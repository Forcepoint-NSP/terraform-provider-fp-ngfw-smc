---
page_title: "smc_sginfo_task"
subcategory: "admin"
description: |-
  This represents an SGInfo Task, which is used to collect system information from the engine. It is a type of task that can be scheduled and executed to gather system diagnostics and configuration details.
---

# smc_sginfo_task (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an SGInfo Task, which is used to collect system information from the engine. It is a type of task that can be scheduled and executed to gather system diagnostics and configuration details.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `destination_path` (String) The destination path where the sginfo result will be stored.
- `include_core_files` (Boolean) Flag to know if the sginfo should include core files.
- `include_slapcat_output` (Boolean) Flag to know if the sginfo should include slapcat output.
- `name` (String) Name of the object.
- `resources` (List of String) URI of the resource.
- `stored_on_server` (Boolean) Flag to know if the sginfo should be stored on the management server.


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
