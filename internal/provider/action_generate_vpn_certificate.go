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
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/common"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

var (
	_ action.ActionWithConfigure = (*GenerateVpnCertificateAction)(nil)
)

type GenerateVpnCertificateActionModel struct {
	GatewayHref              types.String `tfsdk:"gateway_href"`
	CommonName               types.String `tfsdk:"common_name"`
	Organization             types.String `tfsdk:"organization"`
	OrganizationUnit         types.String `tfsdk:"organization_unit"`
	Country                  types.String `tfsdk:"country"`
	Locality                 types.String `tfsdk:"locality"`
	StateProvince            types.String `tfsdk:"state_province"`
	PublicKeyAlgorithm       types.String `tfsdk:"public_key_algorithm"`
	SignatureAlgorithm       types.String `tfsdk:"signature_algorithm"`
	PublicKeyLength          types.Int64  `tfsdk:"public_key_length"`
	CertificateAuthorityHref types.String `tfsdk:"certificate_authority_href"`
	CsrOutputFile            types.String `tfsdk:"csr_output_file"`
	StatusTimeout            types.Int64  `tfsdk:"status_timeout"`
	StatusPollInterval       types.Int64  `tfsdk:"status_poll_interval"`
}

func NewGenerateVpnCertificateAction() action.Action {
	return &GenerateVpnCertificateAction{}
}

type GenerateVpnCertificateAction struct {
	Client *smc.SmcClient
}

func (a *GenerateVpnCertificateAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_generate_vpn_certificate"
}

func (a *GenerateVpnCertificateAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Generate a VPN certificate (or certificate request) for an internal gateway. " +
			"When `certificate_authority_href` is set, the request is signed by that internal certificate authority; " +
			"otherwise only a certificate request is created. The engine must have completed initial contact with " +
			"the SMC before a certificate can be generated.",
		Attributes: map[string]schema.Attribute{
			"gateway_href": schema.StringAttribute{
				MarkdownDescription: "Internal gateway element URL (e.g. from `data.smc_sub_href.internal_gateway.id`). " +
					"The action follows the gateway's `generate_certificate` link automatically.",
				Required: true,
			},
			"common_name": schema.StringAttribute{
				MarkdownDescription: "The common name, used for subject name creation.",
				Required:            true,
			},
			"organization": schema.StringAttribute{
				MarkdownDescription: "The name of the organization, used for subject name creation.",
				Required:            true,
			},
			"organization_unit": schema.StringAttribute{
				MarkdownDescription: "The organization unit, used for subject name creation.",
				Optional:            true,
			},
			"country": schema.StringAttribute{
				MarkdownDescription: "The country code (2 characters), used for subject name creation.",
				Optional:            true,
			},
			"locality": schema.StringAttribute{
				MarkdownDescription: "The locality, used for subject name creation.",
				Optional:            true,
			},
			"state_province": schema.StringAttribute{
				MarkdownDescription: "The state or province, used for subject name creation.",
				Optional:            true,
			},
			"public_key_algorithm": schema.StringAttribute{
				MarkdownDescription: "The public key algorithm used for the certificate (dsa, rsa or ecdsa). Defaults to rsa.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("dsa", "rsa", "ecdsa")},
			},
			"signature_algorithm": schema.StringAttribute{
				MarkdownDescription: "The signature algorithm used for the certificate (e.g. rsa_sha_256, rsa_sha_512, ecdsa_sha_512). Defaults to rsa_sha_512.",
				Optional:            true,
				Validators: []validator.String{stringvalidator.OneOf(
					"dsa_sha_1", "dsa_sha_224", "dsa_rsa_256",
					"rsa_md5", "rsa_sha_1", "rsa_sha_256", "rsa_sha_384", "rsa_sha_512",
					"ecdsa_sha_1", "ecdsa_sha_256", "ecdsa_sha_384", "ecdsa_sha_512",
				)},
			},
			"public_key_length": schema.Int64Attribute{
				MarkdownDescription: "The length in bits of the public key. Defaults to 2048.",
				Optional:            true,
			},
			"certificate_authority_href": schema.StringAttribute{
				MarkdownDescription: "Reference to the internal certificate authority used for signing the generated certificate request. " +
					"When omitted, only a certificate request is generated.",
				Optional: true,
			},
			"csr_output_file": schema.StringAttribute{
				MarkdownDescription: "Path to save the generated certificate request (PEM format). " +
					"Only used when `certificate_authority_href` is not set: get the CSR signed by an external " +
					"CA, then pass the signed certificate back to `smc_import_vpn_certificate`'s `cert_file`. " +
					"The action also writes a `<csr_output_file>.href` companion file next to it, which " +
					"`smc_import_vpn_certificate`'s `csr_href_file` attribute can point to.",
				Optional: true,
			},
			"status_timeout": schema.Int64Attribute{
				MarkdownDescription: "Maximum seconds to wait for the engine to come online before generating the certificate " +
					"— the engine must have completed initial contact with the SMC first. " +
					"The action fails if the engine hasn't come online within this time. Defaults to 300.",
				Optional:   true,
				Validators: []validator.Int64{int64validator.AtLeast(1)},
			},
			"status_poll_interval": schema.Int64Attribute{
				MarkdownDescription: "Seconds between engine status polls. Defaults to 10.",
				Optional:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(1)},
			},
		},
	}
}

func (a *GenerateVpnCertificateAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *GenerateVpnCertificateAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config GenerateVpnCertificateActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	gatewayHref := config.GatewayHref.ValueString()
	if !strings.HasPrefix(gatewayHref, "http") {
		resp.Diagnostics.AddError("Invalid gateway_href", "The 'gateway_href' attribute must be a valid URL starting with http or https.")
		return
	}
	if a.Client == nil {
		resp.Diagnostics.AddError("SMC Client Not Configured", "The SMC client was not properly configured.")
		return
	}

	// Follow the gateway's generate_certificate link so users don't have
	// to resolve the operation URL with an extra data source.
	reader := &smc.GenericCRUDConfig{Client: a.Client}
	href, err := common.ResolveRef2(ctx, reader, gatewayHref,
		"link[?rel=='generate_certificate'] | [0].href")
	if err != nil {
		resp.Diagnostics.AddError("Error resolving generate_certificate link",
			fmt.Sprintf("Could not resolve generate_certificate from gateway %s: %s", gatewayHref, err.Error()))
		return
	}

	// Remove leftovers from a previous run as early as possible so nothing
	// downstream (file watchers, signing scripts) can pick up a stale CSR.
	if csrFile := config.CsrOutputFile.ValueString(); csrFile != "" {
		_ = os.Remove(csrFile)
		_ = os.Remove(csrFile + ".href")
	}

	timeoutSec := int64(300)
	if !config.StatusTimeout.IsNull() {
		timeoutSec = config.StatusTimeout.ValueInt64()
	}
	pollSec := int64(10)
	if !config.StatusPollInterval.IsNull() {
		pollSec = config.StatusPollInterval.ValueInt64()
	}

	if err := a.waitForEngineOnline(ctx, resp, href, timeoutSec, pollSec); err != nil {
		resp.Diagnostics.AddError("Engine not ready", err.Error())
		return
	}

	publicKeyAlgorithm := "rsa"
	if !config.PublicKeyAlgorithm.IsNull() && config.PublicKeyAlgorithm.ValueString() != "" {
		publicKeyAlgorithm = config.PublicKeyAlgorithm.ValueString()
	}
	signatureAlgorithm := "rsa_sha_512"
	if !config.SignatureAlgorithm.IsNull() && config.SignatureAlgorithm.ValueString() != "" {
		signatureAlgorithm = config.SignatureAlgorithm.ValueString()
	}
	publicKeyLength := int64(2048)
	if !config.PublicKeyLength.IsNull() {
		publicKeyLength = config.PublicKeyLength.ValueInt64()
	}

	payload := map[string]any{
		"common_name":          config.CommonName.ValueString(),
		"organization":         config.Organization.ValueString(),
		"public_key_algorithm": publicKeyAlgorithm,
		"signature_algorithm":  signatureAlgorithm,
		"public_key_length":    publicKeyLength,
	}
	setOptionalString := func(key string, value types.String) {
		if !value.IsNull() && value.ValueString() != "" {
			payload[key] = value.ValueString()
		}
	}
	setOptionalString("organization_unit", config.OrganizationUnit)
	setOptionalString("country", config.Country)
	setOptionalString("locality", config.Locality)
	setOptionalString("state_province", config.StateProvince)
	setOptionalString("certificate_authority_href", config.CertificateAuthorityHref)

	body, err := json.Marshal(payload)
	if err != nil {
		resp.Diagnostics.AddError("Error marshalling certificate arguments", err.Error())
		return
	}

	tflog.Debug(ctx, "invoking generate vpn certificate", map[string]any{"url": href, "body": string(body)})

	respData, err := a.Client.DoRequest(smc.Options{
		Method:  "POST",
		URL:     href,
		Body:    body,
		Context: ctx,
	})
	if err != nil {
		resp.Diagnostics.AddError("Error generating VPN certificate", fmt.Sprintf("POST to %s failed: %s", href, err.Error()))
		return
	}
	if respData.StatusCode != 200 && respData.StatusCode != 201 && respData.StatusCode != 202 {
		resp.Diagnostics.AddError("Error generating VPN certificate",
			fmt.Sprintf("POST to %s returned status %d: %s", href, respData.StatusCode, string(respData.Body)))
		return
	}

	withCA := !config.CertificateAuthorityHref.IsNull() && config.CertificateAuthorityHref.ValueString() != ""

	if withCA {
		message := "VPN certificate generated successfully"
		if respData.Location != "" {
			message = fmt.Sprintf("%s: %s", message, respData.Location)
		}
		resp.SendProgress(action.InvokeProgressEvent{Message: message})
		return
	}

	// No CA: SMC created a gateway_certificate_request.
	// The POST response carries a Location header pointing directly to the new CSR element.
	resp.SendProgress(action.InvokeProgressEvent{Message: "VPN certificate request generated successfully"})

	csrOutputFile := config.CsrOutputFile.ValueString()
	if config.CsrOutputFile.IsNull() || csrOutputFile == "" {
		return
	}

	csrLocation := respData.Location
	if csrLocation == "" {
		resp.Diagnostics.AddError("CSR location missing",
			fmt.Sprintf("POST to %s did not return a Location header; cannot export CSR.", href))
		return
	}

	exportHref, err := common.ResolveRef2(ctx, reader, csrLocation,
		"link[?rel=='certificate_export'] | [0].href")
	if err != nil {
		resp.Diagnostics.AddError("Error locating certificate_export link",
			fmt.Sprintf("Could not resolve certificate_export from CSR %s: %s", csrLocation, err.Error()))
		return
	}

	tflog.Debug(ctx, "exporting CSR", map[string]any{"export_href": exportHref, "output_file": csrOutputFile})

	exportResp, err := a.Client.DoRequest(smc.Options{
		Method:  "GET",
		URL:     exportHref,
		Context: ctx,
	})
	if err != nil {
		resp.Diagnostics.AddError("Error exporting CSR", fmt.Sprintf("GET %s failed: %s", exportHref, err.Error()))
		return
	}
	if exportResp.StatusCode != 200 {
		resp.Diagnostics.AddError("Error exporting CSR",
			fmt.Sprintf("GET %s returned status %d: %s", exportHref, exportResp.StatusCode, string(exportResp.Body)))
		return
	}

	if err := os.WriteFile(csrOutputFile, exportResp.Body, 0600); err != nil {
		resp.Diagnostics.AddError("Error writing CSR file",
			fmt.Sprintf("Could not write CSR to '%s': %s", csrOutputFile, err.Error()))
		return
	}

	// Write the CSR element URL last, once the PEM is on disk: consumers can
	// treat the presence of a .href file referencing the current gateway as
	// "the CSR next to it is fresh and ready to sign".
	csrHrefFile := csrOutputFile + ".href"
	if err := os.WriteFile(csrHrefFile, []byte(csrLocation), 0600); err != nil {
		resp.Diagnostics.AddError("Error writing CSR href file",
			fmt.Sprintf("Could not write CSR href to '%s': %s", csrHrefFile, err.Error()))
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("CSR exported to '%s'", csrOutputFile),
	})
}

// waitForEngineOnline auto-derives the node status URL from the
// generate_certificate href and polls until connectivity_status == "OK".
// An OK engine has completed initial contact; no policy upload is required.
// If the status URL cannot be resolved (unexpected URL shape) the wait is skipped.
func (a *GenerateVpnCertificateAction) waitForEngineOnline(
	ctx context.Context,
	resp *action.InvokeResponse,
	certHref string,
	timeoutSec int64,
	pollSec int64,
) error {
	parentURL, err := extractParentElementURL(certHref)
	if err != nil {
		tflog.Warn(ctx, "could not derive parent element URL, skipping engine status wait", map[string]any{"err": err.Error()})
		return nil
	}

	statusHref, err := resolveNodeStatusHref(ctx, a.Client, parentURL)
	if err != nil {
		tflog.Warn(ctx, "could not resolve node status URL, skipping engine status wait", map[string]any{"err": err.Error()})
		return nil
	}

	// Check immediately — if already OK skip polling entirely. A failed
	// pre-check (engine just contacted, first status GET is prone to a
	// transient error) falls through into the resilient waitForEngineStatus
	// loop instead of aborting, same as any other transient error there.
	if s, err := fetchEngineNodeStatus(ctx, a.Client, statusHref); err == nil {
		if s.ConnectivityStatus == "OK" {
			return nil
		}
		resp.SendProgress(action.InvokeProgressEvent{
			Message: fmt.Sprintf("[%s] waiting for OK — current: %s (engine: %s)",
				s.Name, s.ConnectivityStatus, s.EngineNodeStatus),
		})
	}

	return waitForEngineStatus(
		ctx,
		a.Client,
		statusHref,
		func(s engineNodeStatusResponse) string { return s.ConnectivityStatus },
		"OK",
		time.Duration(timeoutSec)*time.Second,
		time.Duration(pollSec)*time.Second,
		func(s engineNodeStatusResponse) {
			resp.SendProgress(action.InvokeProgressEvent{
				Message: fmt.Sprintf("[%s] waiting for OK — current: %s (engine: %s)",
					s.Name, s.ConnectivityStatus, s.EngineNodeStatus),
			})
		},
	)
}
