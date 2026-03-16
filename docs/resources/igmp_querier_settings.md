---
page_title: "smc_igmp_querier_settings"
subcategory: "routing"
description: |-
  This represents the IGMP Querier Settings for Multicast Routing and especially the PIM dynamic routing feature.
---

# smc_igmp_querier_settings (Resource)

This represents the IGMP Querier Settings for Multicast Routing and especially the PIM dynamic routing feature.

## Examples

- [IGMP Querier Settings Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/engines/dynamic_routing/PIM/igmp_querier_settings/main.tf)

Defines IGMP (Internet Group Management Protocol) querier settings for use in dynamic multicast routing environments.

```hcl
resource "smc_igmp_querier_settings" "ipgm_profile" {
  igmp_version   = "igmpv3"
  name           = "tf_igmp_querier_settings"
  query_interval = 180
  comment        = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `igmp_version` (String) The IGMP Version.
- `name` (String) Name of the object.
- `query_interval` (Number) The Query Interval (QDI) in seconds.
- `robustness` (Number) The Robustness value.


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
