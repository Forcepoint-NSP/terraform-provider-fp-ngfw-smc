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
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/common"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

var (
	_ action.ActionWithConfigure = (*ImportVpnCertificateAction)(nil)
)

type ImportVpnCertificateActionModel struct {
	GatewayHref types.String `tfsdk:"gateway_href"`
	CertFile    types.String `tfsdk:"cert_file"`
	CsrHrefFile types.String `tfsdk:"csr_href_file"`
}

func NewImportVpnCertificateAction() action.Action {
	return &ImportVpnCertificateAction{}
}

type ImportVpnCertificateAction struct {
	Client *smc.SmcClient
}

func (a *ImportVpnCertificateAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_import_vpn_certificate"
}

func (a *ImportVpnCertificateAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Import a certificate — signed by an external CA — into the internal gateway's " +
			"pending certificate request created by `smc_generate_vpn_certificate`.",
		Attributes: map[string]schema.Attribute{
			"gateway_href": schema.StringAttribute{
				MarkdownDescription: "Internal gateway element URL (e.g. from `data.smc_sub_href.internal_gateway.id`).",
				Required:            true,
			},
			"cert_file": schema.StringAttribute{
				MarkdownDescription: "Path to the signed PEM certificate to import.",
				Required:            true,
			},
			"csr_href_file": schema.StringAttribute{
				MarkdownDescription: "Path to the `<csr_output_file>.href` file written by `smc_generate_vpn_certificate` " +
					"alongside the CSR (e.g. `smc_generate_vpn_certificate.<label>.config.csr_output_file` + `.href`). " +
					"Set this to import into the exact certificate request that CSR came from. " +
					"When omitted, the action imports into the most recently created pending request for `gateway_href`.",
				Optional: true,
			},
		},
	}
}

func (a *ImportVpnCertificateAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *ImportVpnCertificateAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ImportVpnCertificateActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	gatewayHref := config.GatewayHref.ValueString()
	if !strings.HasPrefix(gatewayHref, "http") {
		resp.Diagnostics.AddError("Invalid gateway_href", "The 'gateway_href' attribute must be a valid URL.")
		return
	}
	if a.Client == nil {
		resp.Diagnostics.AddError("SMC Client Not Configured", "The SMC client was not properly configured.")
		return
	}

	certPEM, err := os.ReadFile(config.CertFile.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading certificate file",
			fmt.Sprintf("Could not read '%s': %s", config.CertFile.ValueString(), err.Error()))
		return
	}

	reader := &smc.GenericCRUDConfig{Client: a.Client}

	var importHref string
	if !config.CsrHrefFile.IsNull() && config.CsrHrefFile.ValueString() != "" {
		// Read the CSR element URL written alongside the PEM by smc_generate_vpn_certificate.
		csrHrefBytes, readErr := os.ReadFile(config.CsrHrefFile.ValueString())
		if readErr != nil {
			resp.Diagnostics.AddError("Error reading CSR href file",
				fmt.Sprintf("Could not read '%s': %s\nEnsure smc_generate_vpn_certificate ran first.", config.CsrHrefFile.ValueString(), readErr.Error()))
			return
		}
		csrHref := strings.TrimSpace(string(csrHrefBytes))
		var navErr error
		importHref, navErr = common.ResolveRef2(ctx, reader, csrHref,
			"link[?rel=='certificate_import'] | [0].href")
		if navErr != nil {
			resp.Diagnostics.AddError("Error resolving certificate_import link",
				fmt.Sprintf("Could not navigate to certificate_import from CSR %s: %s", csrHref, navErr.Error()))
			return
		}
	} else {
		// Fall back to gateway collection traversal, like smc-python's
		// engine.vpn.gateway_certificate_request property. The collection GET
		// returns {"result": [...]}; pick the last (most recent) request.
		var navErr error
		importHref, navErr = common.ResolveRef2(ctx, reader, gatewayHref,
			"gateway_certificate_request/result[-1].href/link[?rel=='certificate_import'] | [0].href")
		if navErr != nil {
			resp.Diagnostics.AddError("Error resolving certificate_import link",
				fmt.Sprintf("Could not navigate to certificate_import from %s: %s\n"+
					"Ensure generate_csr was invoked before running this action.", gatewayHref, navErr.Error()))
			return
		}
	}

	tflog.Debug(ctx, "importing VPN certificate", map[string]any{
		"gateway_href": gatewayHref,
		"import_href":  importHref,
		"cert_file":    config.CertFile.ValueString(),
	})

	body, err := json.Marshal(map[string]string{"certificate": string(certPEM)})
	if err != nil {
		resp.Diagnostics.AddError("Error marshalling certificate", err.Error())
		return
	}

	respData, err := a.Client.DoRequest(smc.Options{
		Method:  "POST",
		URL:     importHref,
		Body:    body,
		Context: ctx,
	})
	if err != nil {
		resp.Diagnostics.AddError("Error importing certificate",
			fmt.Sprintf("POST to %s failed: %s", importHref, err.Error()))
		return
	}
	if respData.StatusCode != 200 && respData.StatusCode != 201 && respData.StatusCode != 204 {
		resp.Diagnostics.AddError("Error importing certificate",
			fmt.Sprintf("POST to %s returned status %d: %s", importHref, respData.StatusCode, string(respData.Body)))
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("certificate imported successfully from '%s'", config.CertFile.ValueString()),
	})
}
