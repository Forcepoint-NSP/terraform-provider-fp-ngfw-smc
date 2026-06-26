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

func IPv4Network() validator.String { return ipv4NetworkValidator{} }

type ipv4NetworkValidator struct{}

func (v ipv4NetworkValidator) Description(_ context.Context) string {
	return "value must be a valid IPv4 network in CIDR notation (e.g. 192.168.0.0/24)"
}
func (v ipv4NetworkValidator) MarkdownDescription(_ context.Context) string {
	return "value must be a valid IPv4 network in CIDR notation (e.g. 192.168.0.0/24)"
}
func (v ipv4NetworkValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	parts := strings.SplitN(val, "/", 2)
	normalizedHost := normalizeIPv4Octets(parts[0])
	normalizedVal := normalizedHost
	if len(parts) == 2 {
		normalizedVal = normalizedHost + "/" + parts[1]
	}
	ip := net.ParseIP(normalizedHost)
	_, _, err := net.ParseCIDR(normalizedVal)
	if err != nil || ip == nil || ip.To4() == nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid IPv4 Network",
			"Expected an IPv4 CIDR network (e.g. 192.168.0.0/24), got: "+val)
	}
}

func IPv6Network() validator.String { return ipv6NetworkValidator{} }

type ipv6NetworkValidator struct{}

func (v ipv6NetworkValidator) Description(_ context.Context) string {
	return "value must be a valid IPv6 network in CIDR notation (e.g. 2001:db8::/32)"
}
func (v ipv6NetworkValidator) MarkdownDescription(_ context.Context) string {
	return "value must be a valid IPv6 network in CIDR notation (e.g. 2001:db8::/32)"
}
func (v ipv6NetworkValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	hostPart := strings.SplitN(val, "/", 2)[0]
	ip := net.ParseIP(hostPart)
	_, _, err := net.ParseCIDR(val)
	if err != nil || ip == nil || ip.To4() != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid IPv6 Network",
			"Expected an IPv6 CIDR network (e.g. 2001:db8::/32), got: "+val)
	}
}
