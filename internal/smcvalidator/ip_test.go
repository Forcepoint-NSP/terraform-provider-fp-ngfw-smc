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

package smcvalidator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terraform-providers/terraform-provider-smc/internal/smcvalidator"
)

func validateString(v validator.String, val string) validator.StringResponse {
	req := validator.StringRequest{
		Path:        path.Root("test"),
		ConfigValue: types.StringValue(val),
	}
	var resp validator.StringResponse
	v.ValidateString(context.Background(), req, &resp)
	return resp
}

// IPv4

func TestIPv4_valid(t *testing.T) {
	for _, val := range []string{
		"192.168.1.1", "10.0.0.1", "0.0.0.0", "255.255.255.255",
		// leading zeros in octets must be accepted
		"192.168.1.01", "010.0.0.1", "192.168.100.00",
	} {
		resp := validateString(smcvalidator.IPv4(), val)
		if resp.Diagnostics.HasError() {
			t.Errorf("IPv4(%q): expected no error, got: %s", val, resp.Diagnostics)
		}
	}
}

func TestIPv4_invalid(t *testing.T) {
	for _, val := range []string{"not-an-ip", "2001:db8::1", "999.0.0.1", "192.168.1"} {
		resp := validateString(smcvalidator.IPv4(), val)
		if !resp.Diagnostics.HasError() {
			t.Errorf("IPv4(%q): expected error but got none", val)
		}
	}
}

func TestIPv4_null_unknown_skipped(t *testing.T) {
	v := smcvalidator.IPv4()
	for _, cfg := range []types.String{types.StringNull(), types.StringUnknown()} {
		req := validator.StringRequest{Path: path.Root("x"), ConfigValue: cfg}
		var resp validator.StringResponse
		v.ValidateString(context.Background(), req, &resp)
		if resp.Diagnostics.HasError() {
			t.Errorf("expected null/unknown to be skipped, got error")
		}
	}
}

// IPv6

func TestIPv6_valid(t *testing.T) {
	for _, val := range []string{"2001:db8::1", "::1", "fe80::1", "2001:0db8:0000:0000:0000:0000:0000:0001"} {
		resp := validateString(smcvalidator.IPv6(), val)
		if resp.Diagnostics.HasError() {
			t.Errorf("IPv6(%q): expected no error, got: %s", val, resp.Diagnostics)
		}
	}
}

func TestIPv6_invalid(t *testing.T) {
	for _, val := range []string{"192.168.1.1", "not-an-ip", "gggg::1"} {
		resp := validateString(smcvalidator.IPv6(), val)
		if !resp.Diagnostics.HasError() {
			t.Errorf("IPv6(%q): expected error but got none", val)
		}
	}
}

// IPv4Range

func TestIPv4Range_valid(t *testing.T) {
	for _, val := range []string{
		"192.168.1.1-192.168.1.254", "10.0.0.1-10.0.0.255",
		// leading zeros in octets must be accepted
		"192.168.1.01-192.168.1.254", "10.0.0.1-10.0.0.0255",
	} {
		resp := validateString(smcvalidator.IPv4Range(), val)
		if resp.Diagnostics.HasError() {
			t.Errorf("IPv4Range(%q): expected no error, got: %s", val, resp.Diagnostics)
		}
	}
}

func TestIPv4Range_invalid(t *testing.T) {
	for _, val := range []string{"192.168.1.1", "192.168.1.1-not-ip", "2001:db8::1-2001:db8::ff", "192.168.1.1-"} {
		resp := validateString(smcvalidator.IPv4Range(), val)
		if !resp.Diagnostics.HasError() {
			t.Errorf("IPv4Range(%q): expected error but got none", val)
		}
	}
}

// IPv6Range

func TestIPv6Range_valid(t *testing.T) {
	for _, val := range []string{"2001:db8::1-2001:db8::ff", "::1-::ffff"} {
		resp := validateString(smcvalidator.IPv6Range(), val)
		if resp.Diagnostics.HasError() {
			t.Errorf("IPv6Range(%q): expected no error, got: %s", val, resp.Diagnostics)
		}
	}
}

func TestIPv6Range_invalid(t *testing.T) {
	for _, val := range []string{"192.168.1.1-192.168.1.254", "not-a-range", "2001:db8::1"} {
		resp := validateString(smcvalidator.IPv6Range(), val)
		if !resp.Diagnostics.HasError() {
			t.Errorf("IPv6Range(%q): expected error but got none", val)
		}
	}
}

// IPRange

func TestIPRange_valid(t *testing.T) {
	for _, val := range []string{"192.168.1.1-192.168.1.254", "2001:db8::1-2001:db8::ff"} {
		resp := validateString(smcvalidator.IPRange(), val)
		if resp.Diagnostics.HasError() {
			t.Errorf("IPRange(%q): expected no error, got: %s", val, resp.Diagnostics)
		}
	}
}

func TestIPRange_invalid(t *testing.T) {
	for _, val := range []string{"192.168.1.1", "not-a-range", "192.168.1.1-2001:db8::1"} {
		resp := validateString(smcvalidator.IPRange(), val)
		if !resp.Diagnostics.HasError() {
			t.Errorf("IPRange(%q): expected error but got none", val)
		}
	}
}
