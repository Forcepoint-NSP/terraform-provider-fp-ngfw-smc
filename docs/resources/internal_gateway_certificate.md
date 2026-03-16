---
page_title: "smc_gateway_certificate"
subcategory: "vpn"
description: |-
  This represents a certificate for an internal gateway, which includes details about the certificate authority, validity dates, and algorithms used.
---

# smc_gateway_certificate (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a certificate for an internal gateway, which includes details about the certificate authority, validity dates, and algorithms used.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `certificate_authority` (String) This represents a VPN Certificate Authority, which is used to manage VPN certificate authorities in the system.
- `certificate_base64` (String) The PEM-encoded certificate for the internal gateway, which contains the public key and other metadata.
- `comment` (String) An optional comment for the element. This field is not required.
- `expiration_date` (String) The date when the certificate expires, indicating when the certificate is no longer valid.
- `name` (String) Name of the object.
- `public_key_algorithm` (String) The public key algorithm used in the certificate, which defines the cryptographic method for the public key.
- `signature_algorithm` (String) The signature algorithm used in the certificate, which defines how the certificate is signed and verified.
- `subject_alt_name` (List of String) The subject alternative names (SAN) of the certificate, which provide additional identities for the certificate beyond the common name.
- `valid_from` (String) The date from which the certificate is valid, indicating when the certificate starts being effective.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
