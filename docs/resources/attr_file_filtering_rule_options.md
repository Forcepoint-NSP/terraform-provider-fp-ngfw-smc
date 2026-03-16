---
page_title: "options"
subcategory: "policy"
description: |-
  This represents the File Filtering Rule Options, which are used in file filtering rules.
---

# options (Nested-Attribute)

This represents the File Filtering Rule Options, which are used in file filtering rules.




## Simple Attributes
- `application_logging` (String) Stores information about Application use, allowing for application-specific logging and monitoring.'off' to disable the option; 'default' to enable the option; 'enforced' to enforce the option.
- `eia_executable_logging` (String) Stores information about EIA Executable use, allowing for executable-specific logging and monitoring.'off' to disable the option; 'default' to enable the option; 'enforced' to enforce the option.
- `log_accounting_info_mode` (Boolean) Specifies whether both connection opening and closing are logged, and information on the volume of traffic is collected. This option is not available for rules that issue Alerts.
- `log_alert` (String) This represents an abstract Alert, which is used to display messages when certain conditions are met.
- `log_closing_mode` (Boolean) Specifies whether log entries are created when connections are closed. If true, both connection opening and closing are logged, but no traffic volume information is collected.
- `log_compression` (String) Specifies the log compression mode, which can be 'off' to not compress logs, 'only_access' to compress only Access Logs, or 'also_inspection' to compress Inspection Logs.
- `log_compression_max_burst_size` (Number) The maximum burst size for compressed logs, which limits the number of log entries that can be created in a burst.
- `log_compression_max_log_rate` (Number) The maximum log rate for compressed logs, which limits the number of log entries that can be created per second.
- `log_level` (String) The Log Level for the rule, which determines how matching packets are logged or alerted.
- `log_payload_excerpt` (Boolean) Stores an excerpt of the packet that matched, allowing quick viewing of the payload in the Logs view. The maximum recorded excerpt size is 4 KB.
- `log_payload_record` (Boolean) Records the traffic up to the limit set in the Record Length field, allowing for detailed analysis of the traffic.
- `log_severity` (Number) The severity level for the log entry when the Log Level is set to Alert, overriding the severity defined in the Alert element.
- `qos_class` (String) This represents a QoS Class, which is an element that links a rule in a QoS Policy to one or more Firewall Actions. The traffic allowed in the access rule is assigned the QoS Class defined for the rule, and the QoS class is used as the matching criteria for applying QoS Policy rules.
- `quarantine_malicious_files` (Boolean) Indicates whether to quarantine malicious files.
- `url_category_logging` (String) Stores information about URL Category use, allowing for URL category-specific logging and monitoring.'off' to disable the option; 'default' to enable the option; 'enforced' to enforce the option.
- `user_logging` (String) Stores information about Users when they are used as the Source or Destination of an Access rule, allowing for user-specific logging and monitoring.'off' to disable the option; 'default' to enable the option; 'enforced' to enforce the option.


