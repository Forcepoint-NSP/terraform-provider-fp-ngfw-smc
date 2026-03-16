---
page_title: "sms_http_channel"
subcategory: ""
description: |-
  This represents the SMS HTTP Channel used by the Mgt Server, which defines the URL, proxy settings, and other HTTP parameters for sending SMS messages.
---

# sms_http_channel (Nested-Attribute)

This represents the SMS HTTP Channel used by the Mgt Server, which defines the URL, proxy settings, and other HTTP parameters for sending SMS messages.




## Simple Attributes
- `account` (String) The SMS account associated with the Mgt Server channel, used for authentication.
- `attr_name_message` (String) The HTTP attribute name for the message content in the SMS HTTP Channel request.
- `attr_name_phone` (String) The HTTP attribute name for the phone number in the SMS HTTP Channel request.
- `conn_timeout` (Number) The connection timeout for the SMS messaging channel, in milliseconds. This defines how long the system will wait for a connection to be established before timing out.
- `debug` (Boolean) Indicates whether debugging is enabled for the SMS messaging channel. If true, additional debug information will be logged.
- `default_success` (Boolean) Indicates whether the SMS HTTP Channel considers a response as successful by default, even if the response body does not match the success response body.
- `failure_response_body` (String) The HTTP response body that indicates a failure when sending an SMS message through the SMS HTTP Channel.
- `follow_redirects` (Boolean) Indicates whether the SMS HTTP Channel should follow HTTP redirects when sending messages.
- `http_parameters` (String) The HTTP parameters encoded string for the SMS HTTP Channel, which includes additional parameters to be sent in the request.
- `name` (String) Name of the object.
- `password` (String) The SMS password associated with the Mgt Server channel, used for authentication.
- `post` (Boolean) Indicates whether the SMS HTTP Channel should use HTTP POST method for sending messages.
- `proxy_host` (String) The host of the HTTP proxy used by the SMS HTTP Channel, if applicable.
- `proxy_port` (Number) The port of the HTTP proxy used by the SMS HTTP Channel, if applicable.
- `rank` (Number) The rank of the SMS messaging channel, used to determine the priority of the channel.
- `success_response_body` (String) The HTTP response body that indicates a successful message delivery when using the SMS HTTP Channel.
- `timeout` (Number) The timeout for the SMS messaging channel, in milliseconds. This defines how long the system will wait for a response before timing out.
- `url` (String) The URL of the SMS HTTP Channel, which is used to send SMS messages.
- `use_http_11` (Boolean) Indicates whether the SMS HTTP Channel should use HTTP/1.1 protocol.


