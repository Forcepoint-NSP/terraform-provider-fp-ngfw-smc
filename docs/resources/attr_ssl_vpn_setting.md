---
page_title: "ssl_vpn_setting"
subcategory: "sslvpn"
description: |-
  This represents the SSL VPN settings, which include options for SSL/TLS versions, cryptography suites, and renegotiation timeout.
---

# ssl_vpn_setting (Nested-Attribute)

This represents the SSL VPN settings, which include options for SSL/TLS versions, cryptography suites, and renegotiation timeout.




## Simple Attributes
- `renegociation_timeout` (Number) The timeout for renegotiation in seconds. This is the time after which a renegotiation will be attempted.
- `ssl_3_0` (Boolean) Indicates whether SSL 3.0 is accepted for the SSL VPN connection.
- `tls13` (Boolean) 
- `tls_1_0` (Boolean) Indicates whether TLS 1.0 is enabled for the SSL VPN connection.
- `tls_1_1` (Boolean) Indicates whether TLS 1.1 is enabled for the SSL VPN connection.
- `tls_1_2` (Boolean) Indicates whether TLS 1.2 is enabled for the SSL VPN connection.
- `tls_cryptography_suite_set` (String) This represents a TLS Cryptography Suite Set Element, which contains a set of cryptographic suites used in SSL VPN configurations.


