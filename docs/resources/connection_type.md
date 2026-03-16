---
page_title: "smc_connection_type"
subcategory: ""
description: |-
  This represents a connection type used in endpoints, which defines the connectivity group, mode, and link type for VPN connections.
---

# smc_connection_type (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a connection type used in endpoints, which defines the connectivity group, mode, and link type for VPN connections.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `connectivity_group` (Number) The connectivity group ID for the VPN connections, starting from 1.
- `link_type_ref` (String) This represents the Link Type, which defines the characteristics and behavior of a link in the VPN configuration.
- `mode` (String) The mode of the connection type, indicating how the connection is balanced.
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
