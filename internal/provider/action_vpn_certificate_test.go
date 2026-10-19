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

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	fwvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/terraform-providers/terraform-provider-smc/internal/common"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

// readerMock implements smc.IResourceReader for navigation tests.
type readerMock struct{ mock.Mock }

func (m *readerMock) ReadResourceByHref(href string) (*smc.ResponseData, error) {
	args := m.Called(href)
	return args.Get(0).(*smc.ResponseData), args.Error(1)
}

const (
	csrHref        = "http://smc:18082/7.6/elements/single_fw/100/internal_gateway/200/gateway_certificate_request/11"
	certExportHref = "http://smc:18082/7.6/elements/single_fw/100/internal_gateway/200/gateway_certificate_request/11/certificate_export"
	certImportHref = "http://smc:18082/7.6/elements/single_fw/100/internal_gateway/200/gateway_certificate_request/11/certificate_import"
)

// csrRespWithExport is the CSR element returned by the Location URL from generate_certificate POST.
const csrRespWithExport = `{
	"name": "new-csr",
	"key": 11,
	"link": [
		{"href": "` + csrHref + `", "rel": "self", "type": "gateway_certificate_request"},
		{"href": "` + certExportHref + `", "rel": "certificate_export"}
	]
}`

// csrRespWithImport is the same CSR element with certificate_import link visible.
const csrRespWithImport = `{
	"name": "new-csr",
	"key": 11,
	"link": [
		{"href": "` + csrHref + `", "rel": "self", "type": "gateway_certificate_request"},
		{"href": "` + certImportHref + `", "rel": "certificate_import"}
	]
}`

func resp(body string) *smc.ResponseData {
	return &smc.ResponseData{Body: []byte(body), StatusCode: 200}
}

// TestCsrExportNavigation verifies the generate action's Location-header approach:
// navigate from the CSR element URL (from POST Location) to its certificate_export link.
func TestCsrExportNavigation(t *testing.T) {
	r := new(readerMock)
	r.On("ReadResourceByHref", csrHref).Return(resp(csrRespWithExport), nil)

	got, err := common.ResolveRef2(context.Background(), r,
		csrHref,
		"link[?rel=='certificate_export'] | [0].href")

	assert.NoError(t, err)
	assert.Equal(t, certExportHref, got)
	r.AssertExpectations(t)
}

// TestCsrImportNavigation verifies the import action's csr_href_file approach:
// navigate from the CSR element URL (read from companion .href file) to its certificate_import link.
func TestCsrImportNavigation(t *testing.T) {
	r := new(readerMock)
	r.On("ReadResourceByHref", csrHref).Return(resp(csrRespWithImport), nil)

	got, err := common.ResolveRef2(context.Background(), r,
		csrHref,
		"link[?rel=='certificate_import'] | [0].href")

	assert.NoError(t, err)
	assert.Equal(t, certImportHref, got)
	r.AssertExpectations(t)
}

// TestCsrImportFallbackNavigation verifies the import action's fallback when no
// csr_href_file is given: gateway → gateway_certificate_request collection
// ({"result": [...]}, as unwrapped by smc-python's SMCResult) → last item →
// certificate_import link.
func TestCsrImportFallbackNavigation(t *testing.T) {
	const (
		gwHref         = "http://smc:18082/7.6/elements/single_fw/100/internal_gateway/200"
		gwcrCollection = gwHref + "/gateway_certificate_request"
		oldCsrHref     = gwcrCollection + "/10"
	)
	gatewayResp := `{
		"name": "myfw_gateway",
		"link": [
			{"href": "` + gwHref + `", "rel": "self", "type": "internal_gateway"},
			{"href": "` + gwcrCollection + `", "rel": "gateway_certificate_request"}
		]
	}`
	collectionResp := `{"result": [
		{"href": "` + oldCsrHref + `", "name": "old-csr", "type": "gateway_certificate_request"},
		{"href": "` + csrHref + `", "name": "new-csr", "type": "gateway_certificate_request"}
	]}`

	r := new(readerMock)
	r.On("ReadResourceByHref", gwHref).Return(resp(gatewayResp), nil)
	r.On("ReadResourceByHref", gwcrCollection).Return(resp(collectionResp), nil)
	r.On("ReadResourceByHref", csrHref).Return(resp(csrRespWithImport), nil)

	got, err := common.ResolveRef2(context.Background(), r,
		gwHref,
		"gateway_certificate_request/result[-1].href/link[?rel=='certificate_import'] | [0].href")

	assert.NoError(t, err)
	assert.Equal(t, certImportHref, got)
	r.AssertExpectations(t)
}

// TestGenerateCertNavigation verifies the generate action's gateway_href approach:
// navigate from the internal gateway element to its generate_certificate link.
func TestGenerateCertNavigation(t *testing.T) {
	const (
		gwHref          = "http://smc:18082/7.6/elements/single_fw/100/internal_gateway/200"
		generateCertURL = gwHref + "/generate_certificate"
	)
	gatewayResp := `{
		"name": "myfw_gateway",
		"link": [
			{"href": "` + gwHref + `", "rel": "self", "type": "internal_gateway"},
			{"href": "` + generateCertURL + `", "rel": "generate_certificate"}
		]
	}`

	r := new(readerMock)
	r.On("ReadResourceByHref", gwHref).Return(resp(gatewayResp), nil)

	got, err := common.ResolveRef2(context.Background(), r,
		gwHref,
		"link[?rel=='generate_certificate'] | [0].href")

	assert.NoError(t, err)
	assert.Equal(t, generateCertURL, got)
	r.AssertExpectations(t)
}

// TestStatusTimeoutValidators verifies status_timeout/status_poll_interval reject
// non-positive values, since they flow into time.NewTimer/NewTicker which panic
// on a zero or negative duration.
func TestStatusTimeoutValidators(t *testing.T) {
	a := NewGenerateVpnCertificateAction()
	var resp action.SchemaResponse
	a.Schema(context.Background(), action.SchemaRequest{}, &resp)

	for _, attrName := range []string{"status_timeout", "status_poll_interval"} {
		attr, ok := resp.Schema.Attributes[attrName].(schema.Int64Attribute)
		if !ok {
			t.Fatalf("%s: expected schema.Int64Attribute", attrName)
		}
		if len(attr.Validators) == 0 {
			t.Fatalf("%s: expected at least one validator", attrName)
		}

		for _, v := range []int64{0, -1} {
			var vResp fwvalidator.Int64Response
			for _, validator := range attr.Validators {
				validator.ValidateInt64(context.Background(), fwvalidator.Int64Request{
					ConfigValue: types.Int64Value(v),
				}, &vResp)
			}
			assert.True(t, vResp.Diagnostics.HasError(), "%s: value %d should be rejected", attrName, v)
		}

		var vResp fwvalidator.Int64Response
		for _, validator := range attr.Validators {
			validator.ValidateInt64(context.Background(), fwvalidator.Int64Request{
				ConfigValue: types.Int64Value(5),
			}, &vResp)
		}
		assert.False(t, vResp.Diagnostics.HasError(), "%s: positive value should be accepted", attrName)
	}
}

// TestAlgorithmValidators verifies public_key_algorithm/signature_algorithm reject
// values outside the SMC API's enum (see Generate_VPN_Certificate__Certificate_Request_Arguments
// in the OpenAPI spec).
func TestAlgorithmValidators(t *testing.T) {
	a := NewGenerateVpnCertificateAction()
	var resp action.SchemaResponse
	a.Schema(context.Background(), action.SchemaRequest{}, &resp)

	cases := []struct {
		attrName string
		valid    string
		invalid  string
	}{
		{"public_key_algorithm", "ecdsa", "ecsda"},
		{"signature_algorithm", "rsa_sha_512", "rsa_sha_1024"},
	}
	for _, c := range cases {
		attr, ok := resp.Schema.Attributes[c.attrName].(schema.StringAttribute)
		if !ok {
			t.Fatalf("%s: expected schema.StringAttribute", c.attrName)
		}
		if len(attr.Validators) == 0 {
			t.Fatalf("%s: expected at least one validator", c.attrName)
		}

		var invalidResp fwvalidator.StringResponse
		for _, validator := range attr.Validators {
			validator.ValidateString(context.Background(), fwvalidator.StringRequest{
				ConfigValue: types.StringValue(c.invalid),
			}, &invalidResp)
		}
		assert.True(t, invalidResp.Diagnostics.HasError(), "%s: value %q should be rejected", c.attrName, c.invalid)

		var validResp fwvalidator.StringResponse
		for _, validator := range attr.Validators {
			validator.ValidateString(context.Background(), fwvalidator.StringRequest{
				ConfigValue: types.StringValue(c.valid),
			}, &validResp)
		}
		assert.False(t, validResp.Diagnostics.HasError(), "%s: value %q should be accepted", c.attrName, c.valid)
	}
}
