---
page_title: "smc_api_user"
subcategory: "admin"
description: |-
  This represents a API Client. It is an element that defines the details of a single person that is allowed to log on to Management Server through REST services.
---

# smc_api_user (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a API Client. It is an element that defines the details of a single person that is allowed to log on to Management Server through REST services.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `allowed_to_login_in_shared` (Boolean) Indicates whether the API Client is allowed to log in to the Shared Domain. If true, the API Client can access shared resources and settings.
- `comment` (String) An optional comment for the element. This field is not required.
- `console_superuser` (Boolean) Indicates whether the administrator user is a console superuser. This applies only to an Appliance and grants additional privileges for console access.
- `name` (String) Name of the object.
- `superuser` (Boolean) Indicates whether the administrator user is a superuser. A superuser has full access to all system functions and settings.

## Nested Attributes
- `permissions` (Single Block, see [here](attr_admin_permissions.md)) 

## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `enabled` (Boolean) The status of the administrator user. It indicates whether the user is enabled or disabled.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `is_user_locked` (Boolean) Indicates whether the administrator user is locked. A locked user cannot log in until it is unlocked by an administrator.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
