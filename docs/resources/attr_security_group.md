---
page_title: "security_group"
subcategory: ""
description: |-
  This represents a Security Group, which is a group of VMs that share the same security policy. Security Groups are dynamically changing sets, and their IP addresses are not known in advance.
---

# security_group (Nested-Attribute)

This represents a Security Group, which is a group of VMs that share the same security policy. Security Groups are dynamically changing sets, and their IP addresses are not known in advance.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `isc_id` (String) ISC VSS ID, which is required for the Security Group.
- `isc_name` (String) ISC name, which is required for the Security Group.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `obsolete` (Boolean) Flag indicating if the object is obsolete. This is not required and defaults to false.


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
