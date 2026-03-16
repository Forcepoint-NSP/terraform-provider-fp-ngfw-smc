---
page_title: "elasticsearch_security_settings"
subcategory: ""
description: |-
  This represents the Elasticsearch security settings, including authentication method, credentials, and TLS settings.
---

# elasticsearch_security_settings (Nested-Attribute)

This represents the Elasticsearch security settings, including authentication method, credentials, and TLS settings.




## Simple Attributes
- `api_key` (String) The API key for authentication to connect to the Elasticsearch cluster. This must be set in clear text when the 'api_key' authentication method is used.
- `login` (String) The login for basic authentication to connect to the Elasticsearch cluster. This must be set when the 'basic' authentication method is used.
- `method` (String) The authentication method used for connecting to the Elasticsearch cluster.
- `password` (String) The password for basic authentication to connect to the Elasticsearch cluster. This must be set in clear text when the 'basic' authentication method is used.
- `tls_credentials` (String) This represents a TLS Server Credentials element, which is used to store the private key and certificate of an internal server. The certificate and the associated private key must be compatible with OpenSSL and be in PEM format. It is used for TLS inspection, securing Web Access Servers, and authenticating Authentication Servers.


