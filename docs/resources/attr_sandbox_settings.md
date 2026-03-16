---
page_title: "sandbox_settings"
subcategory: ""
description: |-
  This represents the definition of Sandbox settings feature, which includes configuration for sandbox service, license key, license token, and HTTP proxies.
---

# sandbox_settings (Nested-Attribute)

This represents the definition of Sandbox settings feature, which includes configuration for sandbox service, license key, license token, and HTTP proxies.




## Simple Attributes
- `http_proxy` (List of String) URI of the HTTP Proxy.
- `sandbox_license_key` (String) The license key for the sandbox service, which is required for activation and usage of the sandbox features.
- `sandbox_license_token` (String) The license token for the sandbox service, which is used for authentication and authorization of sandbox features.
- `sandbox_service` (String) This represents the Sandbox service element (referenced by sandbox settings). It contains information about the portal account, data center, license key and token.


