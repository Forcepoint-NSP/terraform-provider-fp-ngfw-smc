---
page_title: "next_hop"
subcategory: ""
description: |-
  This represents the Next Hop element for Route Map Matching Condition rule, which can specify an IP address, peer address, or a linked network element.
---

# next_hop (Nested-Attribute)

This represents the Next Hop element for Route Map Matching Condition rule, which can specify an IP address, peer address, or a linked network element.




## Simple Attributes
- `next_hop_ip` (String) The IP address (IPv4 or IPv6) for the next hop. Required if no next hop reference or peer address is specified.
- `next_hop_peer_address` (Boolean) Indicates if the next hop is a peer address. Required if no IP address or next hop reference is specified.
- `next_hop_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.


