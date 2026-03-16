---
page_title: "smc_ssl_vpn_portal"
subcategory: "sslvpn"
description: |-
  This represents an Application Access Portal. It contains settings for hostnames, policies, server credentials, and various portal configurations.
---

# smc_ssl_vpn_portal (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Application Access Portal. It contains settings for hostnames, policies, server credentials, and various portal configurations.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `allow_empty_referrer` (Boolean) Indicates whether the Application Access Portal allows empty referrers.
- `brand_color` (String) The brand color of the Application Access Portal. This is used for branding purposes.
- `comment` (String) An optional comment for the element. This field is not required.
- `log_access` (Boolean) Indicates whether access logs are enabled for the Application Access Portal.
- `look_and_feel` (String) The look and feel of the Application Access Portal.
- `name` (String) Name of the object.
- `persistent_session` (Boolean) Indicates whether the Application Access Portal supports persistent sessions.
- `portal_session_timeout` (Number) The session timeout for the Application Access Portal in seconds.
- `portal_theme` (String) The theme of the Application Access Portal.
- `portal_timeout` (Number) The timeout for the Application Access Portal in seconds.
- `self_signed_certificate` (Boolean) Indicates whether the Application Access Portal uses a self-signed certificate.
- `server_credential` (List of String) URI of the TLS Server Credential.
- `session_timeout_unit` (String) The unit of time for the session timeout.
- `ssl_vpn_policy` (String) This represents an Application Access Portal Policy, which is used to define portal access rules for the connections. It includes properties such as portal rules and access control.
- `timeout_unit` (String) The unit of time for the portal timeout.
- `title` (String) The title of the Application Access Portal.

## Nested Attributes
- `ssl_vpn_hostname` (List of Blocks, see [here](attr_ssl_vpn_hostname.md)) The Application Access Hostname elements.

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
