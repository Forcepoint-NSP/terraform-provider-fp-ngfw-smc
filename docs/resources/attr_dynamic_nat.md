---
page_title: "dynamic_nat"
subcategory: ""
description: |-
  This represents the Dynamic NAT (Network Address Translation) entry, a sub part of NAT Rule. Source addresses in matching connections are translated using a smaller pool of IP addresses than there are original source addresses included in the rule. Many hosts can use the same IP address, and the connections are distinguished by allocating a different TCP or UDP port for each connection. Also used for activating an Outbound Multi-Link configuration (IPv4 only). Because ports are needed to keep track of connections, dynamic NAT only works with TCP and UDP protocols. If the protocol used in the communications is not transported on top of TCP or UDP, the communicating applications must encapsulate the packets in TCP or UDP (NAT traversal) to communicate through dynamic NAT.
---

# dynamic_nat (Nested-Attribute)

This represents the Dynamic NAT (Network Address Translation) entry, a sub part of NAT Rule. Source addresses in matching connections are translated using a smaller pool of IP addresses than there are original source addresses included in the rule. Many hosts can use the same IP address, and the connections are distinguished by allocating a different TCP or UDP port for each connection. Also used for activating an Outbound Multi-Link configuration (IPv4 only). Because ports are needed to keep track of connections, dynamic NAT only works with TCP and UDP protocols. If the protocol used in the communications is not transported on top of TCP or UDP, the communicating applications must encapsulate the packets in TCP or UDP (NAT traversal) to communicate through dynamic NAT.




## Simple Attributes
- `automatic_proxy` (Boolean) Indicates whether Automatic Proxy ARP is enabled. This allows the engine to answer address queries regarding the translated address(es).

## Nested Attributes
- `translation_values` (List of Blocks, see [here](attr_translation_value.md)) A list of translation values that define how source addresses are translated in dynamic NAT.

