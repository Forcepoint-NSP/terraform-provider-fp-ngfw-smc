---
page_title: "smc_outbound_multilink"
subcategory: "network_elements"
description: |-
  This represents an Outbound Multi-Link, which is used for load balancing outbound traffic by combining multiple NetLinks. It includes attributes for link selection method, timeout, maximum retries, and lists of Multi-Link Members and QoS Classes.
---

# smc_outbound_multilink (Resource)

This represents an Outbound Multi-Link, which is used for load balancing outbound traffic by combining multiple NetLinks. It includes attributes for link selection method, timeout, maximum retries, and lists of Multi-Link Members and QoS Classes.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `multilink_method` (String) The method for link selection in the Outbound Multi-Link. Options are 'rtt' for Round Trip Time or 'ratio' for Ratio-based distribution.
- `name` (String) Name of the object.
- `retries` (Number) The maximum number of retries for checking each NetLink in the Outbound Multi-Link.
- `timeout` (Number) The timeout in seconds for measuring NetLink performance in the Outbound Multi-Link.

## Nested Attributes
- `multilink_member` (List of Blocks, see [here](attr_multilink_member.md)) The members of the Outbound Multi-Link, which are NetLinks that can be used for load balancing.
- `multilink_qos_class` (List of Blocks, see [here](attr_multi_link_qos_class_member.md)) The QoS class members of the Outbound Multi-Link, which define quality of service parameters for load balancing.

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
