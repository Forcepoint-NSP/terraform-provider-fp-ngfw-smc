---
page_title: "smc_trusted_update_certificate_authority"
subcategory: ""
description: |-
  This represents a Trusted Update Certificate Authority, which is used to manage trusted update certificate authorities in the system.
---

# smc_trusted_update_certificate_authority (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Trusted Update Certificate Authority, which is used to manage trusted update certificate authorities in the system.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `certificate` (String) The certificate of the Certificate Authority in string format.
- `comment` (String) An optional comment for the element. This field is not required.
- `crl_checking_enabled` (Boolean) Indicates whether CRL checking is enabled for this Certificate Authority.
- `internal_ca` (Boolean) Indicates whether this Certificate Authority is an internal CA (used for IPSec).
- `name` (String) Name of the object.
- `ocsp_checking_enabled` (Boolean) Indicates whether OCSP checking is enabled for this Certificate Authority.


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
