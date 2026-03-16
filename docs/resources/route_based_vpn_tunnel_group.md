---
page_title: "smc_rbvpn_tunnel_group"
subcategory: "vpn"
description: |-
  This represents a Route Based VPN Tunnel Group. It is used to group tunnels in the home/monitoring view. It contains a reference to the Link Usage Profile that defines the link usage for the tunnels in this group.
---

# smc_rbvpn_tunnel_group (Resource)

This represents a Route Based VPN Tunnel Group. It is used to group tunnels in the home/monitoring view. It contains a reference to the Link Usage Profile that defines the link usage for the tunnels in this group.

## Examples

- [RBVPN Tunnel Group Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/sdwan/route_based_vpn/vpn_mode/main.tf)

Defines a group of RBVPN tunnels for flexible policy control.

```hcl
resource "smc_rbvpn_tunnel_group" "tf_tunnels" {
  comment = var.resource_comment
  name    = "tf_rbvpn_tunnels"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `link_usage_profile` (String) This represents a Link Usage Profile, which is used to manage link usage settings in a multi-link environment. It includes configurations for link balancing, packet duplication, and forward erasure correction.
- `name` (String) Name of the object.


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
