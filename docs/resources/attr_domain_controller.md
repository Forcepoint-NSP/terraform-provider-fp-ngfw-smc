---
page_title: "domain_controller"
subcategory: ""
description: |-
  This represents the Active Directory Domain Controller settings, including server type, IP address, user credentials, and data expiration time.
---

# domain_controller (Nested-Attribute)

This represents the Active Directory Domain Controller settings, including server type, IP address, user credentials, and data expiration time.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `expiration_time` (Number) The time in seconds to keep user ID data, ranging from 60 seconds (1 minute) to 604800 seconds (7 days). By default, it is set to 28800 seconds (8 hours).
- `ipaddress` (String) The IP address of the domain controller, supporting both IPv4 and IPv6 formats.
- `name` (String) Name of the object.
- `password` (String) The password for the user account with Domain Admin credentials, which is stored in a ciphered format.
- `server_type` (String) The type of server to query for user ID data, such as 'ActiveDirectory'. 'dc' for Domain Controller and 'exchange' for Exchange Server.
- `user` (String) The user name of a user in the domain with permission to execute WMI queries from a remote computer.


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
