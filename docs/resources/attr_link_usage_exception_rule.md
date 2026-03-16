---
page_title: "link_usage_exception_rule"
subcategory: "policy"
description: |-
  This represents a rule that defines which ISP link handles a particular outbound network traffic, including source and destination matching criteria, service matching, and ISP link reference.
---

# link_usage_exception_rule (Nested-Attribute)

This represents a rule that defines which ISP link handles a particular outbound network traffic, including source and destination matching criteria, service matching, and ISP link reference.




## Simple Attributes
- `comment` (String) An optional comment for the link usage exception rule, providing additional context or information.
- `isp_link_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.

## Nested Attributes
- `destinations` (Single Block, see [here](attr_destination_matching_part.md)) 
- `services` (Single Block, see [here](attr_service_match_part.md)) 
- `sources` (Single Block, see [here](attr_source_match_part.md)) 

