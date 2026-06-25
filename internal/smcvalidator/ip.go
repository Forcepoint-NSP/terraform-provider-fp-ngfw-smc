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

package smcvalidator

import (
	"context"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// normalizeIPv4Octets strips leading zeros from each decimal octet of an IPv4
// address string so that e.g. "192.168.100.00" becomes "192.168.100.0".
// Go's net.ParseIP rejects leading zeros since Go 1.17; this normalisation
// lets validators accept the common user typo before handing off to ParseIP.
// Strings that are not four-octet dotted-decimal are returned unchanged.
func normalizeIPv4Octets(s string) string {
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return s
	}
	for i, p := range parts {
		stripped := strings.TrimLeft(p, "0")
		if stripped == "" {
			stripped = "0"
		}
		parts[i] = stripped
	}
	return strings.Join(parts, ".")
}

func IPv4() validator.String { return ipv4Validator{} }

type ipv4Validator struct{}

func (v ipv4Validator) Description(_ context.Context) string {
	return "value must be a valid IPv4 address"
}
func (v ipv4Validator) MarkdownDescription(_ context.Context) string {
	return "value must be a valid IPv4 address"
}
func (v ipv4Validator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	ip := net.ParseIP(normalizeIPv4Octets(val))
	if ip == nil || ip.To4() == nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid IPv4 Address",
			"Expected a valid IPv4 address, got: "+val)
	}
}

func IPv6() validator.String { return ipv6Validator{} }

type ipv6Validator struct{}

func (v ipv6Validator) Description(_ context.Context) string {
	return "value must be a valid IPv6 address"
}
func (v ipv6Validator) MarkdownDescription(_ context.Context) string {
	return "value must be a valid IPv6 address"
}
func (v ipv6Validator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	ip := net.ParseIP(val)
	if ip == nil || ip.To4() != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid IPv6 Address",
			"Expected a valid IPv6 address, got: "+val)
	}
}

func IPv4Range() validator.String { return ipv4RangeValidator{} }

type ipv4RangeValidator struct{}

func (v ipv4RangeValidator) Description(_ context.Context) string {
	return "value must be an IPv4 address range (start-end)"
}
func (v ipv4RangeValidator) MarkdownDescription(_ context.Context) string {
	return "value must be an IPv4 address range (start-end)"
}
func (v ipv4RangeValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	if !isIPv4Range(val) {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid IPv4 Range",
			"Expected an IPv4 range in start-end format (e.g. 192.168.1.1-192.168.1.254), got: "+val)
	}
}

func IPv6Range() validator.String { return ipv6RangeValidator{} }

type ipv6RangeValidator struct{}

func (v ipv6RangeValidator) Description(_ context.Context) string {
	return "value must be an IPv6 address range (start-end)"
}
func (v ipv6RangeValidator) MarkdownDescription(_ context.Context) string {
	return "value must be an IPv6 address range (start-end)"
}
func (v ipv6RangeValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	if !isIPv6Range(val) {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid IPv6 Range",
			"Expected an IPv6 range in start-end format (e.g. 2001:db8::1-2001:db8::ff), got: "+val)
	}
}

func IPRange() validator.String { return ipRangeValidator{} }

type ipRangeValidator struct{}

func (v ipRangeValidator) Description(_ context.Context) string {
	return "value must be an IP address range (start-end), IPv4 or IPv6"
}
func (v ipRangeValidator) MarkdownDescription(_ context.Context) string {
	return "value must be an IP address range (start-end), IPv4 or IPv6"
}
func (v ipRangeValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	if !isIPv4Range(val) && !isIPv6Range(val) {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid IP Range",
			"Expected an IP range in start-end format (IPv4 or IPv6), got: "+val)
	}
}

func isIPv4Range(val string) bool {
	parts := strings.SplitN(val, "-", 2)
	if len(parts) != 2 {
		return false
	}
	start := net.ParseIP(normalizeIPv4Octets(strings.TrimSpace(parts[0])))
	end := net.ParseIP(normalizeIPv4Octets(strings.TrimSpace(parts[1])))
	return start != nil && start.To4() != nil && end != nil && end.To4() != nil
}

func isIPv6Range(val string) bool {
	// Use LastIndex rather than SplitN so the split point is unambiguous: IPv6 addresses
	// never contain '-', so the last '-' in the string is always the range separator.
	idx := strings.LastIndex(val, "-")
	if idx <= 0 {
		return false
	}
	startStr := strings.TrimSpace(val[:idx])
	endStr := strings.TrimSpace(val[idx+1:])
	start := net.ParseIP(startStr)
	end := net.ParseIP(endStr)
	return start != nil && start.To4() == nil && end != nil && end.To4() == nil
}
