---
page_title: "smc_dns_relay_profile"
subcategory: "engines"
description: |-
  This represents a DNS Relay profile element, which contains settings for hostname mappings, domain-specific DNS servers, fixed domain answers, and DNS answer translations.
---

# smc_dns_relay_profile (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a DNS Relay profile element, which contains settings for hostname mappings, domain-specific DNS servers, fixed domain answers, and DNS answer translations.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `dns_answer_translation` (List of Blocks, see [here](attr_dns_rp_dns_answer_translation.md)) The DNS answer translation, which allows for translation of DNS answers based on specific rules.
- `domain_specific_dns_server` (List of Blocks, see [here](attr_domain_specific_dns_server.md)) The domain-specific DNS server, which allows for different DNS servers to be used based on the domain.
- `fixed_domain_answer` (List of Blocks, see [here](attr_dns_rp_fixed_domain_answer.md)) The fixed domain answer, which provides static DNS responses for specific domains.
- `hostname_mapping` (List of Blocks, see [here](attr_dns_rp_hostname_mapping.md)) The hostname mapping, which maps hostnames to IP addresses for DNS resolution.

## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
