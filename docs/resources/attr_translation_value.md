---
page_title: "translation_value"
subcategory: ""
description: |-
  This represents the Firewall NAT Translation value, which includes the translated IP address or network, the translated element, and the port range for source IP address translation.
---

# translation_value (Nested-Attribute)

This represents the Firewall NAT Translation value, which includes the translated IP address or network, the translated element, and the port range for source IP address translation.




## Simple Attributes
- `element` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `ip_descriptor` (String) The IP address pool of IP address(es) used for translation. The minimum size is one IP address. The number of IP addresses required depends on the allowed port range and peak concurrent connections.
- `max_port` (Number) The end of the port range for source IP address translation. The default is the highest possible port, 65535.
- `min_port` (Number) The start of the port range for source IP address translation. The default is the beginning of the 'free' high port range, 1024.


