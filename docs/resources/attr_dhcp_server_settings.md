---
page_title: "dhcp_server_on_interface"
subcategory: ""
description: |-
  This represents the definition of DHCP Server on an interface, which includes settings for IP ranges, DNS servers, WINS servers, gateways, and lease times.
---

# dhcp_server_on_interface (Nested-Attribute)

This represents the definition of DHCP Server on an interface, which includes settings for IP ranges, DNS servers, WINS servers, gateways, and lease times.




## Simple Attributes
- `default_gateway` (String) The IP address of the Default Gateway through which traffic from clients is routed.
- `default_lease_time` (Number) The Default Lease Time in seconds after which IP addresses assigned to clients must be renewed.
- `dhcp_address_range` (String) The manual DHCP address range or address in the same network space defined for the Physical Interface.
- `dhcp_range_ref` (String) This represents an Address Range, which is a Network Element that defines a range of IP addresses. It includes attributes for the range of IP addresses, which must be valid IPv4 or IPv6 addresses.
- `domain_name_search_list` (String) The 'search list' of Domain Names to be used by the client to locate not-fully-qualified domain names.
- `primary_dns_server` (String) The Primary DNS Server IP address that clients use to resolve domain names.
- `primary_wins_server` (String) The Primary WINS Server IP address that clients use to resolve NetBIOS computer names.
- `secondary_dns_server` (String) The Secondary DNS Server IP address that clients use to resolve domain names.
- `secondary_wins_server` (String) The Secondary WINS Server IP address that clients use to resolve NetBIOS computer names.

## Nested Attributes
- `dhcp_range_per_node` (List of Blocks, see [here](attr_dhcp_range_node.md)) The DHCP Range Nodes which contain a DHCP Address Range for each node in a cluster.

