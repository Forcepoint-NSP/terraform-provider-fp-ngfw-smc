---
page_title: "dns_sinkhole_rules"
subcategory: ""
description: |-
  This represents the definition of a DNS Sinkhole Rule, which includes a domain name and the response for DNS queries.
---

# dns_sinkhole_rules (Nested-Attribute)

This represents the definition of a DNS Sinkhole Rule, which includes a domain name and the response for DNS queries.




## Simple Attributes
- `domain_name` (String) This is the base class for all storable elements.
- `response` (String) The response for the DNS query, which can be 'NXDOMAIN' (for error case), an IPv4 address, or an IPv6 address.


