---
page_title: "rbvpn_geneve_settings"
subcategory: "vpn"
description: |-
  This defines the Route-Based Tunnel settings for GENEVE type, including Virtual Network Identifier (VNI) and Destination Port.
---

# rbvpn_geneve_settings (Nested-Attribute)

This defines the Route-Based Tunnel settings for GENEVE type, including Virtual Network Identifier (VNI) and Destination Port.




## Simple Attributes
- `geneve_destination_port` (Number) Destination Port value between 1 and 65 535. When left empty, the default value (6081) is used.
- `geneve_vni` (Number) Virtual Network Identifier (VNI) value between 0 and 16 777 215. When left empty, the default value (0) is used.


