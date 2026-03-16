---
page_title: "third_party_monitoring"
subcategory: ""
description: |-
  This represents the monitoring settings for third-party devices, including SNMP traps, NetFlow, syslog, and status monitoring.
---

# third_party_monitoring (Nested-Attribute)

This represents the monitoring settings for third-party devices, including SNMP traps, NetFlow, syslog, and status monitoring.




## Simple Attributes
- `encoding` (String) The character encoding used for log reception, which determines how text is interpreted.
- `logging_profile_ref` (String) This represents a Logging Profile used for Third Party Monitoring, defining how the Log Server converts Syslog data received from a particular type of third-party component into Stonesoft Management Center log entries.
- `monitoring_log_server_ref` (String) This represents a Log Server, which is a component of the Management Center responsible for storing and managing log (and alert) data, and analyzing and correlating events detected by multiple NGFW Engines.
- `netflow` (Boolean) Indicates whether NetFlow data reception is enabled for this device.
- `probing_profile_ref` (String) This represents a Probing Profile used in Third Party Monitoring. It contains settings that define how a Log Server monitors third-party components.
- `snmp_trap` (Boolean) Indicates whether SNMP trap reception is enabled for this device.
- `time_zone` (String) The time zone ID for log reception, which determines the time zone used for timestamps in logs.


