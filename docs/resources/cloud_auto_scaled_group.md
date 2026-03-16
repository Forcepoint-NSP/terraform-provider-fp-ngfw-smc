---
page_title: "smc_scale_set"
subcategory: ""
description: |-
  This represents a Cloud Auto-Scaled Group, which is used to manage a group of cloud firewalls that can scale automatically based on load. It includes configurations for monitoring, installed policies, and the elements within the group.
---

# smc_scale_set (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Cloud Auto-Scaled Group, which is used to manage a group of cloud firewalls that can scale automatically based on load. It includes configurations for monitoring, installed policies, and the elements within the group.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `element` (List of String) URI of the Cloud Firewall which belongs to the Group.
- `installed_policy` (String) The installed policy found on the first cloud firewall in the group.
- `is_monitored` (Boolean) When selected, adds the Cloud Auto-Scaled Group to the System Status tree. This is not required and defaults to true.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.


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
