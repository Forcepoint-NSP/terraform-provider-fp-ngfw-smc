---
page_title: "dynamic_routing"
subcategory: ""
description: |-
  This represents the Dynamic Routing Settings, which includes BGP and OSPFv2 settings, antispoofing networks, and path count for equal cost multi-path routing.
---

# dynamic_routing (Nested-Attribute)

This represents the Dynamic Routing Settings, which includes BGP and OSPFv2 settings, antispoofing networks, and path count for equal cost multi-path routing.




## Simple Attributes
- `antispoofing_ne_ref` (List of String) URI of the Network Element used for antispoofing.
- `path_count` (Number) The path count for equal cost multi-path routing, which determines the number of paths used for routing decisions.

## Nested Attributes
- `bgp` (Single Block, see [here](attr_bgp_settings.md)) 
- `ospfv2` (Single Block, see [here](attr_ospfv2_settings.md)) 

