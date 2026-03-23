---
page_title: "smc_active_directory_server"
subcategory: "network_elements"
description: |-
  This represents an Active Directory Server, which is used to store user information in a user database and can be queried during the user authentication process. It includes attributes for authentication IP, domain controllers, IAS service, retries, shared secret, and authentication port.
---

# smc_active_directory_server (Resource)

This represents an Active Directory Server, which is used to store user information in a user database and can be queried during the user authentication process. It includes attributes for authentication IP, domain controllers, IAS service, retries, shared secret, and authentication port.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.741.0/examples/network_elements/servers/active_directory_server) for a complete minimal example

This example creates an Active Directory server object in SMC with typical LDAP configuration attributes and associates it with an authentication service.

```hcl
resource "smc_active_directory_server" "ad" {
  address                       = "myad.example.com"
  comment                       = var.resource_comment
  auth_port                     = 1812
  base_dn                       = "dc=example,dc=com"
  bind_password                 = "my_bind_password00!!"
  bind_user_id                  = "cn=admin,dc=example,dc=com"
  client_cert_based_user_search = ""
  display_name_attr_name        = "displayName"
  email                         = "mail"
  frame_ip_attr_name            = "msRADIUSFramedIPAddress"
  group_member_attr             = "member"
  group_object_class            = ["sggroup", "groupOfNames", "country", "organization", "organizationalUnit", "group"]
  # group_object_id_attr_name = ""
  internet_auth_service_enabled = true
  job_title_attr_name           = "title"
  # long_user_id_attr = "userPrincipalName"
  max_search_result         = 0
  mobile_attr_name          = "mobile"
  name                      = "ad"
  office_location_attr_name = "physicalDeliveryOfficeName"
  page_size                 = 1000
  photo_attr_name           = "photo"
  port                      = 389
  protocol                  = "ldap"
  retries                   = 2
  secondary                 = []
  shared_secret             = "*****"
  # short_user_id_attr = "sAMAccountName"
  supported_method  = [data.smc_href.user_password_auth_method.id]
  timeout           = 10
  user_object_class = ["sguser", "person", "organizationalPerson", "inetOrgPerson"]
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the device, which is used for network communication.
- `auth_ipaddress` (String) The IPv4 address used for authentication on the Active Directory server. If the authentication service on the Active Directory server uses a different IP address than the server itself, enter the IPv4 address for authentication.
- `auth_port` (Number) The port number for the Active Directory server authentication.
- `base_dn` (String) The base distinguished name (DN) for the LDAP server, which defines the starting point for user searches.
- `bind_password` (String) The password for the user ID used for binding to the LDAP server.
- `bind_user_id` (String) The distinguished name (DN) of the user ID used for binding to the LDAP server.
- `client_cert_based_user_search` (String) The LDAP query filter used to find user entries based on client certificate subject distinguished name.
- `comment` (String) An optional comment for the element. This field is not required.
- `display_name_attr_name` (String) The attribute name used for the display name in the LDAP directory.
- `email` (String) The attribute name used for the user's email address in the LDAP directory.
- `frame_ip_attr_name` (String) The attribute name used for the framed IP address in the LDAP directory.
- `group_member_attr` (String) The attribute name used for group membership in the LDAP directory.
- `group_object_class` (List of String) The LDAP object classes used for group entries in the LDAP directory.
- `group_object_id_attr_name` (String) The attribute name used for the unique identifier of the LDAP group object in the LDAP directory.
- `internet_auth_service_enabled` (Boolean) Indicates whether the IAS service is enabled on the Active Directory server.
- `ipv6_address` (String) The primary IPv6 address of the device, which is used for network communication.
- `job_title_attr_name` (String) The attribute name used for the user's job title in the LDAP directory.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `long_user_id_attr` (String) The attribute name used for the user's UPN in the LDAP directory.
- `max_search_result` (Number) The maximum number of LDAP entries that can be returned in a search result.
- `mobile_attr_name` (String) The attribute name used for the user's mobile phone number in the LDAP directory.
- `name` (String) Name of the object.
- `office_location_attr_name` (String) The attribute name used for the user's office location in the LDAP directory.
- `page_size` (Number) The maximum number of LDAP entries returned on each page of the LDAP response.
- `photo_attr_name` (String) The attribute name used for the user's photo in the LDAP directory.
- `port` (Number) The port number for the LDAP server, which is used for communication.
- `protocol` (String) The protocol used for LDAP communication, such as ldap, ldaps, or ldap_tls.
- `retries` (Number) The number of times the Firewalls try to connect to the RADIUS or TACACS+ authentication server if the connection fails.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `shared_secret` (String) The shared secret defined for RADIUS clients on the Active Directory server. This is required at creation in clear text, and if not set during a modification, it keeps the current password.
- `short_user_id_attr` (String) The attribute name used for the user ID in the LDAP directory.
- `supported_method` (List of String) URI of the Authentication Method.
- `timeout` (Number) The timeout value in seconds for the LDAP server connection.
- `tls_identity` (String) This represents a TLS Identity, which contains data to check server identity when establishing a TLS connection. It includes fields such as DNS name, IP address, common name, distinguished name, and various hash values.
- `tls_profile_ref` (String) This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.
- `user_object_class` (List of String) The LDAP object classes used for user entries in the LDAP directory.

## Nested Attributes
- `domain_controller` (List of Blocks, see [here](attr_domain_controller.md)) The Domain Controllers associated with the Active Directory Server. Users are associated with IP addresses based on logs collected by the Active Directory Domain Controller. It is not recommended to use this feature for terminal servers or other computers that have many users logged on at the same time.
- `third_party_monitoring` (Single Block, see [here](attr_third_party_monitoring.md)) 

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
