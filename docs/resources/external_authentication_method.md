---
page_title: "smc_authentication_service"
subcategory: ""
description: |-
  This represents an external authentication method, which can be used for user authentication in the system. It supports various types of authentication methods such as User Password, IAS, IPSec, Pre-Shared Key, RADIUS, TACACS+, and LDAP.
---

# smc_authentication_service (Resource)

This represents an external authentication method, which can be used for user authentication in the system. It supports various types of authentication methods such as User Password, IAS, IPSec, Pre-Shared Key, RADIUS, TACACS+, and LDAP.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/network_elements/servers/tacacs_server) for a complete minimal example

This example creates an Authentication Service of type TACACS for use with authentication servers.

```hcl
resource "smc_authentication_service" "tf_tacacs_auth" {
  name = "tf_tacacs_auth"
  type = "tacacs"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `open_id_client_id` (String) The Open ID Client ID or API Key, which is used to authenticate the client with the Open ID server.
- `open_id_secret` (String) The Open ID shared secret, which is used to authenticate the client with the Open ID server. This value is ciphered and should be stored securely.
- `open_id_trusted_cert_ref` (String) This represents a Trusted Certificate Authority, which is used to manage trusted certificate authorities in the system.
- `open_id_url` (String) The Open ID Connect Discovery URL, which is used to discover the Open ID server configuration.
- `open_id_user_attribute` (String) The Open ID User Attribute, which is used to identify the user in the Open ID server response. If empty, the default user identifier will be used.
- `saml_group_attr_name` (List of String) The list of group attribute names used in SAML authentication.
- `saml_metadata_file` (String) The SAML Identity Provider Metadata XML File, which contains the metadata of the SAML Identity Provider. This can be a local file path or a remote URI.
- `saml_tls_profile_ref` (String) This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.
- `type` (String) Type of the authentication method, which can be one of the following: user_password, ias, ipsec, preshared_key, radius, tacacs, ldap.

## Nested Attributes
- `authentication_server` (List of Blocks, see [here](attr_authentication_method_container_for_storage.md)) The authentication servers using this authentication method, which can be RADIUS, TACACS+, LDAP, OpenID, or SAML servers.

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
