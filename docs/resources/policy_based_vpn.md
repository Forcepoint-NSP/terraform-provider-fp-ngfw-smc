---
page_title: "smc_vpn"
subcategory: "vpn"
description: |-
  This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.
---

# smc_vpn (Resource)

This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.

## Examples

- [fw_policy_vpn example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/policies/fw_policy_vpn/main.tf): Demonstrates configuring a firewall policy with VPN rules using the SMC provider.

```hcl
resource "smc_vpn_profile" "tf_keepalive_vpn_profile" {
  name       = "tf_keepalive_vpn_profile"
  keep_alive = true
}

# Create VPN with a specific profile to keep alive tunnels
resource "smc_vpn" "policy_vpn" {
  name        = "tf_vpn"
  nat         = false
  vpn_profile = smc_vpn_profile.tf_keepalive_vpn_profile.id
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `dscp_qos_policy` (String) This represents a QoS Policy, which is used for Bandwidth Management and Traffic Prioritization based on QoS Classes or DSCP Matches.
- `link_usage_profile` (String) This represents a Link Usage Profile, which is used to manage link usage settings in a multi-link environment. It includes configurations for link balancing, packet duplication, and forward erasure correction.
- `mobile_vpn_topology_mode` (String) Indicates the mobile VPN topology mode, which defines the type of gateway communicating with the VPN client.
- `name` (String) Name of the object.
- `nat` (Boolean) Indicates whether NAT rules are applied for this VPN.
- `vpn_profile` (String) This represents a VPN Profile. It contains settings for IKE and IPsec lifetimes, keep-alive options, certificate authorities, and authentication methods.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
