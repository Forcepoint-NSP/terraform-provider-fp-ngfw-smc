---
page_title: "ip_netlink_weight"
subcategory: ""
description: |-
  This represents a network element that defines the external NATed destination IP Address for the Server Pool, the NetLink to use, the Network element used for the Server Pool's external NATed address, and the weight of this Net Link in the pool.
---

# ip_netlink_weight (Nested-Attribute)

This represents a network element that defines the external NATed destination IP Address for the Server Pool, the NetLink to use, the Network element used for the Server Pool's external NATed address, and the weight of this Net Link in the pool.




## Simple Attributes
- `arp_generate` (Boolean) Automatically generate a proxy ARP for the NATed address in the selected Network.
- `netlink_ref` (String) This represents a Static NetLink, which is a type of NetLink used for routing in Multi-Link features. It includes attributes for gateway, networks, DNS elements, and outbound IP addresses.
- `network_ref` (String) This represents a Network, which is a Network Element that represents a (sub)network of computers. It includes attributes for IPv4 and IPv6 networks, broadcast address, and validation patterns.
- `weight` (Number) The weight of this Net Link in the pool.


