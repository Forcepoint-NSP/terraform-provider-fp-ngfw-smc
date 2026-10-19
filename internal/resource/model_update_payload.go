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
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// MergeResourceUpdatePayload backfills the user's plan (dest) with
// values from a freshly-read copy of the current SMC object (src), for
// use when building an update request body. dest is mutated in place
// and is meant to be marshaled as the request payload afterwards.
//
// Without this, sending the plan directly means any scalar attribute
// the user never configured (comment, etc.) goes out as null, and SMC
// clears it - so an update to one unrelated attribute silently wipes
// out values that were only ever set out-of-band (e.g. via the SMC
// GUI) and that Terraform was never told to manage.
//
// This intentionally keeps the same src/dest direction and recursion
// as MergeResourceModels/copyComputedField: nested objects and lists
// are matched and walked exactly as mergeStructs/mergeSlice already do
// (by id, arbitrarily deep - e.g. every level of a nested routing_node
// tree), so computed fields like each level's own "key" keep getting
// backfilled from src at every depth, the same as before. The only
// addition is copyPayloadField's plain-scalar branch: unlike
// copyComputedField (which only ever touches fpro-tagged fields),
// it also fills in a non-fpro scalar attribute from src when the plan
// left it unset. Nested objects/lists are never wholesale-replaced
// here - only walked field-by-field - so the plan stays authoritative
// for their membership/structure, same as today.
func MergeResourceUpdatePayload(ctx context.Context, src, dest any) error {
	return mergeStructs(ctx, reflect.ValueOf(src), reflect.ValueOf(dest), copyPayloadField)
}

func copyPayloadField(ctx context.Context, params CopyFieldParams) (bool, error) {
	if params.fieldType.Tag.Get("fpro") != "" || params.fieldType.Tag.Get("tfsdk") == "id" {
		// computed fields: same as the create/update merge path,
		// always take src's (freshly read) value.
		return copyComputedField(ctx, params)
	}

	if params.destField.Kind() == reflect.Struct {
		return copyPayloadScalarField(params)
	}

	// nested objects (Ptr) and lists (Slice): let mergeStructs's
	// default traversal recurse - id-matched via mergeSlice for
	// lists - applying this same copyFieldFunc arbitrarily deep, so
	// nested computed fields get backfilled at every level too. The
	// plan's own membership/structure is never replaced wholesale.
	return false, nil
}

func copyPayloadScalarField(params CopyFieldParams) (bool, error) {
	dest := params.destField
	if !dest.CanInterface() {
		return false, nil
	}

	destAttr, ok := dest.Interface().(attr.Value)
	if !ok {
		// not a terraform value wrapper (e.g. a plain nested struct);
		// let mergeStructs recurse into its fields.
		return false, nil
	}

	if destAttr.IsNull() || destAttr.IsUnknown() {
		// plan didn't configure this attribute: keep src's current
		// (freshly read) value instead of sending null and clearing
		// it on the server.
		if s, ok := params.srcField.Interface().(types.String); ok && isRedactedValue(s.ValueString()) {
			return true, nil
		}
		dest.Set(params.srcField)
	}
	// else: the plan already has an explicit value - leave it as-is.

	return true, nil
}
