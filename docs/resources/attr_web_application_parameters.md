---
page_title: "web_app"
subcategory: ""
description: |-
  This represents the parameters of a Web Application, including its identifier, enabled status, listening address, host name, port, context path, and security settings.
---

# web_app (Nested-Attribute)

This represents the parameters of a Web Application, including its identifier, enabled status, listening address, host name, port, context path, and security settings.




## Simple Attributes
- `client_cert_auth_enabled` (Boolean) Indicates whether client authentication is enabled for the Web Application.
- `context_path` (String) The context path for the Web Application. If null, it uses the root context.
- `eca_enabled` (Boolean) Indicates whether ECA Rollout evaluation is enabled in SMC Download pages.
- `enable_saml_for_web_access` (Boolean) Indicates whether SAML authentication is enabled for Web Access.
- `enabled` (Boolean) Indicates whether the Web Application is enabled or not.
- `host_name` (String) The host name for the Web Application. If null, no specific host name is set.
- `listening_address` (String) The address on which the Web Application listens for incoming connections. If null, it listens on all addresses.
- `log_access` (Boolean) Indicates whether access to this Web Application is logged.
- `port` (Number) The port on which the Web Application listens for incoming connections.
- `server_credentials_ref` (String) This represents a TLS Server Credentials element, which is used to store the private key and certificate of an internal server. The certificate and the associated private key must be compatible with OpenSSL and be in PEM format. It is used for TLS inspection, securing Web Access Servers, and authenticating Authentication Servers.
- `session_timeout` (Number) The session timeout in seconds for the Web Application, used to control user session duration.
- `ssl_session_id` (Boolean) Indicates whether the SSL session ID is used for the Web Application.
- `standalone_enabled` (Boolean) Indicates whether Standalone client bundles download are enabled in SMC Download pages.
- `tls_cipher_suites` (String) This represents a TLS Cryptography Suite Set Element, which contains a set of cryptographic suites used in SSL VPN configurations.
- `web_app_identifier` (String) The unique identifier for the Web Application, used to reference it in configurations.
- `webclient_enabled` (Boolean) Indicates whether the Web UI (React) is enabled in the API page.
- `xvfb_path` (String) The full path to the directory containing the xvfb-run application, used for Web Access web app.

## Nested Attributes
- `saml_settings` (List of Blocks, see [here](attr_saml_settings.md)) The SAML settings associated with the Web Application.

