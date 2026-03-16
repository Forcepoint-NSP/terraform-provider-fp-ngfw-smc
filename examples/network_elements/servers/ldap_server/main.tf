variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ldap_server" "ldap" {
  address = "myldap.example.com"
  name = "tf_ldap"
  comment = var.resource_comment
  base_dn = "dc=example,dc=com"
  bind_password = "MySecretPassword123!"
  bind_user_id = "cn=admin,dc=example,dc=com"
  client_cert_based_user_search = ""
  display_name_attr_name = "displayName"
  email = "mail"
  frame_ip_attr_name = "msRADIUSFramedIPAddress"
  group_member_attr = "member"
  group_object_class = ["sggroup", "groupOfNames", "country", "organization", "organizationalUnit"]
  # group_object_id_attr_name = ""
  job_title_attr_name = "title"
  max_search_result = 0
  mobile_attr_name = "mobile"
  office_location_attr_name = "physicalDeliveryOfficeName"
  page_size = 0
  photo_attr_name = "photo"
  port = 389
  protocol = "ldap"
  secondary = []
  # short_user_id_attr = "cn"
  supported_method = []
  timeout = 10
  user_object_class = ["sguser", "person", "organizationalPerson", "inetOrgPerson"]
}
