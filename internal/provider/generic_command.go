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
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	// "os"
	// "strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/customfield"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

var (
	_ action.ActionWithConfigure = (*SmcGenericAction)(nil)
)

// InitialContactResourceModel describes the resource data model.
type SmcGenericActionModel struct {
	CommandHref types.String       `tfsdk:"command_href"`
	Method      types.String       `tfsdk:"method"`
	Query       *map[string]string `tfsdk:"query"`
	Parameters  *map[string]string `tfsdk:"parameters"`
	Body        types.String       `tfsdk:"body"`
	OutputFile  types.String       `tfsdk:"output_file"`
}

func NewSmcGenericAction() action.Action {
	return &SmcGenericAction{}
}

type SmcGenericAction struct {
	Client *smc.SmcClient
}

func (a *SmcGenericAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_command"
}

func (a *SmcGenericAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Generic action to invoke the smc api",
		Attributes: map[string]schema.Attribute{
			"command_href": schema.StringAttribute{
				MarkdownDescription: "url to request the smc api",
				Required:            true,
			},
			"method": schema.StringAttribute{
				MarkdownDescription: "HTTP method to use for the request (eg GET, POST, DELETE). Default is POST",
				Optional:            true,
			},
			"query": schema.MapAttribute{
				MarkdownDescription: "smc api query parameter (eg ?filter=xxxxx)",
				Optional:            true,
				ElementType:         types.StringType,
				CustomType:          customfield.NewMapType[types.String](ctx),
			},
			"parameters": schema.MapAttribute{
				MarkdownDescription: "smc api body as a json dictionary",
				Optional:            true,
				ElementType:         types.StringType,
				CustomType:          customfield.NewMapType[types.String](ctx),
			},
			"body": schema.StringAttribute{
				MarkdownDescription: "smc api body as raw text",
				Optional:            true,
			},
			"output_file": schema.StringAttribute{
				MarkdownDescription: "Path to save the response output",
				Optional:            true,
			},
		},
	}
}

func (r *SmcGenericAction) Configure(
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

func (a *SmcGenericAction) Invoke(
	ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config SmcGenericActionModel
	diag := &resp.Diagnostics

	diag.Append(req.Config.Get(ctx, &config)...)
	if diag.HasError() {
		return
	}

	tflog.Debug(ctx, "generic action", map[string]any{
		"config": config,
	})

	if config.CommandHref.IsNull() || config.CommandHref.ValueString() == "" {
		diag.AddError(
			"Missing Required Attribute",
			"The 'command_href' attribute is required and must not be empty.",
		)
		return
	}

	command_href := config.CommandHref.ValueString()

	if !strings.HasPrefix(command_href, "http") {
		diag.AddError(
			"Invalid Initial Contact Href",
			"The 'command_href' attribute must be a valid URL starting with http or https.",
		)
		return
	}
	if a.Client == nil {
		diag.AddError(
			"SMC Client Not Configured",
			"The SMC client was not properly configured.",
		)
		return
	}
	params := url.Values{}

	action_url := command_href
	if config.Query != nil && len(*config.Query) > 0 {
		for k, v := range *config.Query {
			tflog.Debug(ctx, "adding query param", map[string]any{
				"key":   k,
				"value": v,
			})
			params.Set(k, v)
		}
		action_url = fmt.Sprintf("%s?%s", command_href, params.Encode())
	}

	var body []byte = []byte{}
	var err error
	if config.Parameters != nil && len(*config.Parameters) > 0 {
		tflog.Debug(ctx, "adding body parameters", map[string]any{
			"parameters": *config.Parameters,
		})
		body, err = json.MarshalIndent(config.Parameters, "", "  ")
		if err != nil {
			diag.AddError(
				"Error Marshalling Body Parameters",
				fmt.Sprintf("Could not marshal body parameters: %s",
					err.Error()),
			)
			return
		}

	}

	method := "POST"
	if !config.Method.IsNull() && config.Method.ValueString() != "" {
		method = strings.ToUpper(config.Method.ValueString())
	}

	respData, err := a.Client.DoRequest(smc.Options{
		Method: method,
		URL:    action_url,
		Body:   body,
	})
	if err != nil {
		diag.AddError(
			"Error Creating Initial Contact Configuration",
			fmt.Sprintf("Could not create command configuration for command_href '%s': %s",
				config.CommandHref.ValueString(), err.Error()),
		)
		return
	}

	if respData.StatusCode != 200 && respData.StatusCode != 201 && respData.StatusCode != 202 {
		diag.AddError(
			"Error invoking smc action",
			fmt.Sprintf("Error invoking command_href '%s': %s",
				config.CommandHref.ValueString(), string(respData.Body)),
		)
		return
	}

	if config.OutputFile.IsNull() || config.OutputFile.ValueString() == "" {
		if len(respData.Body) > 0 {
			resp.SendProgress(action.InvokeProgressEvent{
				Message: fmt.Sprintf("command completed successfully, response: %s",
					string(respData.Body)),
			})
		} else {
			resp.SendProgress(action.InvokeProgressEvent{
				Message: fmt.Sprintf(
					"command completed successfully with status %d",
					respData.StatusCode),
			})
		}
		return
	}

	responseBody := []byte("{}")
	if len(respData.Body) > 0 {
		responseBody = respData.Body
	}
	// save the file to output_file
	err = os.WriteFile(config.OutputFile.ValueString(), responseBody, 0600)
	if err != nil {
		diag.AddError(
			"Error Writing Configuration File",
			fmt.Sprintf("Could not write configuration file to '%s': %s",
				config.OutputFile.ValueString(), err.Error()),
		)
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf(
			"Command completed successfully. output written to '%s'",
			config.OutputFile.ValueString()),
	})
}
