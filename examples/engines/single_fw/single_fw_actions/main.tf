variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "mgt_server" {
  type = "mgt_server"
  name = "*"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

resource "smc_fw_policy" "tf_aaa_policy" {
  name     = "tf_aaa"
  template = data.smc_href.default_policy_template.id
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "allow_ssh" {
  from_ref = smc_fw_policy.tf_aaa_policy.link.fw_ipv4_access_rules
  name     = "aaa"
  sources { any = true }
  destinations { any = true }
  services { any = true }
  action { action = ["allow"] }
}


resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw"
  log_server_ref = data.smc_href.log_server.href
  comment        = var.resource_comment
  nodes {
    firewall_node {
      name   = "myfwnode"
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
          address       = "192.168.100.14"
          network_value = "192.168.100.00/24"
          primary_mgt   = true
        }
      }
    }
  }

  lifecycle {
    action_trigger {
      events = [after_create]
      actions = [
        action.smc_command.bind_license.id,
        action.smc_initial_contact.tf_single_fw_contact.id,
        action.smc_command.upload.id
      ]
    }
  }
}

resource "smc_group" "tf_group1" {
  name = "tf_group1"
  element = [
    smc_single_fw.tf_single_fw.id
  ]
  is_monitored = true
  comment      = var.resource_comment
}


# note: if you need to do again initial contact:
# terraform apply -invoke action.smc_initial_contact.tf_single_fw_contact -auto-approve
action "smc_initial_contact" "tf_single_fw_contact" {
  config {
    initial_contact_href = smc_single_fw.tf_single_fw.nodes[0].firewall_node.link.initial_contact
    enable_ssh           = true
    time_zone            = "CET"
    keyboard             = "us"
    root_password        = "root"
    install_on_server    = true
    export_to_base64     = false
    output_file          = "/tmp/tf_single_fw_initial_contact.txt"
  }
}

action "smc_command" "bind_license" {
  config {
    command_href = "${var.url}/${var.api_version}/system/license_bind"
    parameters = {
      component_href = smc_single_fw.tf_single_fw.nodes[0].firewall_node.link.self
    }
  }
}

# to upload the policy after initial contact, run
# terraform apply -invoke action.smc_command.upload -auto-approve
#
# to refresh the installed policy, run
# terraform apply -invoke action.smc_command.upload -auto-approve -var refresh=true
variable "refresh" {
  type    = bool
  default = false
}
action "smc_command" "upload" {
  config {
    command_href = smc_single_fw.tf_single_fw.link[var.refresh ? "refresh" : "upload"]
    query = {
      filter = smc_fw_policy.tf_aaa_policy.name
    }
  }
}

# terraform apply -invoke action.smc_command.sginfo -var sginfo_file="./sginfo_output.zip" -auto-approve
variable "sginfo_file" {
  type    = string
  default = "/tmp/tf_single_fw_sginfo.zip"
}
action "smc_command" "sginfo" {
  config {
    command_href = smc_single_fw.tf_single_fw.nodes[0].firewall_node.link.sginfo
    method       = "get"
    # query = {
    #   include_core_files = "false"
    #   include_slapcat_output = "false"
    # }
    output_file = var.sginfo_file
  }
}


# example of change_state with parameters from variable
# terraform apply -invoke action.smc_command.change_state  -var new_state=go_online -auto-approve
variable "new_state" {
  type    = string
  default = "go_offline"

  validation {
    condition     = contains(["go_offline", "go_online", "go_standby", "power_off", "reboot"], var.new_state)
    error_message = "invalid change_state"
  }
}
action "smc_command" "change_state" {
  config {
    command_href = smc_single_fw.tf_single_fw.nodes[0].firewall_node.link[var.new_state]
    method       = "put"
  }
}

resource "smc_backup_task" "tf_backup_task" {
  name = "tf_backup_task_test"
  resources = [
    data.smc_href.mgt_server.id
  ]
  backup_comment = "backup_created_by_terraform"

  depends_on = [smc_single_fw.tf_single_fw]
}

# to create extra backup at any time, use:
# terraform apply -invoke action.smc_command.backup_now -auto-approve
action "smc_command" "backup_now" {
  config {
    command_href = smc_backup_task.tf_backup_task.link.start
  }
}

# restart webswing server
# terraform apply -invoke action.smc_command.restart_web_access -auto-approve
action "smc_command" "restart_web_access" {
  config {
    command_href = "${data.smc_href.mgt_server.id}/restart_web_access"
    method       = "put"
  }
}

data "smc_href" "shared_domain" {
  type = "admin_domain"
  name = "Shared Domain"
}

# refresh alert escalation policy of the Shared Domain
# terraform apply -invoke action.smc_command.refresh_all -auto-approve
action "smc_command" "refresh_admin_domain_policy" {
  config {
    command_href = "${data.smc_href.shared_domain.id}/refresh"
  }
}

