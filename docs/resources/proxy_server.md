---
page_title: "smc_proxy_server"
subcategory: "network_elements"
description: |-
  This represents a Proxy Server, which is a server that performs detailed examination of a connection's data and assists in the determination to allow or discard packets. Common examples include virus scanning or filtering of web URLs. Also known as content screening.
---

# smc_proxy_server (Resource)

This represents a Proxy Server, which is a server that performs detailed examination of a connection's data and assists in the determination to allow or discard packets. Common examples include virus scanning or filtering of web URLs. Also known as content screening.

## Examples

- [Proxy Server Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/network_elements/servers/proxy_server/main.tf)

Creates an HTTP proxy server configuration, supporting a variety of modes and service inspection types.

```hcl
resource "smc_proxy_server" "generic_proxy" {
  add_x_forwarded_for = true
  address             = "10.1.1.1"
  comment             = var.resource_comment
  balancing_mode      = "ha"
  fp_proxy_key_id     = 0
  http_proxy          = "generic"
  inspected_service {
    name         = "FTP"
    port         = 21
    service_type = "FTP"
  }
  inspected_service {
    name         = "HTTP"
    port         = 8080
    service_type = "HTTP"
  }
  name              = "tf_generic_proxy_server"
  ip_address        = ["10.1.1.11"]
  trust_host_header = false
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `add_x_forwarded_for` (Boolean) Indicates whether the Proxy Server adds the 'X-Forwarded-For' header in HTTP(S) requests. This is used to preserve the original client's IP address when requests are forwarded.
- `address` (String) The primary IPv4 address of the device, which is used for network communication.
- `balancing_mode` (String) The balancing mode for the Proxy Server. It can be one of the following: 'ha', 'src', 'dst', or 'srcdst'.
- `comment` (String) An optional comment for the element. This field is not required.
- `fp_proxy_key` (String) The password of the Customer ID used in HTTP/HTTPS properties for the Proxy Server. This is required at creation in clear text, and if not set during a modification, it keeps the current password.
- `fp_proxy_key_hashed` (String) A SHA512 + salted hash of the current password of the Customer ID used in HTTP/HTTPS properties for the Proxy Server. This is a read-only property and is not required at creation.
- `fp_proxy_key_id` (Number) The Key ID used in HTTP/HTTPS properties for the Proxy Server.
- `fp_proxy_user_id` (String) The Customer ID used in HTTP/HTTPS properties for the Proxy Server.
- `http_proxy` (String) The service gateway used for HTTP/HTTPS traffic in the Proxy Server. It can be one of the following: 'forcepoint_ap-web_cloud', 'generic', or 'redirect'.
- `ip_address` (List of String) The list of additional IP addresses for the Proxy Server. These addresses allow using multiple IPs for the Proxy Server element.
- `ipv6_address` (String) The primary IPv6 address of the device, which is used for network communication.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.
- `trust_host_header` (Boolean) Indicates whether the Proxy Server trusts the host header in HTTP(S) requests. This is used to determine if the server should trust the host header provided by clients.

## Nested Attributes
- `inspected_service` (List of Blocks, see [here](inspected_service.md)) The list of Inspected Services that define how the Proxy Server inspects traffic. Each service can have its own settings and requirements.
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
