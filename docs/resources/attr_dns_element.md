---
page_title: "domain_server_address"
subcategory: ""
description: |-
  This represents an ordered element in the list of DNS for an engine, including the IP address and rank for sorting.
---

# domain_server_address (Nested-Attribute)

This represents an ordered element in the list of DNS for an engine, including the IP address and rank for sorting.




## Simple Attributes
- `ne_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `rank` (Number) The rank value to sort DNS entries, which must be a positive float. Only distinct values are accepted.
- `value` (String) The IP address of the DNS server that the Firewall Cluster will use to resolve signature mirrors, domain names, and URL filtering categorization services (which are defined as URLs).


