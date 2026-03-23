---
page_title: "smc_vpn_site"
subcategory: "vpn"
description: |-
  This represents a VPN Site, which is a collection of network elements that are part of a VPN. It can be associated with a gateway and has references to VPNs.
---

# smc_vpn_site (Sub-resource)

This represents a VPN Site, which is a collection of network elements that are part of a VPN. It can be associated with a gateway and has references to VPNs.

## Examples

- [satellite_gateway_node/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.741.0/examples/sdwan/policy_based_vpn/satellite_gateway_node/main.tf): Creates a VPN site for a firewall's internal gateway in SD-WAN.

This example demonstrates how to define an `smc_vpn_site` for associating a network element (such as a LAN or subnet) to a firewall's internal gateway in an SD-WAN VPN design.

```hcl
resource "smc_vpn_site" "site_fw1" {
  from_ref     = format("%s#%s", smc_single_fw.tf_single_fw["fw1"].link.internal_gateway, "*/vpn_site")
  name         = "site_fw1"
  site_element = [smc_network.nw1.id]
  # gateway    = data.smc_sub_href.internal_gateway["fw1"].id
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `automatic` (Boolean) Indicates whether this site is automatically created or manually added. If true, the site is automatically created and cannot have VPN references set.
- `comment` (String) An optional comment for the element. This field is not required.
- `gateway` (String) This represents a VPN Gateway. A VPN Gateway is a network element that can be used to establish a VPN connection with other gateways.
- `name` (String) Name of the object.
- `site_element` (List of String) 


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
