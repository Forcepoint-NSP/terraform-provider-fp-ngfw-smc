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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	fwschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/terraform-providers/terraform-provider-smc/internal/schema74"
)

// v0PhysicalInterface builds a physical_interface wrapper with one single_node_interface.
func v0PhysicalInterface() schema74.AbstractPhysicalInterfaceWrapperResourceModel {
	return schema74.AbstractPhysicalInterfaceWrapperResourceModel{
		PhysicalInterface: &schema74.PhysicalInterfaceResourceModel{
			InterfaceId: types.StringValue("0"),
			Interfaces: &[]schema74.EngineInterfaceWrapperResourceModel{
				{
					SingleNodeInterface: &schema74.SingleNodeInterfaceResourceModel{
						Nodeid:       types.Int64Value(1),
						Nicid:        types.StringValue("0"),
						Address:      types.StringValue("192.168.100.14"),
						NetworkValue: types.StringValue("192.168.100.00/24"),
						PrimaryMgt:   types.BoolValue(true),
					},
				},
			},
		},
	}
}

// v0ModemInterface builds a modem_interface wrapper matching the test data:
//
//	modem_interface {
//	  interface_id = "0"
//	  interfaces {
//	    single_node_interface {
//	      dynamic = true; dynamic_index = 1; nicid = "0"; nodeid = 1; pin_code = "123456"
//	    }
//	  }
//	  modem_interface_type = "LTE"
//	  name = "Modem 0"
//	}
//
// modemAuthMethod may be empty string (null) or a real value such as "PAP".
func v0ModemInterface(modemAuthMethod string) schema74.AbstractPhysicalInterfaceWrapperResourceModel {
	var authMethod types.String
	if modemAuthMethod == "" {
		authMethod = types.StringNull()
	} else {
		authMethod = types.StringValue(modemAuthMethod)
	}
	return schema74.AbstractPhysicalInterfaceWrapperResourceModel{
		ModemInterface: &schema74.ModemInterfaceResourceModel{
			InterfaceId:        types.StringValue("0"),
			ModemAuthMethod:    authMethod,
			ModemInterfaceType: types.StringValue("LTE"),
			Name:               types.StringValue("Modem 0"),
			Interfaces: &[]schema74.EngineInterfaceWrapperResourceModel{
				{
					SingleNodeInterface: &schema74.SingleNodeInterfaceResourceModel{
						Dynamic:      types.BoolValue(true),
						DynamicIndex: types.Int64Value(1),
						Nicid:        types.StringValue("0"),
						Nodeid:       types.Int64Value(1),
						PinCode:      types.StringValue("123456"),
					},
				},
			},
		},
	}
}

// buildUpgradeRequest encodes v0State into a tfsdk.State using the v0 PriorSchema,
// then wraps it in an UpgradeStateRequest.
func buildUpgradeRequest(t *testing.T, ctx context.Context, v0State *schema74.SingleFirewallResourceModel) (resource.UpgradeStateRequest, fwschema.Schema) {
	t.Helper()
	priorSchema := buildV0PriorSchema(ctx,
		schema74.GetSingleFirewallSchemaAttributes,
		schema74.GetSingleFirewallSchemaBlocks)

	state := tfsdk.State{Schema: priorSchema}
	diags := state.Set(ctx, v0State)
	require.Falsef(t, diags.HasError(), "failed to encode v0 state: %v", diags)

	return resource.UpgradeStateRequest{State: &state}, priorSchema
}

// runUpgrade invokes upgradeSingleFirewallStateV0toV1 and returns the decoded v1 model.
func runUpgrade(t *testing.T, ctx context.Context, req resource.UpgradeStateRequest) SingleFirewallResourceModel {
	t.Helper()
	v1Schema := fwschema.Schema{
		Attributes: GetSingleFirewallSchemaAttributes(ctx),
		Blocks:     GetSingleFirewallSchemaBlocks(ctx),
	}
	resp := &resource.UpgradeStateResponse{
		State: tfsdk.State{Schema: v1Schema},
	}
	upgradeSingleFirewallStateV0toV1(ctx, req, resp)
	require.Falsef(t, resp.Diagnostics.HasError(), "upgrade produced errors: %v", resp.Diagnostics)

	var v1 SingleFirewallResourceModel
	diags := resp.State.Get(ctx, &v1)
	require.Falsef(t, diags.HasError(), "failed to decode v1 state: %v", diags)
	return v1
}

// TestUpgradeSingleFirewallStateV0toV1 tests the state migration from v0 (API 7.4)
// to v1 (API 7.5) using the data shape described in the file header comment.
func TestUpgradeSingleFirewallStateV0toV1(t *testing.T) {
	ctx := context.Background()

	t.Run("physical_interface fields preserved", func(t *testing.T) {
		v0 := &schema74.SingleFirewallResourceModel{
			ID:   types.StringValue("http://mysmc:8082/7.4/elements/single_fw/1"),
			Name: types.StringValue("test_fw"),
			PhysicalInterfaces: &[]schema74.AbstractPhysicalInterfaceWrapperResourceModel{
				v0PhysicalInterface(),
			},
		}
		req, _ := buildUpgradeRequest(t, ctx, v0)
		v1 := runUpgrade(t, ctx, req)

		require.NotNil(t, v1.PhysicalInterfaces)
		require.Len(t, *v1.PhysicalInterfaces, 1)

		pi := (*v1.PhysicalInterfaces)[0].PhysicalInterface
		require.NotNil(t, pi)
		assert.Equal(t, "0", pi.InterfaceId.ValueString())
		require.NotNil(t, pi.Interfaces)
		require.Len(t, *pi.Interfaces, 1)
		sni := (*pi.Interfaces)[0].SingleNodeInterface
		require.NotNil(t, sni)
		assert.Equal(t, "192.168.100.14", sni.Address.ValueString())
		assert.Equal(t, "192.168.100.00/24", sni.NetworkValue.ValueString())
		assert.True(t, sni.PrimaryMgt.ValueBool())
		assert.Equal(t, "0", sni.Nicid.ValueString())
		assert.Equal(t, int64(1), sni.Nodeid.ValueInt64())
	})

	t.Run("modem_interface fields preserved when modem_auth_method is null", func(t *testing.T) {
		v0 := &schema74.SingleFirewallResourceModel{
			ID:   types.StringValue("http://mysmc:8082/7.4/elements/single_fw/1"),
			Name: types.StringValue("test_fw"),
			PhysicalInterfaces: &[]schema74.AbstractPhysicalInterfaceWrapperResourceModel{
				v0PhysicalInterface(),
				v0ModemInterface(""), // no modem_auth_method
			},
		}
		req, _ := buildUpgradeRequest(t, ctx, v0)
		v1 := runUpgrade(t, ctx, req)

		require.NotNil(t, v1.PhysicalInterfaces)
		require.Len(t, *v1.PhysicalInterfaces, 2)

		modem := (*v1.PhysicalInterfaces)[1].ModemInterface
		require.NotNil(t, modem)
		assert.Equal(t, "0", modem.InterfaceId.ValueString())
		assert.Equal(t, "LTE", modem.ModemInterfaceType.ValueString())
		assert.Equal(t, "Modem 0", modem.Name.ValueString())

		// interfaces inside the modem are preserved
		require.NotNil(t, modem.Interfaces)
		require.Len(t, *modem.Interfaces, 1)
		sni := (*modem.Interfaces)[0].SingleNodeInterface
		require.NotNil(t, sni)
		assert.True(t, sni.Dynamic.ValueBool())
		assert.Equal(t, int64(1), sni.DynamicIndex.ValueInt64())
		assert.Equal(t, "0", sni.Nicid.ValueString())
		assert.Equal(t, int64(1), sni.Nodeid.ValueInt64())

		// null modem_auth_method → no apn_interfaces entry created
		apns := modem.ApnInterfaces
		assert.True(t, apns == nil || len(*apns) == 0,
			"expected no apn_interfaces for null modem_auth_method, got %v", apns)
	})

	t.Run("modem_auth_method migrated to apn_interfaces[0]", func(t *testing.T) {
		v0 := &schema74.SingleFirewallResourceModel{
			ID:   types.StringValue("http://mysmc:8082/7.4/elements/single_fw/1"),
			Name: types.StringValue("test_fw"),
			PhysicalInterfaces: &[]schema74.AbstractPhysicalInterfaceWrapperResourceModel{
				v0PhysicalInterface(),
				v0ModemInterface("PAP"),
			},
		}
		req, _ := buildUpgradeRequest(t, ctx, v0)
		v1 := runUpgrade(t, ctx, req)

		require.NotNil(t, v1.PhysicalInterfaces)
		modem := (*v1.PhysicalInterfaces)[1].ModemInterface
		require.NotNil(t, modem)

		require.NotNil(t, modem.ApnInterfaces, "apn_interfaces must be non-nil when modem_auth_method was set")
		require.Len(t, *modem.ApnInterfaces, 1, "expected exactly one apn_interfaces entry")
		assert.Equal(t, "PAP", (*modem.ApnInterfaces)[0].ModemAuthMethod.ValueString(),
			"modem_auth_method should be migrated into apn_interfaces[0].modem_auth_method")
	})
}
