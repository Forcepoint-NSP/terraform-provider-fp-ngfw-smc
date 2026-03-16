---
page_title: "saml_settings"
subcategory: "admin"
description: |-
  This represents the SAML settings, including authentication method, service provider ID, and usage.
---

# saml_settings (Nested-Attribute)

This represents the SAML settings, including authentication method, service provider ID, and usage.




## Simple Attributes
- `saml_acs_url` (String) The URL where the SAML assertions are sent and processed.
- `saml_metadata_file` (String) The SAML Setting specific IdP Metadata, this overrides IdP Metadata defined in SAML Method.
- `saml_method_ref` (String) This represents an external authentication method, which can be used for user authentication in the system. It supports various types of authentication methods such as User Password, IAS, IPSec, Pre-Shared Key, RADIUS, TACACS+, and LDAP.
- `saml_name_id_policy_format` (String) The Name ID policy format used in SAML authentication. Options include 'persistent', 'transient', 'emailAddress', and 'unspecified'.
- `saml_service_provider_id` (String) The unique identifier for the SAML service provider.
- `saml_setting_usage` (String) The intended usage of the SAML settings. Options include 'bba' for user authentication, 'app_access' for Application Access Portal, and 'web_access' for Web Access.
- `saml_tls_credentials_ref` (String) This represents a TLS Server Credentials element, which is used to store the private key and certificate of an internal server. The certificate and the associated private key must be compatible with OpenSSL and be in PEM format. It is used for TLS inspection, securing Web Access Servers, and authenticating Authentication Servers.
- `saml_user_attribute` (String) The attribute used to identify the user in SAML authentication.


