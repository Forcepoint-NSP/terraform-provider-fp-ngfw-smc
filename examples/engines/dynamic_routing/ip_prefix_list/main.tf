variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ip_prefix_list" "ip_prefixl_list" {
  comment = var.resource_comment
  entries {
    ip_prefix_list_entry {
      action            = "permit"
      max_prefix_length = 28
      min_prefix_length = 25
      subnet            = "192.168.1.0/24"
    }
  }
  entries {
    ip_prefix_list_entry {
      action            = "deny"
      max_prefix_length = 26
      min_prefix_length = 25
      subnet            = "10.100.100.0/24"
    }
  }
  name = "tf_ip_prefixl_list"
}
