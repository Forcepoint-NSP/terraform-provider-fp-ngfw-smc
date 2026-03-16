variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ipv6_access_list" "ipv6_access_element" {
  comment = var.resource_comment
  entries {
    ipv6_access_list_entry {
      action = "deny"
      # need short form of ipv6
      subnet = "2001:db8:85a3::8a2e:370:7334/128"
    }
  }
  entries {
    ipv6_access_list_entry {
      action = "permit"
      subnet = "2606:2800:220:1:248:1893:25c8:1946/64"
    }
  }
  name = "ipv6_access_element"
}
