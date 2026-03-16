---
page_title: "tls_settings"
subcategory: ""
description: |-
  This represents the TLS settings for Elasticsearch/Log Forwarding, including whether to use internal credentials and the TLS server credentials.
---

# tls_settings (Nested-Attribute)

This represents the TLS settings for Elasticsearch/Log Forwarding, including whether to use internal credentials and the TLS server credentials.




## Simple Attributes
- `tls_credentials` (String) This represents a TLS Server Credentials element, which is used to store the private key and certificate of an internal server. The certificate and the associated private key must be compatible with OpenSSL and be in PEM format. It is used for TLS inspection, securing Web Access Servers, and authenticating Authentication Servers.
- `use_internal_credentials` (Boolean) Indicates whether to use the server's internal TLS credentials for Elasticsearch/Log Forwarding. If true, internal credentials are used; if false, tlsServerCredentials are used.


