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

// Hand-written (not generated). Shared helpers for state migrations (v0 → v1).
package schema

import (
	"context"

	fwschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/terraform-providers/terraform-provider-smc/internal/common"
	"github.com/terraform-providers/terraform-provider-smc/internal/schema74"
)

// buildV0PriorSchema builds a prior schema from attribute/block getter functions.
// Pass schema74 getters to produce an accurate v0 schema for state decoding.
func buildV0PriorSchema(
	ctx context.Context,
	getAttributes func(context.Context) map[string]fwschema.Attribute,
	getBlocks func(context.Context) map[string]fwschema.Block,
) fwschema.Schema {
	return fwschema.Schema{
		Attributes: getAttributes(ctx),
		Blocks:     getBlocks(ctx),
	}
}

// migratePhysicalInterfacesV0toV1 rebuilds the physical_interfaces slice for v1.
func migratePhysicalInterfacesV0toV1(
	v0 *[]schema74.AbstractPhysicalInterfaceWrapperResourceModel) *[]AbstractPhysicalInterfaceWrapperResourceModel {
	if v0 == nil {
		return nil
	}
	result := make([]AbstractPhysicalInterfaceWrapperResourceModel, len(*v0))
	for i, w := range *v0 {
		// Recursively copy all wrapper fields. New fields (ApnInterface) stay nil.
		// SwitchInterfacePort (v0 name) → SwitchInterface (v1 name): skipped (name mismatch).
		common.CopyFieldsByNameRecursive(&result[i], &w)
		// Override ModemInterface to apply the modem_auth_method → apn_interfaces migration.
		result[i].ModemInterface = migrateModemInterfaceV0toV1(w.ModemInterface)
	}
	return &result
}

// migrateModemInterfaceV0toV1 migrates modem_auth_method → apn_interfaces[0].modem_auth_method.
func migrateModemInterfaceV0toV1(v0 *schema74.ModemInterfaceResourceModel) *ModemInterfaceResourceModel {
	if v0 == nil {
		return nil
	}
	v1 := &ModemInterfaceResourceModel{}
	common.CopyFieldsByNameRecursive(v1, v0)
	if !v0.ModemAuthMethod.IsNull() && !v0.ModemAuthMethod.IsUnknown() {
		v1.ApnInterfaces = &[]ApnInterfaceResourceModel{{
			ModemAuthMethod: v0.ModemAuthMethod,
		}}
	}
	return v1
}
