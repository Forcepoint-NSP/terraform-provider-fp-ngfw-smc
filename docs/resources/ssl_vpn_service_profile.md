---
page_title: "smc_ssl_vpn_service_profile"
subcategory: "sslvpn"
description: |-
  This represents an Application Access Service Profile. It contains settings for exceptions, HTTP fields, authentication type, NTLM support, and user input formats.
---

# smc_ssl_vpn_service_profile (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Application Access Service Profile. It contains settings for exceptions, HTTP fields, authentication type, NTLM support, and user input formats.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `authentication_type` (String) The type of authentication used in the Application Access service profile.
- `comment` (String) An optional comment for the element. This field is not required.
- `cookie_hiding` (String) The mode for hiding cookies in the Application Access service profile.
- `login_page_url` (String) The URL of the login page for the Application Access service profile.
- `name` (String) Name of the object.
- `ntlm_support` (Boolean) Indicates whether NTLM support is enabled in the Application Access service profile.
- `password_input_name` (String) The name of the password input field in the Application Access service profile.
- `post_request_url` (String) The URL to which the post request is sent in the Application Access service profile.
- `user_input_custom_format` (String) The custom format for user input in the Application Access service profile.
- `user_input_format` (String) The format of user input in the Application Access service profile.
- `user_input_name` (String) The name of the user input field in the Application Access service profile.

## Nested Attributes
- `ssl_vpn_http_field` (List of Blocks, see [here](attr_ssl_vpn_http_field.md)) The Application Access HTTP fields.
- `ssl_vpn_profile_exception` (List of Blocks, see [here](attr_ssl_vpn_hostname.md)) The Application Access Hostname exceptions.

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
