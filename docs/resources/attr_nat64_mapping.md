---
page_title: "nat64_mapping"
subcategory: ""
description: |-
  This represents a static mapping in IPv6 Transition Mechanism settings, including IPv6 and IPv4 addresses, ports, protocol, and rank.
---

# nat64_mapping (Nested-Attribute)

This represents a static mapping in IPv6 Transition Mechanism settings, including IPv6 and IPv4 addresses, ports, protocol, and rank.




## Simple Attributes
- `comment` (String) An optional comment for the NAT64 mapping.
- `nat64_mapping_protocol_ref` (String) This represents an IP-proto service, which is used to define a service based on the IP protocol number. It includes a protocol number that specifies the protocol used by the traffic.
- `rank` (Number) The rank for ordering NAT64 mappings, where lower values indicate higher priority.


