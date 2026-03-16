---
page_title: "multiping_test"
subcategory: ""
description: |-
  This represents the definition of a Multiping test, which sends out a series of ping requests to determine connectivity through a network link.
---

# multiping_test (Nested-Attribute)

This represents the definition of a Multiping test, which sends out a series of ping requests to determine connectivity through a network link.




## Simple Attributes
- `alert_notification` (Boolean) Indicates whether an alert notification is sent if the test fails.
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `nicid` (String) The Interface ID on which the test is run. It can be 'all', 'all_with_cvi', or a specific interface ID. Note! (Firewalls only) Only the first interface that belongs to an Aggregated Link is shown in the list of interfaces. However, the Link Status test checks the status of both interfaces in the Aggregated Link.
- `offline_state` (Boolean) Indicates whether the test is executed when the engine is offline or not.
- `online_state` (Boolean) Indicates whether the test is executed when the engine is online or not.
- `retry_count` (Number) The number of times the Multiping test will retry if it fails. Note! We recommend always setting the retry count to more than 1 to avoid creating overly sensitive tests that burden the system unnecessarily.
- `standby_state` (Boolean) Indicates whether the test is executed when the engine is in standby state or not.
- `test_action` (String) The action to be taken if the test fails. Options include 'none', 'offline', 'forceoffline', and 'forcespeed'.
- `test_active` (Boolean) Indicates whether the test is active or not.
- `test_interval` (Number) The interval in seconds at which the test is executed. Note! Running a test too frequently can increase overhead.
- `test_timeout` (Number) The timeout in milliseconds for the Multiping test. If the test does not return a response within this time, it is considered to have failed.
- `value` (List of String) The target addresses for the Multiping test. Only IPv4 addresses are supported. The test is considered failed if none of the target addresses responds.

## Nested Attributes
- `enabled_interface` (List of Blocks, see [here](attr_enabled_interface_entry.md)) The source interfaces for the Multiping test. If not specified, the source IP address is selected automatically based on standard routing rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
