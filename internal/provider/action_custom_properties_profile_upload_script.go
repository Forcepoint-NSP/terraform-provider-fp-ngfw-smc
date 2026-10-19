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
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

var (
	_ action.ActionWithConfigure = (*SmcUploadScriptAction)(nil)
)

type SmcUploadScriptActionModel struct {
	CustomPropertiesProfileHref types.String `tfsdk:"custom_properties_profile_href"`
	ScriptPath                  types.String `tfsdk:"script_path"`
	FieldName                   types.String `tfsdk:"field_name"`
}

func NewSmcCustomPropertiesProfileUploadScriptAction() action.Action {
	return &SmcUploadScriptAction{}
}

type SmcUploadScriptAction struct {
	Client *smc.SmcClient
}

func (a *SmcUploadScriptAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_properties_profile_upload_script"
}

func (a *SmcUploadScriptAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Upload a script file to an SMC element endpoint via multipart/form-data PUT.",
		Attributes: map[string]schema.Attribute{
			"custom_properties_profile_href": schema.StringAttribute{
				MarkdownDescription: "The `custom_script` link href of the Custom Properties Profile resource (e.g. `smc_custom_properties_profile.example.link.custom_script`).",
				Required:            true,
			},
			"script_path": schema.StringAttribute{
				MarkdownDescription: "Path to the local file to upload.",
				Required:            true,
			},
			"field_name": schema.StringAttribute{
				MarkdownDescription: "Multipart form field name expected by the SMC API. Defaults to `custom_script`.",
				Optional:            true,
			},
		},
	}
}

func (r *SmcUploadScriptAction) Configure(
	ctx context.Context, req action.ConfigureRequest,
	resp *action.ConfigureResponse) {

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

	r.Client = client
}

func (a *SmcUploadScriptAction) Invoke(
	ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config SmcUploadScriptActionModel
	diag := &resp.Diagnostics

	diag.Append(req.Config.Get(ctx, &config)...)
	if diag.HasError() {
		return
	}

	href := config.CustomPropertiesProfileHref.ValueString()
	if href == "" {
		diag.AddError("Missing Required Attribute", "The 'href' attribute must not be empty.")
		return
	}
	if !strings.HasPrefix(href, "http") {
		diag.AddError("Invalid Href", "The 'href' attribute must be a valid URL starting with http or https.")
		return
	}
	if a.Client == nil {
		diag.AddError("SMC Client Not Configured", "The SMC client was not properly configured.")
		return
	}

	fieldName := "custom_script"
	if !config.FieldName.IsNull() && config.FieldName.ValueString() != "" {
		fieldName = config.FieldName.ValueString()
	}

	scriptPath := config.ScriptPath.ValueString()
	fileContent, err := os.ReadFile(scriptPath)
	if err != nil {
		diag.AddError(
			"Error Reading Script File",
			fmt.Sprintf("Could not read file '%s': %s", scriptPath, err.Error()),
		)
		return
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile(fieldName, filepath.Base(scriptPath))
	if err != nil {
		diag.AddError(
			"Error Building Multipart Request",
			fmt.Sprintf("Could not create form field '%s': %s", fieldName, err.Error()),
		)
		return
	}
	if _, err = part.Write(fileContent); err != nil {
		diag.AddError(
			"Error Building Multipart Request",
			fmt.Sprintf("Could not write file content for field '%s': %s", fieldName, err.Error()),
		)
		return
	}
	if err = writer.Close(); err != nil {
		diag.AddError(
			"Error Building Multipart Request",
			fmt.Sprintf("Could not finalize multipart body: %s", err.Error()),
		)
		return
	}

	tflog.Debug(ctx, "uploading script file", map[string]any{
		"href":       href,
		"field_name": fieldName,
		"file":       filepath.Base(scriptPath),
	})

	respData, err := a.Client.DoRequest(smc.Options{
		Method:  "PUT",
		URL:     href,
		Body:    buf.Bytes(),
		Headers: map[string]string{"Content-Type": writer.FormDataContentType()},
	})
	if err != nil {
		diag.AddError(
			"Error Uploading Script",
			fmt.Sprintf("Could not upload script to '%s': %s", href, err.Error()),
		)
		return
	}

	if respData.StatusCode != 200 && respData.StatusCode != 201 && respData.StatusCode != 202 {
		diag.AddError(
			"Error Uploading Script",
			fmt.Sprintf("Unexpected status %d from '%s': %s",
				respData.StatusCode, href, string(respData.Body)),
		)
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Script '%s' uploaded successfully to '%s' (status %d)",
			filepath.Base(scriptPath), href, respData.StatusCode),
	})
}
