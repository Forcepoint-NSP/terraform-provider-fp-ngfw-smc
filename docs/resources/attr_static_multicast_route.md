---
page_title: "static_multicast_route"
subcategory: ""
description: |-
  This represents a Firewall multicast routing entry for Static/IGMP Proxy multicast routing modes, which defines the source and destination IP addresses and interfaces for multicast traffic routing.
---

# static_multicast_route (Nested-Attribute)

This represents a Firewall multicast routing entry for Static/IGMP Proxy multicast routing modes, which defines the source and destination IP addresses and interfaces for multicast traffic routing.




## Simple Attributes
- `dest_interface` (List of String) The destination interfaces for the multicast route specified by their interface IDs. This is a collection of interface IDs that the multicast traffic will be routed to.
- `dest_ip` (String) The destination IP address for the multicast route.
- `source_interface` (String) The source interface ID for the multicast route.
- `source_ip` (String) The source IP address for the multicast route.


