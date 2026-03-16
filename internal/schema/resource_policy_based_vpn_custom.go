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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

// close and save the vpn before deleting it
func (r *PolicyBasedVpnResource) BeforeDelete(
	ctx context.Context, _ *PolicyBasedVpnResourceModel, id string) error {

	mgr := smc.GetVpnEditStateManager()
	state, ok := mgr.VpnEditStates.Load(id)
	if !ok {
		tflog.Debug(ctx, fmt.Sprintf("vpn %s not open for edit, no need to close", id))
		return nil
	}
	tflog.Debug(ctx, fmt.Sprintf("Need to close vpn %s", id))

	// best effort to save and close. Ignore errors
	err := smc.SendVpnCommand(state.(*smc.VpnEditState).SmcClient, id, "save")
	if err != nil {
		tflog.Warn(ctx,
			fmt.Sprintf("Failed to save vpn %s: %s", id, err.Error()))
	}
	err = smc.SendVpnCommand(state.(*smc.VpnEditState).SmcClient, id, "close")
	if err != nil {
		tflog.Error(ctx,
			fmt.Sprintf("Failed to close vpn %s: %s", id, err.Error()))
	}
	mgr.VpnEditStates.Delete(id)
	return nil
}
