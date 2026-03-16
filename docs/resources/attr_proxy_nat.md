---
page_title: "proxy_nat"
subcategory: ""
description: |-
  This represents the Proxy NAT (Network Address Translation) entry, a sub part of NAT Rule. Traffic is redirected to the specified Proxy Server.
---

# proxy_nat (Nested-Attribute)

This represents the Proxy NAT (Network Address Translation) entry, a sub part of NAT Rule. Traffic is redirected to the specified Proxy Server.




## Simple Attributes
- `automatic_proxy` (Boolean) Indicates whether Automatic Proxy ARP is enabled. This allows the engine to answer address queries regarding the translated address(es).
- `proxy_server` (String) This represents a Proxy Server, which is a server that performs detailed examination of a connection's data and assists in the determination to allow or discard packets. Common examples include virus scanning or filtering of web URLs. Also known as content screening.


