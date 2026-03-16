---
page_title: "smc_virtual_resource"
subcategory: "engines"
description: |-
  This represents a set of resources on the Master NGFW Engine that are allocated to each Virtual NGFW Engine. It includes properties such as ID, domain allocation, and resource limits.
---

# smc_virtual_resource (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a set of resources on the Master NGFW Engine that are allocated to each Virtual NGFW Engine. It includes properties such as ID, domain allocation, and resource limits.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `allocated_domain_ref` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `comment` (String) An optional comment for the element. This field is not required.
- `connection_limit` (Number) The limit on concurrent connections from a single source and/or destination IP address. When the limit is reached, further connection attempts are blocked until existing connections are closed.
- `name` (String) Name of the object.
- `rate_limit` (Number) The rate limit for the Virtual Resource in kbps. By default (0), it is considered as unlimited.
- `show_master_nic` (Boolean) Indicates whether the Master NGFW Engine's Physical Interface IDs are shown in the Virtual NGFW Engine's Interface properties.
- `throughput_limit` (Number) The throughput limit for the Virtual Resource in kbps. By default (0), it is considered as unlimited.
- `vfw_id` (Number) The unique identifier for the Virtual Resource, starting from 1.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
