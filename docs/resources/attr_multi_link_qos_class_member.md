---
page_title: "multilink_qos_class"
subcategory: ""
description: |-
  This represents a member of a Multi-Link QOS Class, which includes the QOS Class reference and the balance method used for load balancing.
---

# multilink_qos_class (Nested-Attribute)

This represents a member of a Multi-Link QOS Class, which includes the QOS Class reference and the balance method used for load balancing.




## Simple Attributes
- `balance_method` (String) The method used for load balancing in the Multi-Link QOS Class, which can be 'notdefined', 'ratio', or 'rtt'.
- `key` (Number) The unique key for this Multi-Link QOS Class Member, used for updates. For internal use only.
- `qos_class_ref` (String) This represents a QoS Class, which is an element that links a rule in a QoS Policy to one or more Firewall Actions. The traffic allowed in the access rule is assigned the QoS Class defined for the rule, and the QoS class is used as the matching criteria for applying QoS Policy rules.


