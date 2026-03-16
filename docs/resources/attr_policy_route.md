---
page_title: "policy_route"
subcategory: "policy"
description: |-
  This defines a Policy Route, which is used to specify routing rules for network traffic based on source and destination IP networks and a gateway IP address.
---

# policy_route (Nested-Attribute)

This defines a Policy Route, which is used to specify routing rules for network traffic based on source and destination IP networks and a gateway IP address.




## Simple Attributes
- `comment` (String) An optional comment for the policy route, providing additional context or information.
- `destination` (String) The destination IP network for the policy route, specified in CIDR notation.
- `gateway_ip` (String) The gateway IP address for the policy route, which is used to route traffic to the specified destination network.
- `source` (String) The source IP network for the policy route, specified in CIDR notation.


