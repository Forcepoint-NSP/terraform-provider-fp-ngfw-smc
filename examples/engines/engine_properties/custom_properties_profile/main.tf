variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

resource "smc_custom_properties_profile" "tf_custom_properties_profile" {
  custom_property {
    data_type = "string"
    name      = "text"
    value     = "Hello World"
  }
  name    = "tf_custom_properties_profile"
  type    = "custom"
  comment = var.resource_comment

  lifecycle {
    action_trigger {
      events  = [after_create]
      actions = [action.smc_custom_properties_profile_upload_script.upload_custom_script.id]
    }
  }
}

resource "smc_single_fw" "tf_single_fw" {
  name                      = "tf_single_fw"
  log_server_ref            = data.smc_href.log_server.href
  comment                   = var.resource_comment
  custom_properties_profile = [smc_custom_properties_profile.tf_custom_properties_profile.id]

  nodes {
    firewall_node {
      name   = "node1"
      nodeid = 1
    }
  }

  physical_interfaces {
    physical_interface {
      interface_id = 0
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 0
          address       = "192.168.100.1"
          network_value = "192.168.100.0/24"
          primary_mgt   = true
        }
      }
    }
  }
}

# Path to the custom properties profile content to upload.
# The SMC API expects a ZIP archive for the 'custom_script' content.
variable "script_path" {
  type    = string
  default = "./print_text.py"
}

# Upload (import) the custom properties profile content.
# Runs automatically after_create (see lifecycle above). To re-upload at any
# time, run:
#   terraform apply -invoke action.smc_custom_properties_profile_upload_script.upload_custom_script -auto-approve
action "smc_custom_properties_profile_upload_script" "upload_custom_script" {
  config {
    custom_properties_profile_href = smc_custom_properties_profile.tf_custom_properties_profile.link.custom_script
    script_path                    = var.script_path
  }
}
