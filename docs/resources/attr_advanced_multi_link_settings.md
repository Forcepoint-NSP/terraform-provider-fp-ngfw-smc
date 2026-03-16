---
page_title: "advanced_multilink_settings"
subcategory: ""
description: |-
  This represents the advanced Multi-Link settings for the engine, which includes failover thresholds, probe thresholds, and initial values for various parameters.
---

# advanced_multilink_settings (Nested-Attribute)

This represents the advanced Multi-Link settings for the engine, which includes failover thresholds, probe thresholds, and initial values for various parameters.




## Simple Attributes
- `activate_failover` (Number) The time in seconds to activate failover after the threshold is reached.
- `dpd_initial_values` (String) The initial values for DPD (Dead Peer Detection) settings, which can be used to configure the initial state of the Multi-Link.
- `failover_initial_values` (String) The initial values for failover settings, which can be used to configure the initial state of the Multi-Link.
- `failover_prop_high` (Number) The failover property high, which is a percentage value that indicates the threshold for high failover conditions.
- `failover_threshold` (Number) The failover threshold, which determines the number of failed probes before a failover is triggered.
- `in_traffic_initial_values` (String) The initial values for in-traffic settings, which can be used to configure the initial state of the Multi-Link.
- `latency_initial_values` (String) The initial values for latency settings, which can be used to configure the initial state of the Multi-Link.
- `multilink_extra_options` (String) Additional options for Multi-Link settings, which can be used to configure extra parameters for the Multi-Link.
- `pdr_initial_values` (String) The initial values for PDR (Packet Delivery Ratio) settings, which can be used to configure the initial state of the Multi-Link.
- `probe_threshold` (Number) The probe threshold, which is a percentage value that indicates the threshold for probing the links.
- `remoteload_initial_values` (String) The initial values for remote load settings, which can be used to configure the initial state of the Multi-Link.
- `throughput_initial_values` (String) The initial values for throughput settings, which can be used to configure the initial state of the Multi-Link.


