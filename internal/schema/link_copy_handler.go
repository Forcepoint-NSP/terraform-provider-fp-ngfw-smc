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

package schema

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terraform-providers/terraform-provider-smc/internal/customfield"
	"reflect"
)

// special handling for "link" field to populate "lk" map
func CopyLinkField(ctx context.Context, srcField, destVal reflect.Value) error {
	lkField := destVal.FieldByName("Lk")

	elements := make(map[string]attr.Value)
	link := srcField.Interface().(customfield.NestedObjectList[ApiLinkResourceModel])

	if !link.IsNullOrUnknown() {
		links, diags := link.AsStructSliceT(ctx)
		if diags.HasError() {
			return fmt.Errorf("failed to convert links to struct slice: %v", diags)
		}

		for _, link := range links {
			if !link.Rel.IsNull() && !link.Href.IsNull() {
				elements[link.Rel.ValueString()] = link.Href
			}
		}
	}

	elementsMap, _ := customfield.NewMap[types.String](ctx, elements)
	lkField.Set(reflect.ValueOf(elementsMap))
	return nil
}
