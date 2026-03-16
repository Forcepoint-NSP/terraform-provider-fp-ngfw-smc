---
page_title: "used_on"
subcategory: ""
description: |-
  This represents the Used On NAT rule matching criteria, which can be used to determine if a NAT rule is applied to any or none of the firewalls, or a specific firewall element.
---

# used_on (Nested-Attribute)

This represents the Used On NAT rule matching criteria, which can be used to determine if a NAT rule is applied to any or none of the firewalls, or a specific firewall element.




## Simple Attributes
- `any` (Boolean) Indicates whether any firewall matches the criteria. If true, no specific element is required.
- `firewall_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `none` (Boolean) Indicates whether no firewall matches the criteria. If true, no specific element is required.


