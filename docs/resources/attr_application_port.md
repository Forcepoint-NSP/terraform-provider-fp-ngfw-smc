---
page_title: "application_port"
subcategory: ""
description: |-
  This represents the definition of a port for an Application element, which includes protocol, port range, TLS settings, and version compatibility.
---

# application_port (Nested-Attribute)

This represents the definition of a port for an Application element, which includes protocol, port range, TLS settings, and version compatibility.




## Simple Attributes
- `from` (Number) The minimum port number for this Application Port.
- `max_version` (String) The maximum version of the engine that supports this Application Port. If not defined, it applies to the same versions as the parent element.
- `min_version` (String) The minimum version of the engine that supports this Application Port. If not defined, it applies to the same versions as the parent element.
- `mode` (String) The mode of the Application Port, which can be 'regular' or 'quic'. By default, it is set to 'regular'.
- `protocol_ref` (String) This represents an IP-proto service, which is used to define a service based on the IP protocol number. It includes a protocol number that specifies the protocol used by the traffic.
- `tls` (String) The TLS type for this Application Port, indicating how TLS is handled. By default, it is set to 'no'.
- `to` (Number) The maximum port number for this Application Port.


