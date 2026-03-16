---
page_title: "smc_smtp_server"
subcategory: "network_elements"
description: |-
  This represents a Simple Mail Transfer Protocol (SMTP) server, which is used to process notifications by e-mails. It includes attributes for port, default sender email, and default sender name.
---

# smc_smtp_server (Resource)

This represents a Simple Mail Transfer Protocol (SMTP) server, which is used to process notifications by e-mails. It includes attributes for port, default sender email, and default sender name.

## Examples

- [SMTP Server Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/network_elements/servers/radius_server/main.tf)

Defines an SMTP server for alerting or notification policies.

```hcl
resource "smc_smtp_server" "smtp" {
  address              = "10.1.1.123"
  email_sender_address = "sender@example.com"
  email_sender_name    = "example_sender"
  name                 = "tf_smtp"
  comment              = var.resource_comment
  port                 = 25
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the device, which is used for network communication.
- `comment` (String) An optional comment for the element. This field is not required.
- `email_sender_address` (String) The default sender email address used in the From field of the e-mail. This default value can be overridden in the properties of the element where the SMTP Server is used.
- `email_sender_name` (String) The default sender name used in the From field of the e-mail. This default value can be overridden in the properties of the element where the SMTP Server is used.
- `ipv6_address` (String) The primary IPv6 address of the device, which is used for network communication.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `port` (Number) The port number for the SMTP server.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.

## Nested Attributes
- `third_party_monitoring` (Single Block, see [here](attr_third_party_monitoring.md)) 

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
