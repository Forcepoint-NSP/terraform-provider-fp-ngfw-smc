---
page_title: "bgp"
subcategory: "routing"
description: |-
  This represents the BGP Dynamic Routing Firewall Settings, including BGP Profile, Router ID, and Autonomous System.
---

# bgp (Nested-Attribute)

This represents the BGP Dynamic Routing Firewall Settings, including BGP Profile, Router ID, and Autonomous System.




## Simple Attributes
- `bgp_as_ref` (String) This represents the Autonomous System for Dynamic Routing Firewall functionality. The BGP Version is 4/4+. It includes the AS Number in decimal notation, which is used to identify the autonomous system in BGP routing.
- `bgp_profile_ref` (String) This represents the BGP Profile for Dynamic Routing Firewall functionality, including port settings, distances, and BGP entries.
- `bmp_router_id` (String) The BMP Router ID, which can be in the format [0-255]:[0-65535], V.X.Y.Z:[0-255], or [0-65535]:[0-255].
- `bmp_router_id_type` (String) The BMP Router ID Type, which indicates the format of the BMP Router ID. Allowed values are 0 for [0-255]:[0-65535], 1 for V.X.Y.Z:[0-255], and 2 for [0-65535]:[0-255].
- `enabled` (Boolean) Indicates if BGP settings are enabled. If true, BGP routing is active.
- `router_id` (String) The BGP Router ID, which can be an IPv4 address or null for automatic mode. Required for BGP configuration.

## Nested Attributes
- `announced_ne_setting` (List of Blocks, see [here](attr_bgp_announced_network_association.md)) The BGP Announced Network Element settings, which can include multiple associations.

