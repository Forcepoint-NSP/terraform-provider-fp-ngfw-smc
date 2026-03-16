---
page_title: "netflow_collector"
subcategory: ""
description: |-
  This represents a Netflow collector, which is a sub part of Log Server entity. It allows you to forward log data to external hosts in various formats and with specific filters.
---

# netflow_collector (Nested-Attribute)

This represents a Netflow collector, which is a sub part of Log Server entity. It allows you to forward log data to external hosts in various formats and with specific filters.




## Simple Attributes
- `data_context` (String) This represents the Data Context. It contains a tag that is used to identify the data context in event filtering operations.
- `filter` (String) This represents a container for filter expressions, which can be used to define complex filtering rules. It contains a root node that holds the main filter expression.
- `host` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `kafka_topic` (String) The Kafka Topic used for forwarding logs through Kafka. This is only applicable when the service is set to 'kafka'.
- `netflow_collector_port` (Number) The port used for log forwarding.
- `netflow_collector_service` (String) The network protocol for forwarding the log data, such as 'udp', 'tcp', or 'tcp_with_tls'.
- `netflow_collector_version` (String) The format for forwarding the log data, such as 'cef', 'csv', 'leef', 'netflow_v11', 'ipfix', 'xml', or 'esm'.
- `tls_profile` (String) This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.

## Nested Attributes
- `tls_identity` (Single Block, see [here](attr_tls_identity.md)) 

