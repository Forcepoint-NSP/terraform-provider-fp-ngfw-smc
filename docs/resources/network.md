---
page_title: "smc_network"
subcategory: "network_elements"
description: |-
  This represents a Network, which is a Network Element that represents a (sub)network of computers. It includes attributes for IPv4 and IPv6 networks, broadcast address, and validation patterns.
---

# smc_network (Resource)

This represents a Network, which is a Network Element that represents a (sub)network of computers. It includes attributes for IPv4 and IPv6 networks, broadcast address, and validation patterns.

## Examples

- [getting_started/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/getting_started/main.tf): Defines a basic network element for a test/lab environment.

This snippet shows the use of `smc_network` to define a simple IPv4 network. This network can later be associated with firewall interfaces or VPN sites.

```hcl
resource "smc_network" "tf_sample_network" {
  name         = "tf_sample_network"
  comment      = var.resource_comment
  ipv4_network = local.tf_network
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `broadcast` (Boolean) Indicates whether the broadcast address and network address are included in the definition.
- `comment` (String) An optional comment for the element. This field is not required.
- `ipv4_network` (String) A valid IPv4 address and netmask in CIDR notation.
- `ipv6_network` (String) A valid IPv6 address and prefix length in CIDR notation.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
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
