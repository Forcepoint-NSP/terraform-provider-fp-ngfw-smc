---
page_title: "smc_sandbox_data_center"
subcategory: ""
description: |-
  This represents a Sandbox Data Center element, which is referenced by the sandbox service. It contains information about the sandbox's hostname, server URL, portal URL, API URL, API key, type, and TLS profile.
---

# smc_sandbox_data_center (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Sandbox Data Center element, which is referenced by the sandbox service. It contains information about the sandbox's hostname, server URL, portal URL, API URL, API key, type, and TLS profile.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `api_key` (String) The API key for the Sandbox Data Center. This is required for authentication when accessing the sandbox service.
- `api_url` (String) The API URL of the Sandbox Data Center. This is used for API interactions with the sandbox service.
- `comment` (String) An optional comment for the element. This field is not required.
- `hostname` (String) The hostname of the Sandbox Data Center. This is required for the sandbox service to connect.
- `name` (String) Name of the object.
- `portal_url` (String) The portal URL of the Sandbox Data Center. This is used for accessing the sandbox portal.
- `sandbox_type` (String) The type of the Sandbox Data Center. Valid values are 'forcepoint_sandbox', 'local_shm_sandbox', 'cloud_sandbox', and 'local_sandbox'.
- `server_url` (String) The server URL of the Sandbox Data Center. This is required for the sandbox service to connect.


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
