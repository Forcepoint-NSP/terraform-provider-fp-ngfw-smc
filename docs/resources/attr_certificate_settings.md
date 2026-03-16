---
page_title: "certificate_settings"
subcategory: ""
description: |-
  This represents the certificate settings for an element, including the subject name, subject alt name, certificate type, and revocation check settings.
---

# certificate_settings (Nested-Attribute)

This represents the certificate settings for an element, including the subject name, subject alt name, certificate type, and revocation check settings.




## Simple Attributes
- `certificate_type` (String) The type of certificate to be used.
- `check_revocation` (Boolean) Indicates whether to check the revocation status of the certificate.
- `ignore_revocation_on_failure` (Boolean) Indicates whether to ignore revocation checks when an error occurs.
- `subject_alt_name` (String) The subject alternative name (SAN) for the certificate, typically a DNS name.
- `subject_name` (String) The distinguished name (DN) for the certificate.


