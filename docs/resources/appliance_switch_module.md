---
page_title: "smc_appliance_switch_module"
subcategory: ""
description: |-
  This represents an Appliance Switch Module, which is used to manage the switch modules in the appliance. It includes information such as the number of ports and configuration value.
---

# smc_appliance_switch_module (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Appliance Switch Module, which is used to manage the switch modules in the appliance. It includes information such as the number of ports and configuration value.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `conf_value` (String) The Configuration Value of this Appliance Switch Module. This is the value used in the engine configuration files.
- `name` (String) Name of the object.
- `port_count` (Number) The number of ports on this Appliance Switch Module. Should be greater than zero.


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
