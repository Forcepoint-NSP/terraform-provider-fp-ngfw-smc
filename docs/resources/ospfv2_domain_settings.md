---
page_title: "smc_ospfv2_domain_settings"
subcategory: "routing"
description: |-
  This represents the OSPFv2 Domain Settings for Dynamic Routing Firewall functionality, including properties such as area ID and authentication settings.
---

# smc_ospfv2_domain_settings (Resource)

This represents the OSPFv2 Domain Settings for Dynamic Routing Firewall functionality, including properties such as area ID and authentication settings.

## Examples

- [OSPFv2 Domain Settings Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/engines/dynamic_routing/OSPFv2/ospfv2_domain_settings/main.tf)

Defines domain-wide settings for OSPFv2 routing.

```hcl
resource "smc_ospfv2_domain_settings" "ospfv2_domain_settings" {
  abr_type                = "cisco"
  auto_cost_bandwidth     = 100
  deprecated_algorithm    = false
  initial_delay           = 200
  initial_hold_time       = 1000
  max_hold_time           = 10000
  name                    = "tf_ospfv2_domain_settings"
  shutdown_max_metric_lsa = 0
  startup_max_metric_lsa  = 0
  comment                 = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `abr_type` (String) The Area Border Router type, which can be either 'cisco' or 'standard' or 'shortcut'.
- `auto_cost_bandwidth` (Number) The auto-cost reference bandwidth in MBit/s.
- `comment` (String) An optional comment for the element. This field is not required.
- `deprecated_algorithm` (Boolean) Indicates whether the deprecated algorithm (RFC 1583) is used.
- `initial_delay` (Number) The initial delay in milliseconds before OSPF starts processing.
- `initial_hold_time` (Number) The initial hold time in milliseconds for OSPF.
- `max_hold_time` (Number) The maximum hold time in milliseconds for OSPF.
- `name` (String) Name of the object.
- `shutdown_max_metric_lsa` (Number) The shutdown max metric link-state advertisement in seconds.
- `startup_max_metric_lsa` (Number) The startup max metric link-state advertisement in seconds.


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
