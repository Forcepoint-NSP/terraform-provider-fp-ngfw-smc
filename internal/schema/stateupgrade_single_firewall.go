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

// Hand-written (not generated). State migration for Single_Firewall from
// schema version 0 (API 7.4) to version 1 (API 7.5).
//
// Breaking changes between v0 and v1:
//   - modem_auth_method moved from Modem_Interface top-level into apn_interfaces[0].
//   - pin_code and apn_interfaces added to Modem_Interface.
//   - apn_interface added as a wrapper-level block.
//   - Several top-level fields added (opcua_*, replica_internal_gateways).
//   - ztna_connector_settings removed.
//   - native_vlan_id added to Physical_Interface.
//   - id/from_ref added to Vpn_Broker_Interface, Antivirus_Settings, Ssh_Host_Key,
//     Threat_Seeker_Settings.
package schema

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/terraform-providers/terraform-provider-smc/internal/common"
	"github.com/terraform-providers/terraform-provider-smc/internal/schema74"
)

// ---------------------------------------------------------------------------
// State upgrader registration
// ---------------------------------------------------------------------------

// upgradeSingleFirewallState returns the StateUpgrader map for Single_Firewall.
// Called from the generated UpgradeState method (ctx already has use_hcl2 set).
func upgradeSingleFirewallState(ctx context.Context) map[int64]resource.StateUpgrader {
	priorSchema := buildV0PriorSchema(ctx, schema74.GetSingleFirewallSchemaAttributes, schema74.GetSingleFirewallSchemaBlocks)
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &priorSchema,
			StateUpgrader: upgradeSingleFirewallStateV0toV1,
		},
	}
}

// upgradeSingleFirewallStateV0toV1 migrates a Single_Firewall state from v0 to v1.
//
// Strategy: read the full v0 state into the v0 struct (exact schema match), then
// rebuild the v1 struct using copyFieldsByNameUnsafe for the bulk of fields and
// explicit helpers for the handful of types whose layouts changed.
func upgradeSingleFirewallStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var v0 schema74.SingleFirewallResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v0)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var v1 SingleFirewallResourceModel
	// Recursively copy all fields. Cross-package type differences are handled by
	// creating new target instances at each level — no unsafe.Pointer aliases.
	// Fields only in v0 (e.g. ZtnaConnectorSettings) are skipped; fields only in v1
	// (e.g. OpcuaDecryptionMode, ReplicaInternalGateways) remain null.
	common.CopyFieldsByNameRecursive(&v1, &v0)
	// Override PhysicalInterfaces to apply the modem_auth_method → apn_interfaces migration;
	// the recursive copy above left apn_interfaces null everywhere.
	v1.PhysicalInterfaces = migratePhysicalInterfacesV0toV1(v0.PhysicalInterfaces)

	resp.Diagnostics.Append(resp.State.Set(ctx, &v1)...)
}
