---
page_title: "tls_identity"
subcategory: ""
description: |-
  This represents a TLS Identity, which contains data to check server identity when establishing a TLS connection. It includes fields such as DNS name, IP address, common name, distinguished name, and various hash values.
---

# tls_identity (Nested-Attribute)

This represents a TLS Identity, which contains data to check server identity when establishing a TLS connection. It includes fields such as DNS name, IP address, common name, distinguished name, and various hash values.




## Simple Attributes
- `tls_field` (String) The field used to check the identity, which can be one of the following: DNSName, IPAddress, CommonName, DistinguishedName, SHA-1, SHA-256, SHA-512, MD5, Email.
- `tls_value` (String) The value to check the selected field with, such as a DNS name, IP address, or hash value.


