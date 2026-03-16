---
page_title: "static_nat"
subcategory: ""
description: |-
  This represents the Static NAT (Network Address Translation) entry, a sub part of NAT Rule. Source addresses in matching connections are translated using the same number of IP addresses as there are possible original source addresses. Each translated IP address corresponds to one original IP address.
---

# static_nat (Nested-Attribute)

This represents the Static NAT (Network Address Translation) entry, a sub part of NAT Rule. Source addresses in matching connections are translated using the same number of IP addresses as there are possible original source addresses. Each translated IP address corresponds to one original IP address.




## Simple Attributes
- `automatic_proxy` (Boolean) Indicates whether Automatic Proxy ARP is enabled. This allows the engine to answer address queries regarding the translated address(es).

## Nested Attributes
- `original_value` (Single Block, see [here](attr_translation_value.md)) 
- `translated_value` (Single Block, see [here](attr_translation_value.md)) 

