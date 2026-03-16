---
page_title: "ospfv2"
subcategory: "routing"
description: |-
  This represents the OSPFv2 Dynamic Routing settings for the firewall, including whether it is enabled, the router ID, and the associated OSPFv2 profile.
---

# ospfv2 (Nested-Attribute)

This represents the OSPFv2 Dynamic Routing settings for the firewall, including whether it is enabled, the router ID, and the associated OSPFv2 profile.




## Simple Attributes
- `enabled` (Boolean) Indicates whether OSPFv2 Dynamic Routing is enabled.
- `ospfv2_profile_ref` (String) This represents an abstract OSPF Profile used as Dynamic Routing element. It contains settings for OSPF routing, including distances and redistribution entries.
- `router_id` (String) The OSPFv2 Router ID, which can be an IPv4 address or null for automatic mode.


