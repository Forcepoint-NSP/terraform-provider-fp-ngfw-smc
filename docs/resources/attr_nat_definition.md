---
page_title: "nat_definition"
subcategory: ""
description: |-
  This represents the definition of NAT (Network Address Translation) for Virtual NGFW Engines, allowing you to define how the Firewall Clusters translate network addresses.
---

# nat_definition (Nested-Attribute)

This represents the definition of NAT (Network Address Translation) for Virtual NGFW Engines, allowing you to define how the Firewall Clusters translate network addresses.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `nat_type` (String) The type of NAT translation, which can be 'static' or 'dynamic'. 'Static' means a fixed mapping between private and public IP addresses, while 'dynamic' allows for multiple mappings based on available public IPs.
- `private_ne_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `public_ip` (String) The Public IP Address to be used for NAT, which can be either IPv4 or IPv6.
- `public_ne_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `service_ref` (List of String) URI of the Service Entry.

## Nested Attributes
- `enabled_interface` (List of Blocks, see [here](attr_enabled_interface_entry.md)) The interfaces that are enabled for NAT, which can be used to specify the public interfaces for the NAT definition.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
