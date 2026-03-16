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

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ConvertToHCL2(ctx context.Context, attrs map[string]schema.Attribute, blockDefs map[string]schema.Block) map[string]schema.Attribute {
	for name, block := range ConvertBlocksToHCL2(ctx, blockDefs) {
		attrs[name] = block
	}
	return attrs
}

func ConvertBlocksToHCL2(ctx context.Context, blockDefs map[string]schema.Block) map[string]schema.Attribute {
	attributes := make(map[string]schema.Attribute)

	for name, block := range blockDefs {
		switch b := block.(type) {
		case schema.SingleNestedBlock:
			attributes[name] = schema.SingleNestedAttribute{
				Description:         b.Description,
				MarkdownDescription: b.MarkdownDescription,
				CustomType:          b.CustomType,
				Attributes:          b.Attributes,
				Optional:            true,
				Computed:            false,
			}
		case schema.ListNestedBlock:
			attributes[name] = schema.ListNestedAttribute{
				Description:         b.Description,
				MarkdownDescription: b.MarkdownDescription,
				CustomType:          b.CustomType,
				NestedObject: schema.NestedAttributeObject{
					Attributes: b.NestedObject.Attributes,
				},
				Optional: true,
				Computed: false,
			}
		case schema.SetNestedBlock:
			attributes[name] = schema.SetNestedAttribute{
				Description:         b.Description,
				MarkdownDescription: b.MarkdownDescription,
				CustomType:          b.CustomType,
				NestedObject: schema.NestedAttributeObject{
					Attributes: b.NestedObject.Attributes,
				},
				Optional: true,
				Computed: false,
			}
		}
	}

	return attributes
}
