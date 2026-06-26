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
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// ---- test model types -------------------------------------------------------

type tvA struct {
	Name string `tfsdk:"name"`
}
type tvB struct {
	Value string `tfsdk:"value"`
}

// twrapper is an anyOf discriminator wrapper: all exported fields are *struct.
type twrapper struct {
	A *tvA `tfsdk:"a"`
	B *tvB `tfsdk:"b"`
}

// ---- helpers ----------------------------------------------------------------

func errDetails(d diag.Diagnostics) []string {
	out := make([]string, len(d))
	for i, dd := range d {
		out[i] = dd.Detail()
	}
	return out
}

func hasErrContaining(d diag.Diagnostics, substr string) bool {
	for _, s := range errDetails(d) {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// ---- TestCountNonNilPtrFields -----------------------------------------------

func TestCountNonNilPtrFields(t *testing.T) {
	tests := []struct {
		name string
		v    twrapper
		want int
	}{
		{"both nil", twrapper{}, 0},
		{"only A set", twrapper{A: &tvA{}}, 1},
		{"only B set", twrapper{B: &tvB{}}, 1},
		{"both set", twrapper{A: &tvA{}, B: &tvB{}}, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := countNonNilPtrFields(reflect.ValueOf(tc.v)); got != tc.want {
				t.Errorf("countNonNilPtrFields = %d, want %d", got, tc.want)
			}
		})
	}
}

// ---- TestValidateAnyOf ------------------------------------------------------

func TestValidateAnyOf(t *testing.T) {
	tests := []struct {
		name      string
		val       any
		fieldName string
		wantError bool
		wantInMsg string
	}{
		{
			name:      "nil value",
			val:       nil,
			fieldName: "f",
			wantError: false,
		},
		{
			name:      "nil pointer to wrapper",
			val:       (*twrapper)(nil),
			fieldName: "f",
			wantError: false,
		},
		{
			name:      "valid single: exactly one variant set",
			val:       &twrapper{A: &tvA{}},
			fieldName: "f",
			wantError: false,
		},
		{
			name:      "invalid single: no variant set",
			val:       &twrapper{},
			fieldName: "myfield",
			wantError: true,
			wantInMsg: "myfield",
		},
		{
			name:      "invalid single: both variants set",
			val:       &twrapper{A: &tvA{}, B: &tvB{}},
			fieldName: "myfield",
			wantError: true,
			wantInMsg: "myfield",
		},
		{
			name:      "nil pointer to slice",
			val:       (*[]twrapper)(nil),
			fieldName: "f",
			wantError: false,
		},
		{
			name:      "empty slice",
			val:       &[]twrapper{},
			fieldName: "f",
			wantError: false,
		},
		{
			name:      "valid slice: each element has one variant",
			val:       &[]twrapper{{A: &tvA{}}, {B: &tvB{}}},
			fieldName: "f",
			wantError: false,
		},
		{
			name:      "invalid slice: first element empty",
			val:       &[]twrapper{{}, {A: &tvA{}}},
			fieldName: "items",
			wantError: true,
			wantInMsg: "items[0]",
		},
		{
			name:      "invalid slice: second element empty",
			val:       &[]twrapper{{A: &tvA{}}, {}},
			fieldName: "items",
			wantError: true,
			wantInMsg: "items[1]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var d diag.Diagnostics
			ValidateAnyOf(tc.val, tc.fieldName, &d)
			if d.HasError() != tc.wantError {
				t.Errorf("HasError() = %v, want %v; details: %v", d.HasError(), tc.wantError, errDetails(d))
			}
			if tc.wantError && tc.wantInMsg != "" && !hasErrContaining(d, tc.wantInMsg) {
				t.Errorf("expected %q in error detail; got: %v", tc.wantInMsg, errDetails(d))
			}
		})
	}
}
