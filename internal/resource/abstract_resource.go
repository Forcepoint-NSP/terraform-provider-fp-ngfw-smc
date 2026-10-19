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
	"errors"
	"fmt"
	"net/http"

	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"
	"github.com/terraform-providers/terraform-provider-smc/internal/common"
	"github.com/terraform-providers/terraform-provider-smc/internal/config"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

// Configure adds the provider-configured client to the resource.
func (r *ResourceBase[T]) Configure(
	ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*smc.SmcClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *smc.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.Config = &smc.GenericCRUDConfig{
		ResourceType: r.ResourceType,
		Client:       client,
	}
}

// requireConfig reports a diagnostics error and returns false if the
// resource was never configured with a client (see SMC-67412).
func (r *ResourceBase[T]) requireConfig(diag *diag.Diagnostics) bool {
	if r.Config == nil || r.Config.Client == nil {
		diag.AddError(
			"SMC Provider Never Configured",
			fmt.Sprintf("Could not perform this operation on %s: the smc provider was never configured "+
				"(one of its arguments, e.g. url or api_key, was still unknown). Add an explicit "+
				"depends_on, or set the argument from a variable/data source.", r.ResourceType),
		)
		return false
	}
	return true
}

// Metadata returns the resource type name.
func (r *ResourceBase[T]) Metadata(
	_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.ResourceType
}

// Create creates the resource in SMC. It is a terraform api
func (r *ResourceBase[T]) Create(ctx context.Context,
	req resource.CreateRequest, resp *resource.CreateResponse) {

	tflog.Debug(ctx, fmt.Sprintf("Beginning Create of resource %s", r.ResourceType))

	if !r.requireConfig(&resp.Diagnostics) {
		tflog.Error(ctx, fmt.Sprintf("Create aborted for resource %s: provider not configured", r.ResourceType))
		return
	}

	var data T
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Error reading plan data for resource %s: %v", r.ResourceType, resp.Diagnostics))
		return
	}

	id, err := getId(data)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resource ID",
			fmt.Sprintf("Could not get resource ID: %s", err.Error()),
		)
		return
	}

	var requestResp *smc.ResponseData

	resourceAlreadyExists := (id != "")
	// special case: if ID is set, do an update instead of create
	// we make terraform believe we are creating a new resource, but in reality
	// we are updating an existing one.

	// typical use-case: Routing_Node, Internal_Gateway and
	// Internal_EndPoint are created automatically when creating a
	// Single Firewall. To manage them in terraform, we need to
	// "create" them with the correct ID, which results in an update.
	if resourceAlreadyExists {
		requestResp, err = r.smcUpdateResource(ctx, &data, nil, id, &resp.Diagnostics)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error updating existing resource %s with ID %s: %v", r.ResourceType, id, err))
			return
		}

	} else if r.IsSubResource {
		requestResp, err = r.smcCreateSubResource(ctx, &data, &resp.Diagnostics)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error creating sub-resource %s: %v", r.ResourceType, err))
			return
		}
	} else {
		requestResp, err = r.smcCreateResource(ctx, &data, &resp.Diagnostics)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error creating resource %s: %v", r.ResourceType, err))
			return
		}
	}

	if id == "" {
		id = requestResp.Location
		if id == "" {
			resp.Diagnostics.AddError(
				"Error creating resource",
				fmt.Sprintf("No resource ID returned after creating %s", r.ResourceType),
			)
			return
		}
		setId(&data, id)
	}

	_, err = r.updateTfState(ctx, &data, id, &resp.State, &resp.Diagnostics)
	if err != nil {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Created resource %s successfully", r.ResourceType))
}

// Delete removes the resource from Terraform state and deletes it from the API.
func (r *ResourceBase[T]) Delete(
	ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	if !r.requireConfig(&resp.Diagnostics) {
		tflog.Error(ctx, fmt.Sprintf("Delete aborted for resource %s: provider not configured", r.ResourceType))
		return
	}

	var data T

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.IsSubResource {
		// sub-resources with no "fromRef" were not actually created by terraform.
		// They were created implicitly as part of parent resource creation.
		// So we skip deletion in this case.
		fromRef, err := getFromRefField(data)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting fromRef field",
				fmt.Sprintf("Could not get fromRef field: %s", err.Error()),
			)
			return
		}
		if fromRef == "" {
			tflog.Debug(ctx, fmt.Sprintf("Sub-resource %s implicitly created, skipping delete",
				r.ResourceType))
			return
		}
	}

	href, err := getId(data)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resource ID",
			fmt.Sprintf("Could not get resource ID: %s", err.Error()),
		)
		return
	}
	if href == "" {
		resp.Diagnostics.AddError(
			"Error getting resource ID",
			"Resource ID is missing in the state data.",
		)
		return
	}

	r.smcDeleteResource(ctx, href, &resp.Diagnostics)
}

// Update modifies the resource and updates the Terraform state on success.
func (r *ResourceBase[T]) Update(ctx context.Context, req resource.UpdateRequest,
	resp *resource.UpdateResponse) {

	if !r.requireConfig(&resp.Diagnostics) {
		tflog.Error(ctx, fmt.Sprintf("Update aborted for resource %s: provider not configured", r.ResourceType))
		return
	}

	var data T
	var stateData T

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read current state data to preserve readonly fields
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := getId(stateData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resource ID",
			fmt.Sprintf("Could not get resource ID: %s", err.Error()),
		)
		return
	}
	if id == "" {
		resp.Diagnostics.AddError(
			"Error updating resource",
			"Resource ID is missing in the state data.",
		)
		return
	}
	_, err = r.smcUpdateResource(ctx, &data, nil, id, &resp.Diagnostics)
	if err != nil {
		return
	}

	_, err = r.updateTfState(ctx, &data, id, &resp.State, &resp.Diagnostics)
	if err != nil {
		return
	}
}

func (r *ResourceBase[T]) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Info(ctx, fmt.Sprintf("Importing state for resource %s: id = %s", r.ResourceType, req.ID))
	// resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	if !r.requireConfig(&resp.Diagnostics) {
		tflog.Error(ctx, fmt.Sprintf("Import aborted for resource %s: provider not configured", r.ResourceType))
		return
	}

	var stateData T
	readResp, err := r.Config.ReadResourceByHref(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Import failed",
			fmt.Sprintf("Could not find %s with href '%s': %s", r.Config.ResourceType, req.ID, err.Error()))
		return
	}

	if readResp == nil {
		resp.Diagnostics.AddError(
			"Import failed",
			fmt.Sprintf("No %s found with href '%s'", r.Config.ResourceType, req.ID),
		)
		return
	}

	if err := apijson.UnmarshalRoot(readResp.Body, &stateData); err != nil {
		resp.Diagnostics.AddError(
			"Error parsing resource",
			fmt.Sprintf("Could not parse: %s", err.Error()),
		)
		return
	}

	setId(&stateData, req.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &stateData)...)

}

func (r *ResourceBase[T]) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, fmt.Sprintf("Beginning Read of %s", r.ResourceType))

	// Drift detection is opt-in per resource type via the
	// "drift_detection" key in tfsmc.conf.json: the SMC API can return
	// default values, normalized values, and reordered lists that
	// differ from Terraform state, which would otherwise cause
	// false-positive diffs. See MergeResourceDrift for how those are
	// reconciled. It also gates whether "id"/"link" (and other readOnly
	// attributes) get a UseStateForUnknown plan modifier at all (see
	// schema_objects.j2) - when disabled, Read is a complete no-op, same
	// as before drift detection existed.
	enabled, err := config.IsDriftDetectionEnabled(common.PROVIDER_TYPE_NAME + "_" + r.ResourceType)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading provider configuration",
			fmt.Sprintf("Could not determine if refresh is enabled for %s: %s", r.ResourceType, err.Error()),
		)
		return
	}
	if !enabled {
		return
	}

	if !r.requireConfig(&resp.Diagnostics) {
		tflog.Error(ctx, fmt.Sprintf("Read aborted for resource %s: provider not configured", r.ResourceType))
		return
	}

	var state T
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := getId(state)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting resource ID",
			fmt.Sprintf("Could not get resource ID: %s", err.Error()),
		)
		return
	}
	if id == "" {
		return
	}

	normalizedId, err := r.Config.Client.NormalizeHref(id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error normalizing resource href",
			fmt.Sprintf("Could not normalize href '%s': %s", id, err.Error()),
		)
		return
	}
	id = normalizedId

	readResp, err := r.Config.ReadResourceByHref(id)
	if err != nil {
		var statusErr *smc.StatusError
		if errors.As(err, &statusErr) && statusErr.Code == http.StatusNotFound {
			tflog.Debug(ctx, fmt.Sprintf("Resource %s no longer exists (%s), removing from state", r.ResourceType, id))
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading resource",
			fmt.Sprintf("Could not read %s: %s", r.ResourceType, err.Error()),
		)
		return
	}

	var apiData T
	if err := apijson.UnmarshalRoot(readResp.Body, &apiData); err != nil {
		resp.Diagnostics.AddError(
			"Error parsing resource",
			fmt.Sprintf("Could not parse %s: %s", r.ResourceType, err.Error()),
		)
		return
	}

	if err := MergeResourceDrift(ctx, &apiData, &state); err != nil {
		resp.Diagnostics.AddError(
			"Error reconciling resource drift",
			fmt.Sprintf("Could not reconcile %s: %s", r.ResourceType, err.Error()),
		)
		return
	}
	setId(&state, id)

	// MergeResourceDrift deliberately never refreshes "id"/"link" (Lk)
	// from apiData (see copyDriftField), so a scheme/host/port/version
	// baked into a stale href (e.g. from before an SMC http->https
	// switch, SMC-64735/66510) survives the merge above untouched at
	// every nesting level. Since those attributes carry a
	// UseStateForUnknown plan modifier whenever drift detection is
	// enabled, state must be normalized here or the next plan carries
	// the stale value forward and mismatches what Update's own
	// (unconditional) normalization actually produces - "Provider
	// produced inconsistent result after apply". Pure string rewrite,
	// no extra API call.
	if _, err := normalizeStateHrefs(ctx, r.Config.Client, &state); err != nil {
		resp.Diagnostics.AddError(
			"Error normalizing resource hrefs",
			fmt.Sprintf("Could not normalize hrefs for %s: %s", r.ResourceType, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
