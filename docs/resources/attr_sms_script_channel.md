---
page_title: "sms_script_channel"
subcategory: ""
description: |-
  This represents the SMS Script Channel used by the Mgt Server, which defines the script path, execution path, and timeout settings for executing SMS scripts.
---

# sms_script_channel (Nested-Attribute)

This represents the SMS Script Channel used by the Mgt Server, which defines the script path, execution path, and timeout settings for executing SMS scripts.




## Simple Attributes
- `account` (String) The SMS account associated with the Mgt Server channel, used for authentication.
- `conn_timeout` (Number) The connection timeout for the SMS messaging channel, in milliseconds. This defines how long the system will wait for a connection to be established before timing out.
- `debug` (Boolean) Indicates whether debugging is enabled for the SMS messaging channel. If true, additional debug information will be logged.
- `name` (String) Name of the object.
- `password` (String) The SMS password associated with the Mgt Server channel, used for authentication.
- `proxy_port` (Number) The port of the HTTP proxy used by the SMS Script Channel, if applicable. This is used to route HTTP requests through a proxy server.
- `rank` (Number) The rank of the SMS messaging channel, used to determine the priority of the channel.
- `script_execution_path` (String) The execution path for the SMS script, which defines where the script will be executed on the server.
- `script_path` (String) The path to the SMS script that will be executed by the Mgt Server.
- `timeout` (Number) The timeout for the SMS messaging channel, in milliseconds. This defines how long the system will wait for a response before timing out.


