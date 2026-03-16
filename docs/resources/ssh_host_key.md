---
page_title: "smc_ssh_host_key"
subcategory: "engines"
description: |-
  This represents an SSH Host Key for the Sidewinder SSH Proxy. It stores a key pair for use in the SSH Proxy. The key data is not directly accessible through this interface. During creation of a new key, the type and length fields are required.
---

# smc_ssh_host_key (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an SSH Host Key for the Sidewinder SSH Proxy. It stores a key pair for use in the SSH Proxy. The key data is not directly accessible through this interface. During creation of a new key, the type and length fields are required.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `comment` (String) An optional comment for the element. This field is not required.
- `key_fingerprint` (String) The SHA256 fingerprint of the SSH host key. This is read-only and used for identification purposes.
- `key_length` (Number) The length of the SSH host key. Valid values for RSA and DSA are 1024 and 2048. Valid values for ECDSA are 256, 384, and 521. Required when creating a new key, read-only otherwise.
- `key_service_ref` (List of String) URI of the TCP service.
- `key_type` (String) The type of SSH host key. Valid values are RSA, DSA, and ECDSA. Required when creating a new key, read-only otherwise.
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
