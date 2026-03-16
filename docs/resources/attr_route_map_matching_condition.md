---
page_title: "match_condition"
subcategory: ""
description: |-
  This represents a Route Map Matching Condition, which defines the criteria for matching routes in a Route Map Rule.
---

# match_condition (Nested-Attribute)

This represents a Route Map Matching Condition, which defines the criteria for matching routes in a Route Map Rule.




## Simple Attributes
- `access_list_ref` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `external_bgp_peer_address_ref` (String) This represents the External BGP Peer for Dynamic Routing Firewall functionality. It is used to configure external BGP peers in the firewall's dynamic routing capabilities.
- `fwcluster_peer_address_ref` (String) This represents a group of devices, or nodes, that share a given work load. You can cluster Firewalls to share the load and provide redundancy, allowing, for example, scheduled maintenance that takes one node out of service without interrupting services to the users.
- `metric` (Number) The Metric value used for matching routes in the Route Map Rule.
- `next_hop_ref` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `rank` (Number) The rank of this matching condition, which determines its priority in the Route Map Rule.
- `type` (String) The type of matching condition, which can be 'element', 'next_hop', 'metric', or 'peer_address'.


