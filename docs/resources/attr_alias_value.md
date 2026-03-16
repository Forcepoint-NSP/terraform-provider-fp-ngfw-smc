---
page_title: "alias_value"
subcategory: ""
description: |-
  This represents the mapping between an Alias and its translated elements, allowing the use of Aliases in policies.
---

# alias_value (Nested-Attribute)

This represents the mapping between an Alias and its translated elements, allowing the use of Aliases in policies.




## Simple Attributes
- `alias_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `cluster_ref` (String) This represents a group of devices, or nodes, that share a given work load. You can cluster Firewalls, IPS engines, and Layer 2 Firewalls to share the load and provide redundancy, allowing, for example, scheduled maintenance that takes one node out of service without interrupting services to the users.
- `resolved_value` (List of String) The resolved values for the Alias, which are the actual values used in the policies.
- `translated_element` (List of String) URI of the Network Element used as translated element.


