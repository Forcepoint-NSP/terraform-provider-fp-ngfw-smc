---
page_title: "smc_autonomous_system"
subcategory: "routing"
description: |-
  This represents the Autonomous System for Dynamic Routing Firewall functionality. The BGP Version is 4/4+. It includes the AS Number in decimal notation, which is used to identify the autonomous system in BGP routing.
---

# smc_autonomous_system (Resource)

This represents the Autonomous System for Dynamic Routing Firewall functionality. The BGP Version is 4/4+. It includes the AS Number in decimal notation, which is used to identify the autonomous system in BGP routing.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.741.0/examples/engines/dynamic_routing/BGP/external_bgp_peers) for a minimal example

This example creates an Autonomous System (AS) object to be used in BGP configuration.

```hcl
resource "smc_autonomous_system" "autonomous_system" {
  as_number = "578"
  name = "tf_autonomous_system"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `as_number` (String) AS Number value in decimal notation. Dotted notation cannot be used here.
- `comment` (String) An optional comment for the element. This field is not required.
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
