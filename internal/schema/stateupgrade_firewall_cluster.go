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

// Hand-written (not generated). State migration for Firewall_Cluster from
// schema version 0 (API 7.4) to version 1 (API 7.5).
//
// Breaking changes between v0 and v1 (same as Single_Firewall):
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

// upgradeFirewallClusterState returns the StateUpgrader map for Firewall_Cluster.
// Called from the generated UpgradeState method (ctx already has use_hcl2 set).
func upgradeFirewallClusterState(ctx context.Context) map[int64]resource.StateUpgrader {
	priorSchema := buildV0PriorSchema(ctx, schema74.GetFirewallClusterSchemaAttributes, schema74.GetFirewallClusterSchemaBlocks)
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &priorSchema,
			StateUpgrader: upgradeFirewallClusterStateV0toV1,
		},
	}
}

// upgradeFirewallClusterStateV0toV1 migrates a Firewall_Cluster state from v0 to v1.
func upgradeFirewallClusterStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var v0 schema74.FirewallClusterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v0)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var v1 FirewallClusterResourceModel
	// Recursively copy all fields. Fields only in v0 (e.g. ZtnaConnectorSettings) are
	// skipped; fields only in v1 (e.g. OpcuaDecryptionMode, ReplicaInternalGateways) remain null.
	common.CopyFieldsByNameRecursive(&v1, &v0)
	// Override PhysicalInterfaces to apply the modem_auth_method → apn_interfaces migration.
	v1.PhysicalInterfaces = migratePhysicalInterfacesV0toV1(v0.PhysicalInterfaces)

	resp.Diagnostics.Append(resp.State.Set(ctx, &v1)...)
}
