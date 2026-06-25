---
page_title: "application_access_allowed_url"
subcategory: "applicationaccess"
description: |-
  This defines an allowed URL in an Application Access, including the protocol, host, and port. It is used to specify which URLs are permitted for access through the SSL VPN.
---

# application_access_allowed_url (Nested-Attribute)

This defines an allowed URL in an Application Access, including the protocol, host, and port. It is used to specify which URLs are permitted for access through the SSL VPN.




## Simple Attributes
- `port` (String) The port number for the allowed URL, which is typically 80 for HTTP or 443 for HTTPS. 'any' can be used to allow all ports.
- `protocol` (String) The protocol used for the allowed URL, such as 'HTTP', 'HTTPS', etc.
- `url_host` (String) The host of the allowed URL, which can be a domain name or an IP address.


