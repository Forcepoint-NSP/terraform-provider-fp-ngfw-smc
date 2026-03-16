---
page_title: "smc_bgp_profile"
subcategory: "routing"
description: |-
  This represents the BGP Profile for Dynamic Routing Firewall functionality, including port settings, distances, and BGP entries.
---

# smc_bgp_profile (Resource)

This represents the BGP Profile for Dynamic Routing Firewall functionality, including port settings, distances, and BGP entries.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/BGP/bgp_profile) for an example

This example creates a BGP profile with aggregation, redistribution, and monitoring.

```hcl
resource "smc_bgp_profile" "bgp_profile" {
  comment = var.resource_comment
  aggregation_entry {
    mode   = "as_set_and_summary"
    subnet = smc_network.network1.id
  }
  bmp_entry {
    bmp_address                = "192.168.10.1"
    bmp_connect_through_master = true
    bmp_port                   = 179
  }
  bmp_entry {
    bmp_address                = "10.100.100.14"
    bmp_connect_through_master = false
    bmp_port                   = 1179
  }
  distance_entry {
    distance = 220
    subnet   = smc_network.network1.id
  }
  distance_entry {
    distance = 255
    subnet   = smc_network.network2.id
  }
  external = 40
  internal = 220
  local    = 250
  name     = "tf_bgp_profile"
  port     = 179
  redistribution_entry {
    enabled     = false
    filter_type = "none"
    type        = "kernel"
  }
  redistribution_entry {
    enabled     = true
    filter_type = "none"
    metric      = 120
    type        = "static"
  }
  redistribution_entry {
    enabled     = false
    filter_type = "none"
    type        = "connected"
  }
  redistribution_entry {
    enabled     = false
    filter_type = "none"
    type        = "ospfv2"
  }
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `external` (Number) The External distance value, which should be between 1 and 255.
- `internal` (Number) The Internal distance value, which should be between 1 and 255.
- `local` (Number) The Local distance value, which should be between 1 and 255.
- `name` (String) Name of the object.
- `port` (Number) The BGP port value.

## Nested Attributes
- `aggregation_entry` (List of Blocks, see [here](attr_bgp_aggregation_entry.md)) A list of BGP aggregation entries, which define the aggregation of BGP routes.
- `bmp_entry` (List of Blocks, see [here](attr_bgp_bmp_entry.md)) A list of BGP BMP entries, which define the BGP Monitoring Protocol entries.
- `distance_entry` (List of Blocks, see [here](attr_bgp_distance_entry.md)) A list of BGP distance entries, which define the distance for specific BGP routes.
- `redistribution_entry` (List of Blocks, see [here](attr_dynamic_routing_redistribution_entry.md)) A list of BGP Redistribution entries, which define how routes are redistributed in BGP.

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
