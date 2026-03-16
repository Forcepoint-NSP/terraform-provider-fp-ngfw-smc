---
page_title: "smc_link_selection"
subcategory: ""
description: |-
  This represents the link selection value used in QoS settings, defining various network parameters such as bandwidth, jitter, latency, packet loss, and stability.
---

# smc_link_selection (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the link selection value used in QoS settings, defining various network parameters such as bandwidth, jitter, latency, packet loss, and stability.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `bandwidth` (Number) The bandwidth value for the link selection, ranging from 0 to 4.
- `comment` (String) An optional comment for the element. This field is not required.
- `jitter` (Number) The jitter value for the link selection, ranging from 0 to 4.
- `latency` (Number) The latency value for the link selection, ranging from 0 to 4.
- `name` (String) Name of the object.
- `packet_loss` (Number) The packet loss value for the link selection, ranging from 0 to 4.
- `stability` (Number) The stability value for the link selection, ranging from 0 to 4.


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
