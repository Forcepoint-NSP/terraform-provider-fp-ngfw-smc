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
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// ValidateAnyOf validates that exactly one pointer field in an anyOf wrapper struct is set.
// val must be a *WrapperStruct or *[]WrapperStruct, where WrapperStruct contains only
// pointer fields (one per variant). nil/zero val is silently skipped.
func ValidateAnyOf(val any, fieldName string, diags *diag.Diagnostics) {
	if val == nil {
		return
	}
	rv := reflect.ValueOf(val)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			count := countNonNilPtrFields(rv.Index(i))
			if count != 1 {
				diags.AddError(
					fmt.Sprintf("Invalid %s entry", fieldName),
					fmt.Sprintf("%s[%d] must have exactly one type defined, got %d", fieldName, i, count),
				)
			}
		}
	case reflect.Struct:
		count := countNonNilPtrFields(rv)
		if count != 1 {
			diags.AddError(
				fmt.Sprintf("Invalid %s", fieldName),
				fmt.Sprintf("%s must have exactly one type defined, got %d", fieldName, count),
			)
		}
	}
}

// countNonNilPtrFields counts the number of non-nil pointer fields in a struct value.
func countNonNilPtrFields(v reflect.Value) int {
	count := 0
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Ptr && !f.IsNil() {
			count++
		}
	}
	return count
}
