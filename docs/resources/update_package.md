---
page_title: "smc_update_package"
subcategory: "admin"
description: |-
  This represents an Update Package. It contains information about the package ID, release date, state, activation date, and the SMC version where the activation has been done.
---

# smc_update_package (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Update Package. It contains information about the package ID, release date, state, activation date, and the SMC version where the activation has been done.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.


## Readonly Attributes
- `activation_date` (Number) Timestamp when the update package was activated.
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `package_id` (Number) Unique identifier for the update package.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `release_date` (String) Release date of the update package.
- `smc_version` (String) Version of the SMC where the update package was activated.
- `state` (String) State of the update package.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
