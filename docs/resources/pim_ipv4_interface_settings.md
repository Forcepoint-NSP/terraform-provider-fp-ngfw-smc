---
page_title: "smc_pim_ipv4_interface_settings"
subcategory: "routing"
description: |-
  This represents the PIM IPv4 Interface Settings for Dynamic Routing Firewall functionality.
---

# smc_pim_ipv4_interface_settings (Resource)

This represents the PIM IPv4 Interface Settings for Dynamic Routing Firewall functionality.

## Examples

- [PIM IPv4 Interface Settings Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/PIM/pim_ipv4_interface_settings/main.tf)

Defines PIM-specific multicast settings for a given interface.

```hcl
resource "smc_pim_ipv4_interface_settings" "pim_interface_settings" {
  dr_priority               = 14
  igmp_querier_settings_ref = data.smc_href.igmp_querier_settings.id
  name                      = "tf_pim_interface_settings"
  random_delay              = 10
  zbr                       = "224.0.1.1,239.255.255.250"
  comment                   = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `dr_priority` (Number) The Designated Router (DR) priority for the PIM Interface. This value determines the priority of the router in the PIM network.
- `igmp_querier_settings_ref` (String) This represents the IGMP Querier Settings for Multicast Routing and especially the PIM dynamic routing feature.
- `name` (String) Name of the object.
- `random_delay` (Number) The Random Delay in seconds for the PIM Interface. This value is used to introduce a random delay in the PIM operations.
- `zbr` (String) The Zone Boundary Router (ZBR) group for the PIM Interface. Listed Multicast Groups do not go through the interface.


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
