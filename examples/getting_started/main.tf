/*
 Terraform configuration to create a single firewall with AAA policy,
 network, router, and initial contact action.
 */
variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  tf_single_fw_name    = "tf_sample_fw"
  tf_single_fw_address = "192.168.1.1"
  tf_router            = "192.168.1.254"
  tf_network           = "192.168.1.0/24"
  tf_dns               = "8.8.8.8"
  tf_fw_ref            = smc_single_fw.tf_sample_fw
  tf_node_ref          = local.tf_fw_ref.nodes[0].firewall_node
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "any_network" {
  type = "network"
  name = "Any network"
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

# Allow all rule for testing purposes
resource "smc_fw_ipv4_access_rule" "allow_ssh" {
  from_ref = smc_fw_policy.tf_aaa_policy.link.fw_ipv4_access_rules
  name     = "aaa"
  sources      { any = true }
  destinations { any = true }
  services     { any = true }
  action       { action = ["allow"] }
}

resource "smc_network" "tf_sample_network" {
  name         = "tf_sample_network"
  comment      = var.resource_comment
  ipv4_network = local.tf_network

}

resource "smc_router" "tf_sample_router" {
  address = local.tf_router
  name    = "tf_sample_router"
  comment = var.resource_comment
}

module "tf_sample_fw_node1" {
  source = "./modules/simple_firewall_node"
  name   = "${local.tf_single_fw_name}_node1"
  nodeid = 1
}

module "intf0" {
  source       = "./modules/simple_physical_interface"
  nodeid       = module.tf_sample_fw_node1.nodeid
  interface_id = 0
  nicid        = 0
  address      = local.tf_single_fw_address
  network      = smc_network.tf_sample_network.ipv4_network
  mgmt         = true
}

module "route_intf0" {
  source      = "./modules/simple_route"
  interface   = "Interface 0"
  network     = smc_network.tf_sample_network.id
  gateway     = smc_router.tf_sample_router.id
  destination = data.smc_href.any_network.id
}

resource "smc_single_fw" "tf_sample_fw" {
  name                  = local.tf_single_fw_name
  comment               = var.resource_comment
  log_server_ref        = data.smc_href.log_server.href
  nodes                 = [module.tf_sample_fw_node1.value]
  physical_interfaces   = [module.intf0.value]
  domain_server_address = [{ value = local.tf_dns, rank = 1.0 }]

  depends_on = [
    smc_router.tf_sample_router,
    smc_network.tf_sample_network,
    smc_fw_policy.tf_aaa_policy
  ]

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

resource "smc_routing_node" "tf_single_fw" {
  id           = smc_single_fw.tf_sample_fw.link.routing
  level        = "engine_cluster"
  routing_node = [module.route_intf0.value]
}

action "smc_initial_contact" "tf_single_fw_contact" {
  config {
    initial_contact_href = local.tf_node_ref.link.initial_contact
    enable_ssh           = true
    time_zone            = "CET"
    keyboard             = "us"
    root_password        = "root"
    install_on_server    = true
    export_to_base64     = false
    output_file          = "/tmp/${local.tf_single_fw_name}_initial_contact.txt"
  }
}

action "smc_command" "bind_license" {
  config {
    command_href = "${var.url}/${var.api_version}/system/license_bind"
    parameters = {
      component_href = local.tf_node_ref.link.self
    }
  }
}

action "smc_command" "upload" {
  config {
    command_href = local.tf_fw_ref.link.upload
    query = {
      filter = smc_fw_policy.tf_aaa_policy.name
    }
  }
}
