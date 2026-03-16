---
page_title: "v3user"
subcategory: ""
description: |-
  This represents an SNMP User for SNMP version 3. It includes authentication and privacy protocols and passwords.
---

# v3user (Nested-Attribute)

This represents an SNMP User for SNMP version 3. It includes authentication and privacy protocols and passwords.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `snmp_authentication_password` (String) The password used for authentication by the SNMP user.
- `snmp_authentication_protocol` (String) The authentication protocol used by the SNMP user. It can be one of the following: MD5, SHA, SHA-224, SHA-256, SHA-384, or SHA-512.
- `snmp_privacy_protocol` (String) The privacy protocol used by the SNMP user. It can be one of the following: DES, AES-128, AES-192, or AES-256.
- `snmp_private_password` (String) The password used for privacy by the SNMP user.


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
