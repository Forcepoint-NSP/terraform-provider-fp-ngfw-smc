---
page_title: "trap_destination"
subcategory: ""
description: |-
  This represents a destination for SNMP traps, including the IP address and port.
---

# trap_destination (Nested-Attribute)

This represents a destination for SNMP traps, including the IP address and port.




## Simple Attributes
- `address` (String) The IP address of the SNMP trap destination.
- `comment` (String) An optional comment for the element. This field is not required.
- `destination_port` (Number) The port number of the SNMP trap destination.
- `name` (String) Name of the object.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
