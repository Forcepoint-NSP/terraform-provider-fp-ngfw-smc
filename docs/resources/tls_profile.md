---
page_title: "smc_tls_profile"
subcategory: ""
description: |-
  This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.
---

# smc_tls_profile (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `accept_wildcard` (Boolean) Is server identity check accepts wildcards. Default is false, can be overridden by the chosen server identity method.
- `check_revocation` (Boolean) Is certificate revocation checked. Default is true.
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `tls_cryptography_suites` (String) This represents a TLS Cryptography Suite Set Element, which contains a set of cryptographic suites used in SSL VPN configurations.
- `tls_trusted_ca_ref` (List of String) URI of the trusted CA.
- `use_only_subject_alt_name` (Boolean) Use Only Subject Alt Name When the TLS identity is a DNS name, uses only Subject Alternative Name (SAN) certificate matching.


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
