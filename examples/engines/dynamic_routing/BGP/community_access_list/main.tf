variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_community_access_list" "community_access_list" {
  comment = var.resource_comment
  entries {
    community_access_list_entry {
      action    = "deny"
      community = "178"
    }
  }
  entries {
    community_access_list_entry {
      action    = "permit"
      community = "1022"
    }
  }
  name = "tf_community_access_list"
  type = "standard"
}
