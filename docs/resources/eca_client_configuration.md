---
page_title: "smc_eca_client_config"
subcategory: ""
description: |-
  This represents a ECA Client Configuration, which is used to manage the client-side configuration for ECA (Endpoint Compliance Agent). It includes settings such as trusted certificate authorities and auto-discovery options.
---

# smc_eca_client_config (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a ECA Client Configuration, which is used to manage the client-side configuration for ECA (Endpoint Compliance Agent). It includes settings such as trusted certificate authorities and auto-discovery options.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `auto_discovery` (Boolean) Indicates whether auto-discovery is enabled for the ECA client configuration.
- `comment` (String) An optional comment for the element. This field is not required.
- `eca_ca_ref` (List of String) URI of the trusted CA.
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
