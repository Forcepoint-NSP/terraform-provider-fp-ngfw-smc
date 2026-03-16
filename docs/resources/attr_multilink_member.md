---
page_title: "multilink_member"
subcategory: ""
description: |-
  This represents a member of an Outbound Multi-Link element, which defines the NetLink used as a member, the network it connects to, and its role in the Multi-Link configuration.
---

# multilink_member (Nested-Attribute)

This represents a member of an Outbound Multi-Link element, which defines the NetLink used as a member, the network it connects to, and its role in the Multi-Link configuration.




## Simple Attributes
- `netlink_ref` (String) This represents a NetLink, which is used for implementing routing of Multi-Link features. NetLinks can represent any IP-based network links (such as ISP routers, xDSL, leased lines, dial-up modems). NetLinks are combined together into an Outbound Multi-link.
- `netlink_role` (String) The role of the NetLink in the Outbound Multi-Link configuration, either 'active' or 'standby'.
- `network_ref` (String) This represents a Network, which is a Network Element that represents a (sub)network of computers. It includes attributes for IPv4 and IPv6 networks, broadcast address, and validation patterns.

## Nested Attributes
- `qos_role` (List of Blocks, see [here](attr_qos_class_with_role.md)) A list of QoS classes with roles associated with this NetLink, defining how traffic is prioritized and managed.

