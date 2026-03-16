---
page_title: "smc_integrated_uis"
subcategory: "services"
description: |-
  This represents an Integrated User Identification Service, which is used to identify users based on their network activity. It allows for integration with authentication domains and provides configuration options for initial query time, polling interval, and ignore values.
---

# smc_integrated_uis (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Integrated User Identification Service, which is used to identify users based on their network activity. It allows for integration with authentication domains and provides configuration options for initial query time, polling interval, and ignore values.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `iuis_domain` (String) This represents the Authentication Domain. Each LDAP server has its own authentication domain in the SMC. One authentication domain can be selected as the default authentication domain, so that users can leave out this information when they authenticate (users can type username instead of username@domain). Users that are stored under non-default authentication domains must always include the domain in the username. If administrative Domains have been configured, you can create separate Authentication Domains for each administrative Domain and select one Authentication Domain in each administrative Domain as the Default. Alternatively, you can select one of the Authentication Domains in the Shared Domain as the Default Authentication Domain across all the administrative Domains.
- `iuis_initial_query_time` (Number) The initial query time in seconds, which defines how long the system waits before performing the first user identification query.
- `iuis_polling_interval` (Number) The polling interval in seconds, which defines how often the system checks for user identification updates.
- `name` (String) Name of the object.

## Nested Attributes
- `iuis_ignore` (List of Blocks, see [here](attr_integrated_uis_ignore_value.md)) The ignore values, which define specific IP addresses or networks that should be ignored by the user identification service.

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
