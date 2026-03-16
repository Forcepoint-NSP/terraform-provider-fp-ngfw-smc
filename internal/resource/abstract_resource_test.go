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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"testing"
)

type HostResourceModel struct {
	ID types.String `tfsdk:"id"`
}

func TestGetId(t *testing.T) {
	type SampleResourceModel struct {
		ID types.String `tfsdk:"id"`
	}

	sample := SampleResourceModel{
		ID: types.StringValue("sample-id"),
	}

	idValue, err := GetField(sample, "ID")
	if err != nil {
		t.Fatalf("GetField returned an error: %v", err)
	}
	id, ok := idValue.(types.String)
	if !ok {
		t.Fatalf("Expected types.String, got %T", idValue)
	}
	if id.ValueString() != "sample-id" {
		t.Fatalf("Expected id 'sample-id', got '%s'", id.ValueString())
	}
}

func TestSetId(t *testing.T) {
	const newID = "new-id"
	const testID = "test-id"

	// Test successful ID setting
	model := &HostResourceModel{
		ID: types.StringValue("old-id"),
	}

	err := setId(model, newID)
	if err != nil {
		t.Fatalf("setId returned error: %v", err)
	}

	if model.ID.ValueString() != newID {
		t.Errorf("expected ID to be '%s', got '%s'", newID, model.ID.ValueString())
	}

	// Test error on non-pointer input
	err = setId(HostResourceModel{}, testID)
	if err == nil {
		t.Error("expected error for non-pointer input")
	}

	// Test error on nil pointer
	var nilModel *HostResourceModel
	err = setId(nilModel, testID)
	if err == nil {
		t.Error("expected error for nil pointer")
	}

	// Test error on struct without ID field
	type NoIdModel struct {
		Name types.String `tfsdk:"name"`
	}
	noIdModel := &NoIdModel{Name: types.StringValue("test")}
	err = setId(noIdModel, testID)
	if err == nil {
		t.Error("expected error for struct without ID field")
	}

	// Test error on wrong field type
	type WrongTypeModel struct {
		ID string `tfsdk:"id"`
	}
	wrongTypeModel := &WrongTypeModel{ID: "old-id"}
	err = setId(wrongTypeModel, newID)
	if err == nil {
		t.Error("expected error for wrong field type")
	}
}
