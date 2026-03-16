---
page_title: "smc_user_response_entry"
subcategory: ""
description: |-
  This represents an entry in a User Response, which includes the reason for the response, type of response, and additional details such as redirect URL, text, title, and message.
---

# smc_user_response_entry (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an entry in a User Response, which includes the reason for the response, type of response, and additional details such as redirect URL, text, title, and message.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `reason` (String) The reason for sending this response, which can be used to provide context or explanation for the response.
- `redirect` (String) The type of redirect for the response, which can be 'automatic' or 'manual'.
- `type` (String) The type of response being sent, which can be one of the predefined response types such as 'tcp_close', 'url_redirect', or 'html_page'.
- `user_response_message` (String) The message content of the response, which is applicable for HTML advanced responses.
- `user_response_text` (String) The text content of the response, which can be a URL for URL redirects or HTML content for HTML responses.
- `user_response_title` (String) The title of the response message, which is applicable for HTML advanced responses.


