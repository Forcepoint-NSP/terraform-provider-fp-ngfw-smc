---
page_title: "smc_web_access_server"
subcategory: ""
description: |-
  This represents a Web Access Server, which is used to provide web-based access to the management interface. It includes attributes for log target server, web application parameters, certificate settings, TLS profile, and UIID.
---

# smc_web_access_server (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Web Access Server, which is used to provide web-based access to the management interface. It includes attributes for log target server, web application parameters, certificate settings, TLS profile, and UIID.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the server. Must be a valid IPv4 address.
- `alert_server` (String) This represents a Log Server, which is a component of the Management Center responsible for storing and managing log (and alert) data, and analyzing and correlating events detected by multiple NGFW Engines.
- `comment` (String) An optional comment for the element. This field is not required.
- `ipv6_address` (String) The primary IPv6 address of the server. Must be a valid IPv6 address.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `tls_profile` (String) This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.
- `uiid` (String) The UIID (Unique Installation ID) for the Web Access Server, which uniquely identifies the installation.

## Nested Attributes
- `external_pki_certificate_settings` (Single Block, see [here](attr_certificate_settings.md)) 
- `third_party_monitoring` (Single Block, see [here](attr_third_party_monitoring.md)) 
- `web_app` (List of Blocks, see [here](attr_web_application_parameters.md)) The list of Web Application parameters for the Web Access Server. Each application can have its own settings and requirements.

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
