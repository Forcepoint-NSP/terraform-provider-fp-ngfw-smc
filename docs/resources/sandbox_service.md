---
page_title: "smc_sandbox_service"
subcategory: "services"
description: |-
  This represents the Sandbox service element (referenced by sandbox settings). It contains information about the portal account, data center, license key and token.
---

# smc_sandbox_service (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the Sandbox service element (referenced by sandbox settings). It contains information about the portal account, data center, license key and token.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `portal_username` (String) The username of the portal account associated with the Sandbox service.
- `sandbox_data_center` (String) This represents a Sandbox Data Center element, which is referenced by the sandbox service. It contains information about the sandbox's hostname, server URL, portal URL, API URL, API key, type, and TLS profile.
- `sandbox_license_key` (String) The license key for the Sandbox service.
- `sandbox_license_token` (String) The license token for the Sandbox service.


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
