---
page_title: "connection_timeout"
subcategory: ""
description: |-
  This represents the Idle Timeout settings for a protocol or TCP connection state, which defines the timeout duration for idle connections.
---

# connection_timeout (Nested-Attribute)

This represents the Idle Timeout settings for a protocol or TCP connection state, which defines the timeout duration for idle connections.




## Simple Attributes
- `protocol` (String) The protocol for which the idle timeout is defined. Default values are 'udp', 'tcp', 'icmp', and 'other'. Advanced values are 'related', 'tcp_closing', 'tcp_syn_seen', 'tcp_fin_wait_1', 'tcp_fin_wait_2', 'tcp_time_wait', 'tcp_close_wait', 'tcp_last_ack', 'tcp_syn_ack_seen', 'tcp_time_wait_ack', 'tcp_closing_ack', 'tcp_close_wait_ack', 'tcp_last_ack_wait', 'tcp_syn_fin_seen', 'ipsec_established', 'tcp_syn_return', 'loose_graceful', 'loose'.
- `timeout` (Number) The timeout duration in seconds for idle connections.


