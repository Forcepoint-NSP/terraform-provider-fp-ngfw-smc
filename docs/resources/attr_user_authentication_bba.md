---
page_title: "web_authentication"
subcategory: ""
description: |-
  This represents the Browser-Based User Authentication settings for a NGFW, including HTTP/HTTPS ports, session handling, and TLS profile.
---

# web_authentication (Nested-Attribute)

This represents the Browser-Based User Authentication settings for a NGFW, including HTTP/HTTPS ports, session handling, and TLS profile.




## Simple Attributes
- `all_interfaces` (Boolean) Indicates if the service listens on all interfaces. If true, it listens on all available network interfaces.
- `authentication_idle_timeout` (Number) The authentication idle timeout in seconds. This defines how long a user session remains active without activity before requiring re-authentication.
- `authentication_timeout` (Number) The authentication timeout in seconds. This defines how long a user session remains active before requiring re-authentication.
- `enforce_https` (Boolean) Indicates if HTTPS must be enforced. If true, all HTTP requests are redirected to HTTPS.
- `http_port` (Number) The port on which HTTP listens. If null, HTTP is not enabled.
- `https_port` (Number) The port on which HTTPS listens. If null, HTTPS is not enabled.
- `keep_alive_rate` (Number) The rate at which the service refreshes status, in seconds. If null, keep-alive is disabled.
- `key_length` (Number) The key length for the certificate, specified in bits. This is used for certificate-based authentication.
- `page_ref` (String) This represents a Web Authentication HTML Page, which is used for browser-based user authentication. It contains the HTML content that is displayed to users when they are not authorized.
- `session_handling` (Boolean) Indicates if session handling is enabled. If true, sessions are managed for user authentication.
- `subject_name` (String) The subject name for the certificate. This is used for certificate-based authentication.
- `tls_profile` (String) This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.
- `use_cert_bba` (Boolean) Indicates if certificate-based authentication is used. If true, users authenticate using their certificates.

## Nested Attributes
- `enabled_interface` (List of Blocks, see [here](attr_enabled_interface_entry.md)) The list of interfaces on which the service listens. Each entry represents an enabled interface.

