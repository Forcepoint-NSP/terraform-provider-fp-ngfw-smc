---
page_title: "scan_detection"
subcategory: ""
description: |-
  This represents the definition of Scan Detection settings on a NGFW, including scan detection type, time windows, events sensitivity, and logging options.
---

# scan_detection (Nested-Attribute)

This represents the definition of Scan Detection settings on a NGFW, including scan detection type, time windows, events sensitivity, and logging options.




## Simple Attributes
- `alert_ref` (String) This represents an abstract Alert, which is used to display messages when certain conditions are met.
- `log_level` (String) The Log Level for Scan Detection. 'none': does not create any log entry, 'transient': creates a log entry that is displayed in the Current Events mode in the Logs view (if someone is viewing it at the moment) but is not stored,'stored': creates a log entry that is stored on the Log Server, 'essential': creates a log entry that is shown in the Logs view and saved for further use. When the Log Server is unavailable, log entries are temporarily stored on the engine. When the engine is running out of space to store the log entries, it begins discarding log data in the order of importance. Monitoring data is discarded first, followed by log entries marked as Transient and Stored, and finally log entries marked as Essential. The Alert entries are the last log entries to be discarded. Note! The settings for storing the logs temporarily on the engine are defined in the log spooling policy, and 'alert': triggers the alert you add to the Alert field.
- `scan_detection_icmp_events` (Number) The ICMP Scan Events sensitivity for Scan Detection. If Scan Detection is set to Default On or Default Off, this defines the number of events needed in the selected time period to generate an alert.
- `scan_detection_icmp_timewindow` (Number) The ICMP Time window in seconds for Scan Detection. If Scan Detection is set to Default On or Default Off, this defines the number of events needed in the selected time period to generate an alert.
- `scan_detection_icmp_unit` (String) The ICMP Scan time window unit for Scan Detection. Possible values are 'second', 'minute', or 'hour'. Default is 'minute'.
- `scan_detection_tcp_events` (Number) The TCP Scan Events sensitivity for Scan Detection. If Scan Detection is set to Default On or Default Off, this defines the number of events needed in the selected time period to generate an alert.
- `scan_detection_tcp_timewindow` (Number) The TCP Time window in seconds for Scan Detection. If Scan Detection is set to Default On or Default Off, this defines the number of events needed in the selected time period to generate an alert.
- `scan_detection_tcp_unit` (String) The TCP Scan time window unit for Scan Detection. Possible values are 'second', 'minute', or 'hour'. Default is 'minute'.
- `scan_detection_type` (String) The type of Scan Detection. Possible values are 'off' (Scan Detection is not enabled), 'default off' (Scan Detection is not enabled, but you can override this setting in individual Access rules. This is the default setting), and 'default on' (Scan Detection is enabled. You can override this setting in individual Access rules if scan detection is not needed or to avoid false positives).
- `scan_detection_udp_events` (Number) The UDP Scan Events sensitivity for Scan Detection. If Scan Detection is set to Default On or Default Off, this defines the number of events needed in the selected time period to generate an alert.
- `scan_detection_udp_timewindow` (Number) The UDP Time window in seconds for Scan Detection. If Scan Detection is set to Default On or Default Off, this defines the number of events needed in the selected time period to generate an alert.
- `scan_detection_udp_unit` (String) The UDP Scan time window unit for Scan Detection. Possible values are 'second', 'minute', or 'hour'. Default is 'minute'.
- `severity` (Number) If Log Level is set to Alert, allows you to override the severity defined in the Alert element. Possible values are '1' (Info), '2-4' (Low), '5-7' (High), and '8-10' (Critical).


