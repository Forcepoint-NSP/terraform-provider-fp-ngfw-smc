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

// Code not auto-generated
package provider

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

var (
	_ action.ActionWithConfigure = (*InitialContactAction)(nil)
)

// InitialContactResourceModel describes the resource data model.
type InitialContactActionModel struct {
	InitialContactHref types.String `tfsdk:"initial_contact_href"`
	EnableSsh          types.Bool   `tfsdk:"enable_ssh"`
	TimeZone           types.String `tfsdk:"time_zone"`
	Keyboard           types.String `tfsdk:"keyboard"`
	RootPassword       types.String `tfsdk:"root_password"`
	InstallOnServer    types.Bool   `tfsdk:"install_on_server"`
	ExportToBase64     types.Bool   `tfsdk:"export_to_base64"`
	OutputFile         types.String `tfsdk:"output_file"`
}

func NewInitialContactAction() action.Action {
	return &InitialContactAction{}
}

type InitialContactAction struct {
	Client *smc.SmcClient
}

func (a *InitialContactAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_initial_contact"
}

func (a *InitialContactAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Action for initial contact configuration for a firewall node.",
		Attributes: map[string]schema.Attribute{
			"initial_contact_href": schema.StringAttribute{
				MarkdownDescription: "href of the initial contact endpoint.",
				Required:            true,
			},
			"enable_ssh": schema.BoolAttribute{
				MarkdownDescription: "Enable SSH access on the firewall node.",
				Optional:            true,
			},
			"time_zone": schema.StringAttribute{
				MarkdownDescription: "The time zone to configure on the firewall node.",
				Optional:            true,
			},
			"keyboard": schema.StringAttribute{
				MarkdownDescription: "The keyboard layout to configure on the firewall node.",
				Optional:            true,
			},
			"root_password": schema.StringAttribute{
				MarkdownDescription: "The root password for the firewall node.",
				Optional:            true,
			},
			"install_on_server": schema.BoolAttribute{
				MarkdownDescription: "Whether to install on the server.",
				Optional:            true,
			},
			"export_to_base64": schema.BoolAttribute{
				MarkdownDescription: "Export the configuration as base64 encoded string.",
				Optional:            true,
			},
			"output_file": schema.StringAttribute{
				MarkdownDescription: "Path to save the base64 encoded configuration file locally.",
				Required:            true,
			},
		},
	}
}

func (r *InitialContactAction) Configure(
	ctx context.Context, req action.ConfigureRequest,
	resp *action.ConfigureResponse) {

	if req.ProviderData == nil {
		// we must not set error here, as the provider might be not configured yet.
		// the Configure method will be called again once the provider is configured.
		return
	}

	client, ok := req.ProviderData.(*smc.SmcClient)
	if !ok || client == nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *smc.SmClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.Client = client
}

func (r *InitialContactAction) sendInitialContactRequest(
	ctx context.Context, data *InitialContactActionModel, diag *diag.Diagnostics) {

	tflog.Debug(ctx, "Creating initial contact configuration", map[string]any{
		"initial_contact_href": data.InitialContactHref.ValueString(),
	})

	// Validate mandatory fields
	if data.InitialContactHref.IsNull() || data.InitialContactHref.ValueString() == "" {
		diag.AddError(
			"Missing Required Attribute",
			"The 'initial_contact_href' attribute is required and must not be empty.",
		)
		return
	}

	initial_contact_href := data.InitialContactHref.ValueString()

	if !strings.HasPrefix(initial_contact_href, "http") {
		diag.AddError(
			"Invalid Initial Contact Href",
			"The 'initial_contact_href' attribute must be a valid URL starting with http or https.",
		)
		return
	}
	if !strings.HasSuffix(initial_contact_href, "initial_contact") {
		diag.AddError(
			"Invalid Initial Contact Href",
			"The 'initial_contact_href' attribute must end with 'initial_contact'",
		)
		return
	}
	if r.Client == nil {
		diag.AddError(
			"SMC Client Not Configured",
			"The SMC client was not properly configured.",
		)
		return
	}
	params := url.Values{
		"enable_ssh":        {strconv.FormatBool(data.EnableSsh.ValueBool())},
		"time_zone":         {data.TimeZone.ValueString()},
		"keyboard":          {data.Keyboard.ValueString()},
		"root_password":     {data.RootPassword.ValueString()},
		"install_on_server": {strconv.FormatBool(data.InstallOnServer.ValueBool())},
		"export_to_base64":  {strconv.FormatBool(data.ExportToBase64.ValueBool())},
	}

	initial_contact_url := fmt.Sprintf("%s?%s", initial_contact_href, params.Encode())

	respData, err := r.Client.DoRequest(smc.Options{
		Method: "POST",
		URL:    initial_contact_url,
	})
	if err != nil {
		diag.AddError(
			"Error Creating Initial Contact Configuration",
			fmt.Sprintf("Could not create initial contact configuration for initial_contact_href '%s': %s",
				data.InitialContactHref.ValueString(), err.Error()),
		)
		return
	}
	if respData.StatusCode != 200 {
		diag.AddError(
			"Error Creating Initial Contact Configuration",
			fmt.Sprintf("Could not create initial contact configuration for initial_contact_href '%s': %s",
				data.InitialContactHref.ValueString(), string(respData.Body)),
		)
		return
	}

	// save the file to output_file
	err = os.WriteFile(data.OutputFile.ValueString(), respData.Body, 0600)
	if err != nil {
		diag.AddError(
			"Error Writing Configuration File",
			fmt.Sprintf("Could not write configuration file to '%s': %s",
				data.OutputFile.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("Initial contact configuration created for initial_contact_href '%s'", data.InitialContactHref.ValueString()))

}

func (a *InitialContactAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config InitialContactActionModel

	tflog.Debug(ctx, "Invoking Initial Contact action")
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	a.sendInitialContactRequest(ctx, &config, &resp.Diagnostics)
	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("initial contact in progress..."),
	})
}
