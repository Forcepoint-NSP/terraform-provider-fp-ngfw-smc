---
page_title: "smc_delete_old_unused_smc_web_access_data_task"
subcategory: "admin"
description: |-
  This represents a Webswing User File Cleaner Task, which is used to clean up old and unused SMC Web Access user files. It is a type of system task that can be scheduled and executed to ensure that user files are maintained properly.
---

# smc_delete_old_unused_smc_web_access_data_task (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Webswing User File Cleaner Task, which is used to clean up old and unused SMC Web Access user files. It is a type of system task that can be scheduled and executed to ensure that user files are maintained properly.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `resources` (List of String) URI of the resource.


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
