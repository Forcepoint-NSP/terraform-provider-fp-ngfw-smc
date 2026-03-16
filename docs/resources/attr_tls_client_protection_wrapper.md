---
page_title: "tls_client_protection"
subcategory: ""
description: |-
  This represents a wrapper for TLS Client protection settings, including proxy usage and trusted certificate authorities.
---

# tls_client_protection (Nested-Attribute)

This represents a wrapper for TLS Client protection settings, including proxy usage and trusted certificate authorities.




## Simple Attributes
- `ca_for_signing_ref` (String) This represents a Signing Certificate Authority, which is used to manage signing certificate authorities in the system.
- `proxy_usage` (String) The purpose of the proxy usage, such as 'tls_inspection' or 'opcua_inspection'. This field indicates the intended use of the proxy in TLS Client protection settings.
- `tls_trusted_ca_ref` (List of String) URI of the TLS trusted CA.
- `tls_trusted_ca_tag_ref` (List of String) URI of the trusted CA tag.


