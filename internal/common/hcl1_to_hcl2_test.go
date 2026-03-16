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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/terraform-providers/terraform-provider-smc/internal/customfield"
)

func TestConvertBlocksToHCL2_SingleNestedBlock(t *testing.T) {
	ctx := context.Background()

	blocks := map[string]schema.Block{
		"test_single": schema.SingleNestedBlock{
			Description: "Test single nested block",
			Attributes: map[string]schema.Attribute{
				"field1": schema.StringAttribute{
					Description: "Field 1",
					Optional:    true,
				},
			},
		},
	}

	attrs := ConvertBlocksToHCL2(ctx, blocks)
	if len(attrs) != 1 {
		t.Fatalf("Expected 1 attribute, got %d", len(attrs))
	}

	attr, ok := attrs["test_single"]
	if !ok {
		t.Fatal("Expected 'test_single' attribute not found")
	}

	singleAttr, ok := attr.(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("Expected SingleNestedAttribute, got %T", attr)
	}

	if singleAttr.Description != "Test single nested block" {
		t.Errorf("Expected description 'Test single nested block', got '%s'", singleAttr.Description)
	}

	if !singleAttr.Optional {
		t.Error("Expected Optional to be true")
	}

	if singleAttr.Computed {
		t.Error("Expected Computed to be false")
	}

	if len(singleAttr.Attributes) != 1 {
		t.Errorf("Expected 1 nested attribute, got %d", len(singleAttr.Attributes))
	}
}

func TestConvertBlocksToHCL2_ListNestedBlock(t *testing.T) {
	ctx := context.Background()

	blocks := map[string]schema.Block{
		"test_list": schema.ListNestedBlock{
			Description: "Test list nested block",
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"field1": schema.StringAttribute{
						Description: "Field 1",
						Required:    true,
					},
					"field2": schema.Int64Attribute{
						Description: "Field 2",
						Optional:    true,
					},
				},
			},
		},
	}

	attrs := ConvertBlocksToHCL2(ctx, blocks)

	if len(attrs) != 1 {
		t.Fatalf("Expected 1 attribute, got %d", len(attrs))
	}

	attr, ok := attrs["test_list"]
	if !ok {
		t.Fatal("Expected 'test_list' attribute not found")
	}

	listAttr, ok := attr.(schema.ListNestedAttribute)
	if !ok {
		t.Fatalf("Expected ListNestedAttribute, got %T", attr)
	}

	if listAttr.Description != "Test list nested block" {
		t.Errorf("Expected description 'Test list nested block', got '%s'", listAttr.Description)
	}

	if !listAttr.Optional {
		t.Error("Expected Optional to be true")
	}

	if listAttr.Computed {
		t.Error("Expected Computed to be false")
	}

	if len(listAttr.NestedObject.Attributes) != 2 {
		t.Errorf("Expected 2 nested attributes, got %d", len(listAttr.NestedObject.Attributes))
	}
}

func TestConvertBlocksToHCL2_SetNestedBlock(t *testing.T) {
	ctx := context.Background()

	blocks := map[string]schema.Block{
		"test_set": schema.SetNestedBlock{
			Description: "Test set nested block",
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "Name field",
						Required:    true,
					},
				},
			},
		},
	}

	attrs := ConvertBlocksToHCL2(ctx, blocks)

	if len(attrs) != 1 {
		t.Fatalf("Expected 1 attribute, got %d", len(attrs))
	}

	attr, ok := attrs["test_set"]
	if !ok {
		t.Fatal("Expected 'test_set' attribute not found")
	}

	setAttr, ok := attr.(schema.SetNestedAttribute)
	if !ok {
		t.Fatalf("Expected SetNestedAttribute, got %T", attr)
	}

	if setAttr.Description != "Test set nested block" {
		t.Errorf("Expected description 'Test set nested block', got '%s'", setAttr.Description)
	}

	if !setAttr.Optional {
		t.Error("Expected Optional to be true")
	}

	if setAttr.Computed {
		t.Error("Expected Computed to be false")
	}

	if len(setAttr.NestedObject.Attributes) != 1 {
		t.Errorf("Expected 1 nested attribute, got %d", len(setAttr.NestedObject.Attributes))
	}
}

func TestConvertBlocksToHCL2_EmptyMap(t *testing.T) {
	ctx := context.Background()

	blocks := map[string]schema.Block{}

	attrs := ConvertBlocksToHCL2(ctx, blocks)

	if len(attrs) != 0 {
		t.Fatalf("Expected 0 attributes, got %d", len(attrs))
	}
}

func TestConvertBlocksToHCL2_MultipleBlocks(t *testing.T) {
	ctx := context.Background()

	blocks := map[string]schema.Block{
		"single_block": schema.SingleNestedBlock{
			Description: "Single block",
			Attributes: map[string]schema.Attribute{
				"field": schema.StringAttribute{Optional: true},
			},
		},
		"list_block": schema.ListNestedBlock{
			Description: "List block",
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"item": schema.StringAttribute{Required: true},
				},
			},
		},
		"set_block": schema.SetNestedBlock{
			Description: "Set block",
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"value": schema.Int64Attribute{Optional: true},
				},
			},
		},
	}

	attrs := ConvertBlocksToHCL2(ctx, blocks)

	if len(attrs) != 3 {
		t.Fatalf("Expected 3 attributes, got %d", len(attrs))
	}

	if _, ok := attrs["single_block"].(schema.SingleNestedAttribute); !ok {
		t.Error("Expected 'single_block' to be SingleNestedAttribute")
	}

	if _, ok := attrs["list_block"].(schema.ListNestedAttribute); !ok {
		t.Error("Expected 'list_block' to be ListNestedAttribute")
	}

	if _, ok := attrs["set_block"].(schema.SetNestedAttribute); !ok {
		t.Error("Expected 'set_block' to be SetNestedAttribute")
	}
}

func TestConvertBlocksToHCL2_WithCustomType(t *testing.T) {
	ctx := context.Background()

	type TestResourceModel struct{}

	customType := customfield.NewNestedObjectType[TestResourceModel](ctx)

	blocks := map[string]schema.Block{
		"test_custom": schema.SingleNestedBlock{
			Description: "Test with custom type",
			CustomType:  customType,
			Attributes: map[string]schema.Attribute{
				"field": schema.StringAttribute{Optional: true},
			},
		},
	}

	attrs := ConvertBlocksToHCL2(ctx, blocks)

	singleAttr := attrs["test_custom"].(schema.SingleNestedAttribute)

	if singleAttr.CustomType == nil {
		t.Error("Expected CustomType to be set")
	}

	if singleAttr.CustomType.String() != customType.String() {
		t.Error("CustomType was not preserved correctly")
	}
}
