---
page_title: "ipv6_prefix_list_entry"
subcategory: ""
description: |-
  This represents an entry in an IPv6 Prefix List, which is used to filter BGP routes based on IPv6 prefixes.
---

# ipv6_prefix_list_entry (Nested-Attribute)

This represents an entry in an IPv6 Prefix List, which is used to filter BGP routes based on IPv6 prefixes.




## Simple Attributes
- `action` (String) The action to be taken for this access list entry.
- `comment` (String) A comment for the access or prefix list entry.
- `max_prefix_length` (Number) The maximum prefix length for the IPv6 prefix list entry, used to filter routes.
- `min_prefix_length` (Number) The minimum prefix length for the IPv6 prefix list entry, used to filter routes.
- `name` (String) Name of the object.
- `ne_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `subnet` (String) The IP network or subnet for the access or prefix list entry.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
