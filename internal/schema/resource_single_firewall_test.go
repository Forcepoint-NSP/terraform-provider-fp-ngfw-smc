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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	smcresource "github.com/terraform-providers/terraform-provider-smc/internal/resource"
)

func newTestSingleFirewallResource() *SingleFirewallResource {
	return &SingleFirewallResource{
		ResourceBase: smcresource.ResourceBase[SingleFirewallResourceModel]{
			ResourceType: "single_fw",
		},
	}
}

func TestValidatePhysicalInterfaces(t *testing.T) {
	r := newTestSingleFirewallResource()

	tests := []struct {
		name       string
		interfaces *[]AbstractPhysicalInterfaceWrapperResourceModel
		wantError  bool
	}{
		{
			name:       "nil physical_interfaces",
			interfaces: nil,
			wantError:  false,
		},
		{
			name:       "empty physical_interfaces",
			interfaces: &[]AbstractPhysicalInterfaceWrapperResourceModel{},
			wantError:  false,
		},
		{
			name: "one entry, one type set",
			interfaces: &[]AbstractPhysicalInterfaceWrapperResourceModel{
				{PhysicalInterface: &PhysicalInterfaceResourceModel{}},
			},
			wantError: false,
		},
		{
			name: "two entries, each with one type set",
			interfaces: &[]AbstractPhysicalInterfaceWrapperResourceModel{
				{PhysicalInterface: &PhysicalInterfaceResourceModel{}},
				{TunnelInterface: &TunnelInterfaceResourceModel{}},
			},
			wantError: false,
		},
		{
			name: "one entry, no type set",
			interfaces: &[]AbstractPhysicalInterfaceWrapperResourceModel{
				{},
			},
			wantError: true,
		},
		{
			name: "one entry, two types set",
			interfaces: &[]AbstractPhysicalInterfaceWrapperResourceModel{
				{
					PhysicalInterface: &PhysicalInterfaceResourceModel{},
					TunnelInterface:   &TunnelInterfaceResourceModel{},
				},
			},
			wantError: true,
		},
		{
			name: "two entries, second one invalid",
			interfaces: &[]AbstractPhysicalInterfaceWrapperResourceModel{
				{PhysicalInterface: &PhysicalInterfaceResourceModel{}},
				{},
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := SingleFirewallResourceModel{PhysicalInterfaces: tc.interfaces}
			resp := &resource.ValidateConfigResponse{}
			r.validatePhysicalInterfaces(data, resp)
			if resp.Diagnostics.HasError() != tc.wantError {
				t.Errorf("HasError() = %v, want %v; diagnostics: %v",
					resp.Diagnostics.HasError(), tc.wantError, resp.Diagnostics)
			}
		})
	}
}
