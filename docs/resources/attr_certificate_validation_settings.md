---
page_title: "certificate_validation_settings"
subcategory: ""
description: |-
  This defines settings for certificate validation on the engine, including OCSP and CRL lookups, timeout settings, and optional HTTP proxy configuration.
---

# certificate_validation_settings (Nested-Attribute)

This defines settings for certificate validation on the engine, including OCSP and CRL lookups, timeout settings, and optional HTTP proxy configuration.




## Simple Attributes
- `crl_distribution_point_list` (List of String) A list of URLs representing the distribution points to use for CRL (Certificate Revocation List) cache on the engine.
- `http_proxy_ref` (String) This represents an HTTP Proxy, which is used to route HTTP traffic through a proxy server. It includes attributes for IP address, port, username, and password.
- `ocsp_crl_timeout` (Number) The timeout in seconds for OCSP and CRL lookups. This setting determines how long the engine will wait for a response before timing out.


