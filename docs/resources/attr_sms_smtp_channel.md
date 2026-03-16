---
page_title: "sms_smtp_channel"
subcategory: ""
description: |-
  This represents the SMS SMTP Channel used by the Mgt Server, which defines the SMTP server, SMS gateway, subject, body, and other SMTP parameters for sending SMS messages.
---

# sms_smtp_channel (Nested-Attribute)

This represents the SMS SMTP Channel used by the Mgt Server, which defines the SMTP server, SMS gateway, subject, body, and other SMTP parameters for sending SMS messages.




## Simple Attributes
- `account` (String) The SMS account associated with the Mgt Server channel, used for authentication.
- `body` (String) The body of the SMS message sent through the SMS SMTP Channel, which contains the actual content of the SMS message.
- `close_socket` (Boolean) Indicates whether the SMTP socket should be closed after sending the SMS message. If true, the socket will be closed; if false, it will remain open for further communication.
- `conn_timeout` (Number) The connection timeout for the SMS messaging channel, in milliseconds. This defines how long the system will wait for a connection to be established before timing out.
- `debug` (Boolean) Indicates whether debugging is enabled for the SMS messaging channel. If true, additional debug information will be logged.
- `from` (String) The 'From' address used in the SMTP message, which is typically the sender's email address or identifier.
- `local_hostname` (String) The local hostname used by the SMTP server, which is typically the hostname of the server where the Mgt Server is running.
- `name` (String) Name of the object.
- `password` (String) The SMS password associated with the Mgt Server channel, used for authentication.
- `rank` (Number) The rank of the SMS messaging channel, used to determine the priority of the channel.
- `sms_gateway` (String) The SMS gateway used by the SMS SMTP Channel, which defines how SMS messages are sent.
- `smtp_server_ref` (String) This represents a Simple Mail Transfer Protocol (SMTP) server, which is used to process notifications by e-mails. It includes attributes for port, default sender email, and default sender name.
- `start_tls` (Boolean) Indicates whether STARTTLS should be used for the SMTP connection. If true, the connection will be upgraded to a secure TLS connection.
- `subject` (String) The subject of the SMS message sent through the SMS SMTP Channel, which is typically used for email-based SMS gateways.
- `timeout` (Number) The timeout for the SMS messaging channel, in milliseconds. This defines how long the system will wait for a response before timing out.


