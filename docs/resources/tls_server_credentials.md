---
page_title: "smc_tls_server_credentials"
subcategory: "admin"
description: |-
  This represents a TLS Server Credentials element, which is used to store the private key and certificate of an internal server. The certificate and the associated private key must be compatible with OpenSSL and be in PEM format. It is used for TLS inspection, securing Web Access Servers, and authenticating Authentication Servers.
---

# smc_tls_server_credentials (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a TLS Server Credentials element, which is used to store the private key and certificate of an internal server. The certificate and the associated private key must be compatible with OpenSSL and be in PEM format. It is used for TLS inspection, securing Web Access Servers, and authenticating Authentication Servers.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `certificate_state` (String) The state of the certificate, which can be 'initial', 'request', or 'certificate'. This indicates whether the element is a new TLS Server Credential, a request, or a certificate.
- `comment` (String) An optional comment for the element. This field is not required.
- `key_length` (Number) The key length for the certificate, specified in bits. This is required for creating a certificate request.
- `name` (String) Name of the object.
- `public_key_algorithm` (String) The public key algorithm used for the certificate, which can be 'dsa', 'rsa', or 'ecdsa'. This is required for creating a certificate request.
- `signature_algorithm` (String) The signature algorithm used for the certificate, which can be one of the following: 'dsa_sha_1', 'dsa_sha_224', 'dsa_rsa_256', 'rsa_md5', 'rsa_sha_1', 'rsa_sha_256', 'rsa_sha_384', 'rsa_sha_512', 'ecdsa_sha_1', 'ecdsa_sha_256', 'ecdsa_sha_384', 'ecdsa_sha_512'. This is required for creating a certificate request.
- `subject_alt_names` (String) A formatted representation of subject alternative names (SANs) for the certificate. This is used to specify additional domain names or IP addresses that the certificate should be valid for.
- `subject_name` (String) The Distinguished Name (DN) of the certificate subject, following the syntax: CN=CommonName,O=Organization,OU=Unit,C=Country,ST=State,L=City. This is required for creating a certificate request.
- `valid_from` (Number) The possible valid from date for the associated certificate, specified in milliseconds since epoch. This indicates when the certificate becomes valid.
- `valid_to` (Number) The possible valid to date for the associated certificate, specified in milliseconds since epoch. This indicates when the certificate expires.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
