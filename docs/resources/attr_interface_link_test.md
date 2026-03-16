---
page_title: "linkstatus_test"
subcategory: ""
description: |-
  This represents the definition of an Interface Link test, which checks whether a network port reports the link as up or down. The link status depends on the type of interface and its configuration.
---

# linkstatus_test (Nested-Attribute)

This represents the definition of an Interface Link test, which checks whether a network port reports the link as up or down. The link status depends on the type of interface and its configuration.




## Simple Attributes
- `alert_notification` (Boolean) Indicates whether an alert notification is sent if the test fails.
- `comment` (String) An optional comment for the element. This field is not required.
- `down_ratio` (Number) The ratio used to determine if an aggregated link in Load Balancing mode is down. It is ignored for non-aggregated links or links not configured in Load Balancing mode. Accepted values are: 0 (not set), or any of the values in 15, 20, 25, 30, 35, 40, 45, 50.
- `name` (String) Name of the object.
- `nicid` (String) The Interface ID on which the test is run. It can be 'all', 'all_with_cvi', or a specific interface ID. Note! (Firewalls only) Only the first interface that belongs to an Aggregated Link is shown in the list of interfaces. However, the Link Status test checks the status of both interfaces in the Aggregated Link.
- `offline_state` (Boolean) Indicates whether the test is executed when the engine is offline or not.
- `online_state` (Boolean) Indicates whether the test is executed when the engine is online or not.
- `standby_state` (Boolean) Indicates whether the test is executed when the engine is in standby state or not.
- `test_action` (String) The action to be taken if the test fails. Options include 'none', 'offline', 'forceoffline', and 'forcespeed'.
- `test_active` (Boolean) Indicates whether the test is active or not.
- `test_interval` (Number) The interval in seconds at which the test is executed. Note! Running a test too frequently can increase overhead.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
