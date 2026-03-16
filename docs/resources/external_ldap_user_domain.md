---
page_title: "smc_external_ldap_user_domain"
subcategory: ""
description: |-
  This represents the Authentication Domain. Each LDAP server has its own authentication domain in the SMC. One authentication domain can be selected as the default authentication domain, so that users can leave out this information when they authenticate (users can type username instead of username@domain). Users that are stored under non-default authentication domains must always include the domain in the username. If administrative Domains have been configured, you can create separate Authentication Domains for each administrative Domain and select one Authentication Domain in each administrative Domain as the Default. Alternatively, you can select one of the Authentication Domains in the Shared Domain as the Default Authentication Domain across all the administrative Domains.
---

# smc_external_ldap_user_domain (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the Authentication Domain. Each LDAP server has its own authentication domain in the SMC. One authentication domain can be selected as the default authentication domain, so that users can leave out this information when they authenticate (users can type username instead of username@domain). Users that are stored under non-default authentication domains must always include the domain in the username. If administrative Domains have been configured, you can create separate Authentication Domains for each administrative Domain and select one Authentication Domain in each administrative Domain as the Default. Alternatively, you can select one of the Authentication Domains in the Shared Domain as the Default Authentication Domain across all the administrative Domains.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `auth_method` (String) This represents an external authentication method, which can be used for user authentication in the system. It supports various types of authentication methods such as User Password, IAS, IPSec, Pre-Shared Key, RADIUS, TACACS+, and LDAP.
- `browse_ldap_automatically` (Boolean) Indicates whether SMC can connect to the directory server for browsing users and groups.
- `comment` (String) An optional comment for the element. This field is not required.
- `engine_can_connect` (Boolean) Indicates whether the engine can connect to the directory server to resolve users and groups.
- `isdefault` (Boolean) Indicates whether this User Domain is the default one for authentication. Only one User Domain can be set as default.
- `ldap_server` (List of String) URI of the Authentication Server (Active Directory Server or LDAP Server).
- `name` (String) Name of the object.

## Nested Attributes
- `additional_username_suffix` (List of Blocks, see [here](attr_additional_username_suffix.md)) The list of Additional Username Suffixes. You can specify additional username suffixes that users can use when they log in. For example, if your domain is example.com, you can specify @example.com as an additional username suffix. Users can then log in using either their username or their

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
