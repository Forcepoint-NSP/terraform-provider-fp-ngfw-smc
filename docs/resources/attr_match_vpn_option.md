---
page_title: "match_vpn_options"
subcategory: "vpn"
description: |-
  This represents the Match VPN Options, which is a sub part of Firewall Access Rule used for matching the source VPN.
---

# match_vpn_options (Nested-Attribute)

This represents the Match VPN Options, which is a sub part of Firewall Access Rule used for matching the source VPN.




## Simple Attributes
- `match_type` (String) The type of match for the VPN. 'normal': the rule only matches traffic from the specified VPNs. 'no vpn': the rule only matches non-VPN traffic. 'mobile vpn': the rule only matches traffic from IPsec VPN client.
- `match_vpns` (List of String) URI of the Policy Based VPN.


