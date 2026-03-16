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
	"fmt"

	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"

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

// Metadata returns the resource type name.
func (r *ResourceBase[T]) Metadata(
	_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.ResourceType
}

// Create creates the resource in SMC. It is a terraform api
func (r *ResourceBase[T]) Create(ctx context.Context,
	req resource.CreateRequest, resp *resource.CreateResponse) {

	tflog.Debug(ctx, fmt.Sprintf("Beginning Create of resource %s", r.ResourceType))

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
	// Reading the resource state from the SMC is not supported because
	// the returned object differs significantly from the Terraform object
	// due to default values. This causes frequent false-positive change
	// detections.
	tflog.Debug(ctx, fmt.Sprintf("Beginning Read of %s", r.ResourceType))

}
