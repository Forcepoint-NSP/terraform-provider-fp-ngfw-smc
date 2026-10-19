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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

var (
	_ action.ActionWithConfigure = (*EngineRefreshPolicyAction)(nil)
)

type EngineRefreshPolicyActionModel struct {
	Href                types.String `tfsdk:"href"`
	PreserveConnections types.Bool   `tfsdk:"preserve_connections"`
	Comment             types.String `tfsdk:"comment"`
	Wait                types.Bool   `tfsdk:"wait"`
	Timeout             types.Int64  `tfsdk:"timeout"`
	Color               types.Bool   `tfsdk:"color"`
}

func NewEngineRefreshPolicyAction() action.Action {
	return &EngineRefreshPolicyAction{}
}

type EngineRefreshPolicyAction struct {
	Client *smc.SmcClient
}

func (a *EngineRefreshPolicyAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_engine_refresh_policy"
}

func (a *EngineRefreshPolicyAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Refresh the policy on an engine. Polls until completion by default.",
		Attributes: map[string]schema.Attribute{
			"href": schema.StringAttribute{
				MarkdownDescription: "Refresh operation URL (value of `link.refresh` from the engine element).",
				Required:            true,
			},
			"preserve_connections": schema.BoolAttribute{
				MarkdownDescription: "When true (default), existing TCP, UDP, and VPN connections are kept alive during the policy refresh. When false, all active connections are terminated and must be re-established after the refresh completes.",
				Optional:            true,
			},
			"comment": schema.StringAttribute{
				MarkdownDescription: "Comment recorded with the refresh operation.",
				Optional:            true,
			},
			"wait": schema.BoolAttribute{
				MarkdownDescription: "Wait for the operation to complete before returning. Defaults to true.",
				Optional:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Maximum seconds to wait when `wait` is true. Defaults to 600.",
				Optional:            true,
			},
			"color": schema.BoolAttribute{
				MarkdownDescription: "Render the progress bar with ANSI green color. Requires a color-capable terminal. Defaults to true.",
				Optional:            true,
			},
		},
	}
}

func (a *EngineRefreshPolicyAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*smc.SmcClient)
	if !ok || client == nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *smc.SmcClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	a.Client = client
}

func (a *EngineRefreshPolicyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config EngineRefreshPolicyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	href := config.Href.ValueString()
	if !strings.HasPrefix(href, "http") {
		resp.Diagnostics.AddError("Invalid href", "The 'href' attribute must be a valid URL starting with http or https.")
		return
	}
	if a.Client == nil {
		resp.Diagnostics.AddError("SMC Client Not Configured", "The SMC client was not properly configured.")
		return
	}

	// Commit pending VPN topology edits before refreshing.
	if err := smc.SaveAndCloseVpns(ctx); err != nil {
		resp.Diagnostics.AddError("Failed to commit VPN topology edits", err.Error())
		return
	}

	params := url.Values{}
	if !config.PreserveConnections.IsNull() {
		params.Set("preserve_connections", fmt.Sprintf("%t", config.PreserveConnections.ValueBool()))
	}
	if !config.Comment.IsNull() && config.Comment.ValueString() != "" {
		params.Set("comment", config.Comment.ValueString())
	}

	actionURL := href
	if len(params) > 0 {
		actionURL = fmt.Sprintf("%s?%s", href, params.Encode())
	}

	tflog.Debug(ctx, "invoking engine refresh policy", map[string]any{"url": actionURL})

	respData, err := a.Client.DoRequest(smc.Options{
		Method:  "POST",
		URL:     actionURL,
		Context: ctx,
	})
	if err != nil {
		resp.Diagnostics.AddError("Error invoking refresh policy", fmt.Sprintf("POST to %s failed: %s", href, err.Error()))
		return
	}
	if respData.StatusCode != 200 && respData.StatusCode != 201 && respData.StatusCode != 202 {
		resp.Diagnostics.AddError("Error invoking refresh policy",
			fmt.Sprintf("POST to %s returned status %d: %s", href, respData.StatusCode, string(respData.Body)))
		return
	}

	// wait=true is the default when not set
	wait := config.Wait.IsNull() || config.Wait.ValueBool()
	if !wait {
		resp.SendProgress(action.InvokeProgressEvent{Message: "refresh policy started (not following)"})
		return
	}

	status, parseErr := parseOperationStatus(respData.Body)
	if parseErr != nil {
		resp.Diagnostics.AddError("Error parsing operation status", parseErr.Error())
		return
	}

	timeout := 600 * time.Second
	if !config.Timeout.IsNull() {
		timeout = time.Duration(config.Timeout.ValueInt64()) * time.Second
	}

	color := config.Color.IsNull() || config.Color.ValueBool()
	if err := followOperation(ctx, a.Client, status, timeout, func(progress float64, message string) {
		resp.SendProgress(action.InvokeProgressEvent{
			Message: renderProgressBar(progress, message, color),
		})
	}); err != nil {
		resp.Diagnostics.AddError("Refresh policy failed", err.Error())
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{Message: "refresh policy completed successfully"})
}
