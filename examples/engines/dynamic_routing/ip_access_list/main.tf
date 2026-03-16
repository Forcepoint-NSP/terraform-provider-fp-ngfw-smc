variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ip_access_list" "ip_access_list" {
  name = "tf_ip_access_list"
  comment = var.resource_comment

  entries {
    ip_access_list_entry {
      action = "deny"
      subnet = "172.16.16.0/21"
    }
  }
  entries {
    ip_access_list_entry {
      action = "permit"
      subnet = "192.168.100.0/24"
    }
  }
}
