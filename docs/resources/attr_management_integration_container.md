---
page_title: "mgt_integration_container"
subcategory: ""
description: |-
  This represents the container for Management Integration entries installed on the Management Server, such as NSX integration.
---

# mgt_integration_container (Nested-Attribute)

This represents the container for Management Integration entries installed on the Management Server, such as NSX integration.




## Simple Attributes
- `admin` (String) This represents the base structure for an administrator user in the system, including permissions, status, and superuser flags.
- `comment` (String) An optional comment for the element. This field is not required.
- `enabled` (Boolean) Indicates whether the Management Integration is enabled or not.
- `integration_ip` (String) The IP address of the Management Integration, such as the NSX Manager IP.
- `integration_port` (Number) The port used by the Management Integration, default is 8090.
- `integration_type` (String) The type of the Management Integration, such as 'nsx' for NSX integration.
- `listening_ip` (String) The IP address on which the Management Integration listens for requests.
- `name` (String) Name of the object.
- `password` (String) The password used for the Management Integration, such as the NSX Manager password.
- `username` (String) The username used for the Management Integration, such as the NSX Manager username.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
