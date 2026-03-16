variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ipv6_prefix_list" "ipv6_prefix_list" {
  comment = var.resource_comment
  entries {
    ipv6_prefix_list_entry {
      action            = "deny"
      max_prefix_length = 78
      min_prefix_length = 65
      subnet            = "2001:4860:4860::/64"
    }
  }
  name = "tf_ipv6_prefix_list"
}
