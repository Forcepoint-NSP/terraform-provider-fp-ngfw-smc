---
page_title: "smc_external_ldap_user"
subcategory: "admin"
description: |-
  This represents an external LDAP User defined on the external LDAP server. It contains user details such as display name, email, phone number, job title, office location, frame IP, activation date, expiration date, and days left until expiration.
---

# smc_external_ldap_user (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an external LDAP User defined on the external LDAP server. It contains user details such as display name, email, phone number, job title, office location, frame IP, activation date, expiration date, and days left until expiration.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `activation_date` (Number) The activation date of the user in milliseconds since epoch.
- `comment` (String) An optional comment for the element. This field is not required.
- `days_left` (Number) The number of days left until the user expires. -1 means that it is impossible to know, e.g., if the user has no expiration date.
- `display_name` (String) The display name of the user.
- `email` (String) The email address of the user.
- `expiration_date` (Number) The expiration date of the user in milliseconds since epoch. If the user is not set to expire, this value may be null.
- `frame_ip` (String) The frame IP address of the user.
- `job_title` (String) The job title of the user.
- `name` (String) Name of the object.
- `office_location` (String) The office location of the user.
- `phone_number` (String) The phone number of the user.
- `unique_id` (String) The unique id of the User/User Group element.
- `user_domain` (String) This represents a User Domain, which is used to define the authentication domain for users. It can be either an authentication domain or the internal domain.


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
