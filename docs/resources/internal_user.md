---
page_title: "smc_internal_user"
subcategory: "admin"
description: |-
  This represents an Internal User. It contains user details such as display name, days left until expiration, user groups, password, and pre-shared key.
---

# smc_internal_user (Resource)

This represents an Internal User. It contains user details such as display name, days left until expiration, user groups, password, and pre-shared key.

## Examples

- [fw_policy_with_match_expression/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/policies/fw_policy_with_match_expression/main.tf): Adds a test internal SMC user account for authentication testing in policies.

This snippet defines an `smc_internal_user` resource representing a local user with a password for test login/auth policy scenarios.

```hcl
resource "smc_internal_user" "todd" {
  authentication_method = [data.smc_href.user_password_auth_method.id]
  days_left             = -1
  name                  = "tf_user"
  password              = "SuperSecretPassword"
  unique_id             = "cn=tf_user,dc=stonegate,domain=InternalDomain"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `activation_date` (Number) The optional activation date in milliseconds since epoch.
- `authentication_method` (List of String) URI of the authentication method.
- `comment` (String) An optional comment for the element. This field is not required.
- `days_left` (Number) The number of days left until the user expires or is activated.
- `display_name` (String) The display name of the user.
- `expiration_after` (Number) The optional expiration delay in days. If not set, the user will never expire.
- `expiration_date` (Number) The optional expiration date in milliseconds since epoch. If not set, the user will never expire.
- `name` (String) Name of the object.
- `password` (String) The password of the user in case of password authentication method.
- `pre_shared_key` (String) The pre-shared key of the user in case of pre-shared key authentication method.
- `subject_alt_names` (String) The subject alternative names required in case of IPSec certificate authentication method. This is a comma-separated list of subject alternative names.
- `unique_id` (String) The unique id of the User/User Group element.
- `user_domain` (String) This represents a User Domain, which is used to define the authentication domain for users. It can be either an authentication domain or the internal domain.
- `user_group` (List of String) URI of the user group.


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
