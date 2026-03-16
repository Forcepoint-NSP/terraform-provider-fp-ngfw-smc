---
page_title: "link_usage_preference"
subcategory: ""
description: |-
  This represents a Link Usage Preference, which includes a QoS Class and a preference type (prefer, avoid, do not use).
---

# link_usage_preference (Nested-Attribute)

This represents a Link Usage Preference, which includes a QoS Class and a preference type (prefer, avoid, do not use).




## Simple Attributes
- `preference` (String) The type of preference for the QoS Class, which can be 'prefer', 'avoid', or 'do not use'.
- `qos_class` (String) This represents a QoS Class, which is an element that links a rule in a QoS Policy to one or more Firewall Actions. The traffic allowed in the access rule is assigned the QoS Class defined for the rule, and the QoS class is used as the matching criteria for applying QoS Policy rules.


