---
page_title: "inline_test"
subcategory: ""
description: |-
  This represents a test that checks whether the network settings (speed/duplex) match on the two ports that form the inline pair and can force both ports to use the same settings.
---

# inline_test (Nested-Attribute)

This represents a test that checks whether the network settings (speed/duplex) match on the two ports that form the inline pair and can force both ports to use the same settings.




## Simple Attributes
- `alert_notification` (Boolean) Indicates whether an alert notification is sent if the test fails.
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `offline_state` (Boolean) Indicates whether the test is executed when the engine is offline or not.
- `online_state` (Boolean) Indicates whether the test is executed when the engine is online or not.
- `standby_state` (Boolean) Indicates whether the test is executed when the engine is in standby state or not.
- `test_action` (String) The action to be taken if the test fails. Options include 'none', 'offline', 'forceoffline', and 'forcespeed'.
- `test_active` (Boolean) Indicates whether the test is active or not.
- `test_interval` (Number) The interval in seconds at which the test is executed. Note! Running a test too frequently can increase overhead.
- `test_timeout` (Number) The timeout in milliseconds for the inline pair link speed test.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
