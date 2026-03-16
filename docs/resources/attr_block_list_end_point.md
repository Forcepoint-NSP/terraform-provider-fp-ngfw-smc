---
page_title: "block_list_endpoint"
subcategory: ""
description: |-
  This represents a block list Endpoint, a sub part of Rule Action Options block list Scope to configure the block list settings of the Rule.
---

# block_list_endpoint (Nested-Attribute)

This represents a block list Endpoint, a sub part of Rule Action Options block list Scope to configure the block list settings of the Rule.




## Simple Attributes
- `address_mode` (String) The type of address mode for the block list endpoint. 'any': Matches any IP address. 'attacker': Matches the IP address identified as the originator of an attack by the Situation element that is triggered. 'victim': Matches the IP address identified as the target of an attack by the Situation element that is triggered. 'ip_source': Matches the IP address that is the source of the packet(s) that trigger the detected situation. 'ip_destination': Matches the IP address that is the destination of the packet(s) that trigger the detected situation. 'tcp_client': Matches the IP address that is the source of the TCP connection that triggers the detected situation. 'tcp_server': Matches the IP address that is the destination of the TCP connection that triggers the detected situation. 'address': Matches only the fixed IP address.
- `ip_network` (String) The IP network range for the block list endpoint.
- `port1` (Number) The first port value for the block list endpoint.
- `port2` (Number) The second port value for the block list endpoint.
- `port_mode` (String) The type of port mode for the block list endpoint. 'not_used': Matches any IP traffic regardless of the protocol or port. 'from_traffic': Matches the IP protocol and the port number in the traffic that triggered the block list entry. 'predefined_tcp': Matches only TCP traffic through the TCP port or the range of TCP ports. 'predefined_udp': Matches only UDP traffic through the UDP port or the range of UDP ports.


