variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_as_path_access_list" "as_path_access_list" {
  comment = var.resource_comment
  entries {
    as_path_access_list_entry {
      action     = "permit"
      expression = "^65001_"
    }
  }
  entries {
    as_path_access_list_entry {
      action     = "deny"
      expression = "_65002$"
    }
  }
  name = "tf_as_path_access_list"
}
