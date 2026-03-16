// Copyright 2026 Forcepoint LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"fmt"
	"reflect"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type ResourceModelWithId interface {
	// note that the Id here has nothing to do with the
	// terraform resource id. It is an identifier used to
	// match elements in lists of nested objects.
	GetSliceIds(ctx context.Context) []string
}

type SpecialCopyHandler func(ctx context.Context, srcField, destVal reflect.Value) error

var specialCopyHandlers map[string]SpecialCopyHandler = make(map[string]SpecialCopyHandler)

func AddSpecialCopyHandlers(fieldName string, copyHandler SpecialCopyHandler) {
	specialCopyHandlers[fieldName] = copyHandler
}

// dereferenceValue safely dereferences a pointer value
func dereferenceValue(val reflect.Value) reflect.Value {
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return reflect.Value{}
		}
		return val.Elem()
	}
	return val
}

func intersect[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if slices.Contains(b, v) {
			set = append(set, v)
		}
	}

	return set
}

// MergeResourceModels copies fields from src to dest if they have the
// `fpro` tag (that is only the readonly/computed fields).

// It handles nested structs, lists of structs, and pointers, with
// context for logging.
// attrs can be used to limit the merge to specific fields.
func MergeResourceModels(ctx context.Context, src, dest any, attrs ...string) error {
	return mergeStructs(ctx, reflect.ValueOf(src),
		reflect.ValueOf(dest),
		copyComputedField,
		attrs...)
}

func MergeResourceFull(ctx context.Context, src, dest any) error {
	return mergeStructs(ctx, reflect.ValueOf(src),
		reflect.ValueOf(dest),
		copyAllField)
}

func mergeStructs(ctx context.Context,
	srcVal, destVal reflect.Value,
	copyFieldFunc func(ctx context.Context, params CopyFieldParams) (bool, error),
	attrs ...string,
) error {
	srcVal = dereferenceValue(srcVal)
	destVal = dereferenceValue(destVal)

	if !srcVal.IsValid() || !destVal.IsValid() {
		tflog.Debug(ctx, "invalid srcVal or destVal")
		return fmt.Errorf("invalid srcVal or destVal")
	}
	tflog.Debug(ctx, "entering mergeStructs", map[string]any{
		"srcVal_kind": srcVal.Kind().String(),
	})
	if srcVal.Type() != destVal.Type() {
		return fmt.Errorf("type mismatch: %s vs %s", srcVal.Type(), destVal.Type())
	}

	if srcVal.Kind() == reflect.Slice {
		return mergeSlice(ctx, srcVal, destVal, copyFieldFunc, attrs...)
	}

	if srcVal.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct kind, got %s", srcVal.Kind())
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		destField := destVal.Field(i)
		fieldType := srcVal.Type().Field(i)

		tflog.Debug(ctx, "processing srcField", map[string]any{
			"field_name": fieldType.Name,
		})

		copyFieldParams := CopyFieldParams{
			srcField:  srcField,
			destField: destField,
			fieldType: fieldType,
			attrs:     attrs,
			destVal:   destVal,
		}
		copied, err := copyFieldFunc(ctx, copyFieldParams)
		if err != nil {
			return err
		}
		if copied {
			continue
		}

		switch srcField.Kind() {
		case reflect.Ptr:
			if destField.IsNil() {
				tflog.Debug(ctx, "ignoring destField - is nil", map[string]any{
					"field_name": fieldType.Name,
				})
				continue
			}
			mergeStructs(ctx, srcField, destField, copyFieldFunc, attrs...)
		case reflect.Slice:
			err := mergeStructs(ctx, srcField, destField, copyFieldFunc, attrs...)
			if err != nil {
				return err
			}
		case reflect.Struct:
			if srcField.Type() == reflect.TypeOf(types.String{}) ||
				srcField.Type() == reflect.TypeOf(types.Bool{}) ||
				srcField.Type() == reflect.TypeOf(types.Int64{}) {
				tflog.Debug(ctx, "ignoring terraform types field", map[string]any{
					"field_name": fieldType.Name,
					"field_type": srcField.Type().String(),
				})
				continue
			}

			err := mergeStructs(ctx, srcField, destField, copyFieldFunc, attrs...)
			if err != nil {
				return err
			}
		} // end switch
	} // end for each field
	return nil
}

type CopyFieldParams struct {
	srcField  reflect.Value
	destField reflect.Value
	fieldType reflect.StructField // reflect type of the field
	attrs     []string            // white-list of attrs to merge
	destVal   reflect.Value       // the parent struct of destField
}

// copyComputedfield copies only fields with 'fpro' tag ('fpro' stands
// for forcepoint read-only)
// returns true if the field was dealt with (copied), false otherwise
func copyComputedField(ctx context.Context, params CopyFieldParams) (bool, error) {
	fproTag := params.fieldType.Tag.Get("fpro")
	isIdField := (params.fieldType.Tag.Get("tfsdk") == "id")

	if isIdField {
		// note: special handling for 'id' field.
		// When the struct is both used as a nested attribute and a
		// resource model (eg PhysicalInterface), it contains an 'id' field.

		// when used as a nested attribute, we don't use the 'id' field,
		// but we still need to set its value because it is declared as
		// "computed" and terraform framework expects it to be set.

		if params.destField.IsValid() == false {
			return true, nil
		}
		idStr := params.destField.Interface().(types.String)

		if idStr.IsNull() || idStr.IsUnknown() {
			empty_str_val := types.StringValue("")
			params.destField.Set(reflect.ValueOf(empty_str_val))
		}
		return true, nil
	}

	if fproTag == "" {
		return false, nil
	}

	// if the function was called with a limited set of
	// attrs, only merge those
	if len(params.attrs) > 0 && (!slices.Contains(params.attrs, fproTag)) {
		tflog.Debug(ctx, "skipping field not in attrs", map[string]any{
			"field_name": params.fieldType.Name,
			"tag":        fproTag,
		})
		return true, nil
	}

	copyFunc, ok := specialCopyHandlers[fproTag]
	if ok {
		// special handling for link field to populate lk map. Note we
		// don't return here and actual link is copied below. We might
		// consider changing this in the future.
		copyFunc(ctx, params.srcField, params.destVal)
	}

	tflog.Debug(ctx, "merging field with fpro tag", map[string]any{
		"field_name": params.fieldType.Name,
	})

	params.destField.Set(params.srcField)
	return true, nil
}

func copyAllField(ctx context.Context, params CopyFieldParams) (bool, error) {
	if params.destField.IsValid() {
		// dest params take precedence
		return false, nil
	}

	params.destField.Set(params.srcField)
	return true, nil
}

// findNestedId recursively searches through struct fields for an ID
// returns e.g. "FirewallNode/Nodeid/2"
func findNestedIds(ctx context.Context, val reflect.Value) ([]string, error) {
	tflog.Debug(ctx, "findNestedIds: searching type", map[string]any{
		"type_name": val.Type().Name(),
		"numFields": val.NumField(),
	})
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := val.Type().Field(i)

		if fieldType.Name == "Link" || fieldType.Name == "Lk" {
			continue
		}
		// Skip invalid fields and anonymous (embedded) fields
		if !fieldVal.IsValid() || fieldType.Anonymous {
			continue
		}

		// Dereference pointer fields
		fieldVal = dereferenceValue(fieldVal)
		if !fieldVal.IsValid() {
			continue
		}

		// Recursively search struct fields
		if fieldVal.Kind() == reflect.Struct && !isTypesField(fieldVal.Type()) {
			nestedIds, err := GetElementIds(ctx, fieldVal)
			if err == nil && nestedIds != nil {
				var results []string
				for _, nestedId := range nestedIds {
					results = append(results, fmt.Sprintf("%s/%s",
						fieldType.Name, nestedId))
				}
				return results, nil
			}
		}
	}

	return nil, fmt.Errorf("Could not find slice ID in type %s",
		val.Type().Name())
}

// isTypesField checks if the field is a Terraform types field that should not be recursed into
func isTypesField(t reflect.Type) bool {
	return t == reflect.TypeOf(types.String{}) ||
		t == reflect.TypeOf(types.Bool{}) ||
		t == reflect.TypeOf(types.Int64{})
}

// GetElementId gets alist of possible identifiers from a struct value
//
// It returns a list because some structs may have multiple identifiers.
// eg RoutingNodeResourceModel has both Href and Name as identifiers.

// e.g. {"FirewallNode/Nodeid/2"} or nil if not found.
//
// refer to test "TestGetElementId" for usage examples.
func GetElementIds(ctx context.Context, val reflect.Value) ([]string, error) {
	// Dereference pointer if needed
	val = dereferenceValue(val)
	if !val.IsValid() || val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("value must be a valid struct, got %s", val.Kind())
	}

	// try to cast to ResourceModelWithId
	if val.CanInterface() {
		var modelWithId ResourceModelWithId
		var ok bool
		if val.CanAddr() {
			modelWithId, ok = val.Addr().Interface().(ResourceModelWithId)
		} else {
			modelWithId, ok = val.Interface().(ResourceModelWithId)
		}
		if ok {
			id := modelWithId.GetSliceIds(ctx)
			if len(id) > 0 {
				return id, nil
			}
			return nil, fmt.Errorf("GetSliceIds() returned empty list")
		}
	}

	// If no direct ID field found, recursively search nested structs
	return findNestedIds(ctx, val)
}

// GetElementById finds the matching element in the dest slice.
//
// params:
// - destSlice reflect.Value of a slice of structs
// - ids eg {"FirewallNode/Nodeid/2"}
//
// returns the matching element reflect.Value or error if not found.
func GetElementById(ctx context.Context, destSlice reflect.Value, ids []string) (reflect.Value, error) {
	for i := 0; i < destSlice.Len(); i++ {
		destId, err := GetElementIds(ctx, destSlice.Index(i))
		if err != nil {
			continue
		}

		commonIds := intersect(destId, ids)
		if len(commonIds) > 0 {
			return destSlice.Index(i), nil
		}
	}
	return reflect.Value{}, fmt.Errorf("element with id %s not found", ids)
}

// mergeSlice merges two slices of structs by matching elements based
// on their IDs.

// The reason for these complex functions is to handle lists of
// nested objects that are not in the same order and to merge their
// fields based on their identifiers (the smc api returns list of
// nested object in unpredictable order).
//
// see test "TestMergeCluster" for usage examples (in this test the
// firewall nodes are not in the same order, but they are merged
// correctly).
func mergeSlice(ctx context.Context, srcSlice, destSlice reflect.Value,
	copyFieldFunc func(ctx context.Context, params CopyFieldParams) (bool, error),
	attrs ...string) error {
	if srcSlice.Len() == 0 {
		tflog.Debug(ctx, "ignoring empty srcField", map[string]any{
			"type_name": srcSlice.Type().Name(),
		})
		return nil
	}

	for j := 0; j < srcSlice.Len(); j++ {
		srcItem := srcSlice.Index(j)

		srcItem = dereferenceValue(srcItem)
		tflog.Debug(ctx, "mergeSlice: processing type", map[string]any{
			"type_name":    srcItem.Type().Name(),
			"kind":         srcItem.Kind().String(),
			"canInterface": srcItem.CanInterface(),
		})
		var destItem reflect.Value
		ids, err := GetElementIds(ctx, srcItem)
		if err != nil || ids == nil {
			tflog.Debug(ctx, "mergeSlice: could not get id for src element, skipping", map[string]any{
				"src_element": srcItem,
			})
			continue
		}
		tflog.Debug(ctx, "mergeSlice: found id:", map[string]any{
			"ids": ids,
		})
		destItem, err = GetElementById(ctx, destSlice, ids)
		if err != nil {
			tflog.Debug(ctx, "mergeSlice: dest element not found for id:",
				map[string]any{
					"ids": ids,
				})
			continue
		}

		err = mergeStructs(ctx, srcItem, destItem, copyFieldFunc, attrs...)
		if err != nil {
			tflog.Debug(ctx, "mergeSlice: error merging structs", map[string]any{
				"error": err.Error(),
			})
			return err
		}
	}

	return nil
}
