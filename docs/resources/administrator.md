---
page_title: "smc_administrator"
subcategory: "admin"
description: |-
  This represents an administrator user. It contains the user's credentials, authentication method, and permissions for managing the system. Administrators can be local or LDAP-based, and they can have sudo rights on engines.
---

# smc_administrator (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an administrator user. It contains the user's credentials, authentication method, and permissions for managing the system. Administrators can be local or LDAP-based, and they can have sudo rights on engines.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `allow_sudo` (Boolean) Indicates whether the administrator has sudo rights on the engine. If true, the administrator can execute commands with elevated privileges on the engine.
- `allowed_to_login_in_shared` (Boolean) Indicates whether the administrator is allowed to log in to the Shared Domain. If true, the administrator can access shared resources and settings.
- `auth_method` (String) This represents an external authentication method, which can be used for user authentication in the system. It supports various types of authentication methods such as User Password, IAS, IPSec, Pre-Shared Key, RADIUS, TACACS+, and LDAP.
- `can_use_api` (Boolean) Indicates whether the administrator can use the SMC API. If true, the administrator can access the API for automation and integration purposes.
- `comment` (String) An optional comment for the element. This field is not required.
- `console_superuser` (Boolean) Indicates whether the administrator user is a console superuser. This applies only to an Appliance and grants additional privileges for console access.
- `engine_target` (List of String) URI of the target engine.
- `ldap_group` (String) This represents a group of external LDAP Users defined on the external LDAP server. It allows browsing and invalidating user entries.
- `ldap_user` (String) This represents an external LDAP User defined on the external LDAP server. It contains user details such as display name, email, phone number, job title, office location, frame IP, activation date, expiration date, and days left until expiration.
- `local_admin` (Boolean) Indicates whether the administrator is a local administrator on the engine. If true, the administrator can log in to the engine directly.
- `name` (String) Name of the object.
- `password` (String) The password of the administrator user. This is required for local administrators and optional for LDAP-based administrators.
- `superuser` (Boolean) Indicates whether the administrator user is a superuser. A superuser has full access to all system functions and settings.

## Nested Attributes
- `client_identity` (Single Block, see [here](attr_tls_identity.md)) 
- `permissions` (Single Block, see [here](attr_admin_permissions.md)) 

## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `enabled` (Boolean) The status of the administrator user. It indicates whether the user is enabled or disabled.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `is_user_locked` (Boolean) Indicates whether the administrator user is locked. A locked user cannot log in until it is unlocked by an administrator.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
