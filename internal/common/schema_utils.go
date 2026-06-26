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

package common

import (
	"context"
	"reflect"
)

// CopyFieldsByName copies struct fields from src into dst by matching field names.
// Only fields whose reflect.Type is identical in both structs are copied; fields
// present in only one struct, or with differing types, are silently skipped.
// Both dst and src must be non-nil pointers to structs.
//
// This is the recommended pattern for migrating between schema versions when the
// source and destination types differ only by package path (e.g. schema74 → schema)
// but share many value-type fields (types.String, types.Bool, types.Int64,
// customfield.Map[types.String], etc.) that have the same concrete reflect.Type.
func CopyFieldsByName(dst, src any) {
	dv := reflect.ValueOf(dst).Elem()
	sv := reflect.ValueOf(src).Elem()
	st := sv.Type()
	for i := 0; i < st.NumField(); i++ {
		sf := sv.Field(i)
		df := dv.FieldByName(st.Field(i).Name)
		if df.IsValid() && df.CanSet() && sf.Type() == df.Type() {
			df.Set(sf)
		}
	}
}

// CopyFieldsByNameRecursive deep-copies src into dst by matching field names at every
// level of nesting. For fields with identical reflect.Type the value is copied directly.
// For cross-package pointer and slice fields (same logical structure, different package
// path), new target instances are allocated and fields are recursively copied.
// Fields present in only one struct, or whose kinds are incompatible, are skipped.
// Both dst and src must be non-nil pointers to structs.
//
// This is the safe alternative to unsafe.Pointer aliasing for schema version migrations
// where the type trees may have diverged (fields added or removed in nested types).
func CopyFieldsByNameRecursive(dst, src any) {
	copyRecursive(reflect.ValueOf(dst).Elem(), reflect.ValueOf(src).Elem())
}

func copyRecursive(dst, src reflect.Value) {
	switch src.Kind() {
	case reflect.Struct:
		st := src.Type()
		for i := 0; i < st.NumField(); i++ {
			sf := src.Field(i)
			df := dst.FieldByName(st.Field(i).Name)
			if !df.IsValid() || !df.CanSet() {
				continue
			}
			if sf.Type() == df.Type() {
				df.Set(sf)
			} else {
				copyRecursive(df, sf)
			}
		}
	case reflect.Pointer:
		if src.IsNil() || dst.Kind() != reflect.Pointer {
			return
		}
		if src.Type() == dst.Type() {
			dst.Set(src)
			return
		}
		newPtr := reflect.New(dst.Type().Elem())
		copyRecursive(newPtr.Elem(), src.Elem())
		dst.Set(newPtr)
	case reflect.Slice:
		if src.IsNil() || dst.Kind() != reflect.Slice {
			return
		}
		if src.Type() == dst.Type() {
			dst.Set(src)
			return
		}
		n := src.Len()
		newSlice := reflect.MakeSlice(dst.Type(), n, n)
		sElem := src.Type().Elem()
		dElem := newSlice.Type().Elem()
		for i := 0; i < n; i++ {
			if sElem == dElem {
				newSlice.Index(i).Set(src.Index(i))
			} else {
				copyRecursive(newSlice.Index(i), src.Index(i))
			}
		}
		dst.Set(newSlice)
	}
	// Other differing kinds (value-type scalars, maps, etc.): skip.
}

// CopyStruct creates a new *V1 and copies all exact-type-matching fields from *V0
// using CopyFieldsByName. Returns nil if src is nil.
// Use for types that only gained new scalar fields in v1.
func CopyStruct[V0, V1 any](src *V0) *V1 {
	if src == nil {
		return nil
	}
	dst := new(V1)
	CopyFieldsByName(dst, src)
	return dst
}

// CopyStructSlice creates a new *[]V1, copying each v0 element with CopyStruct semantics.
// Returns nil if src is nil.
func CopyStructSlice[V0, V1 any](src *[]V0) *[]V1 {
	if src == nil {
		return nil
	}
	result := make([]V1, len(*src))
	for i := range *src {
		CopyFieldsByName(&result[i], &(*src)[i])
	}
	return &result
}

func UseHCL2(ctx context.Context) bool {
	useHcl2_val := ctx.Value("use_hcl2")
	if useHcl2_val == nil {
		return false
	}
	useHcl2, ok := useHcl2_val.(bool)
	if !ok {
		return false
	}
	return useHcl2
}
