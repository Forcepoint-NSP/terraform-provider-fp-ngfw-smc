---
page_title: "smc_ssl_vpn_web_service"
subcategory: "sslvpn"
description: |-
  This represents an Application Access Web Service. It contains settings for alternative hostnames, allowed URLs, trusted certificate authorities, and various service configurations.
---

# smc_ssl_vpn_web_service (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Application Access Web Service. It contains settings for alternative hostnames, allowed URLs, trusted certificate authorities, and various service configurations.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `client_trust` (String) This represents a Trusted Certificate Authority, which is used to manage trusted certificate authorities in the system.
- `comment` (String) An optional comment for the element. This field is not required.
- `cookie_protection` (Boolean) Indicates whether cookie protection is enabled for the Application Access Web Service. This helps to prevent session hijacking by securing cookies.
- `description` (String) Description of the Application Access Web Service.
- `disable_rewrite` (Boolean) Indicates whether the rewrite feature is disabled for the Application Access Web Service. If true, it means that HTML content rewriting is not applied.
- `external_url` (String) External URL of the Application Access Web Service. This is the URL that users will access from outside the network.
- `internal_url` (String) Internal URL of the Application Access Web Service. This is the URL that users will access from inside the network.
- `name` (String) Name of the object.
- `rewrite_html` (Boolean) Indicates whether HTML content should be rewritten for the Application Access Web Service. This is useful for modifying web pages to work correctly over the VPN.
- `routing_method` (String) The routing method for the Application Access Web Service. It defines how the service routes traffic, such as DNS mapping or direct routing.
- `self_signed_certificate` (Boolean) Indicates whether the Application Access Web Service uses a self-signed certificate. If true, it means the service is using a self-signed certificate for secure communication.
- `server_credential` (String) This represents a TLS Server Credentials element, which is used to store the private key and certificate of an internal server. The certificate and the associated private key must be compatible with OpenSSL and be in PEM format. It is used for TLS inspection, securing Web Access Servers, and authenticating Authentication Servers.
- `ssl_vpn_service_profile` (String) This represents an Application Access Service Profile. It contains settings for exceptions, HTTP fields, authentication type, NTLM support, and user input formats.
- `ssl_vpn_sso_domain` (String) This represents an SSL VPN Single Sign On Domain, which includes properties such as SSO mode and timeout settings.
- `start_page` (String) The start page for the Application Access Web Service. This is the page that users will see when they first access the service.
- `title` (String) Title of the Application Access Web Service.
- `trusted_ca` (List of String) URI of the trusted CA.
- `url_prefix` (String) URL prefix for the Application Access Web Service. This is used to define a base path for the service URLs.
- `visible_in_portal` (Boolean) Indicates whether the Application Access Web Service is visible in the portal. If true, it will be displayed in the user portal for easy access.

## Nested Attributes
- `ssl_vpn_allowed_url` (List of Blocks, see [here](attr_ssl_vpn_allowed_url.md)) List of allowed URLs for the Application Access Web Service.
- `ssl_vpn_althost` (List of Blocks, see [here](attr_ssl_vpn_hostname.md)) List of alternative hostnames for the Application Access Web Service.

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
