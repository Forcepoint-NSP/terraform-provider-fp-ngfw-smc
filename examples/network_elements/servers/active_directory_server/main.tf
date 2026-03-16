variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "user_password_auth_method" {
  name = "User password"
  type = "authentication_service"
}

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
