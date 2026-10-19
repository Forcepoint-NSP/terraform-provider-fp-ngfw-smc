/*
  Generate and sign a VPN certificate for a firewall's internal gateway
  using an internal RSA certificate authority.

  A single apply performs the full lifecycle:
    1. Create the firewall.
    2. Trigger: bind license, run initial contact, upload policy.
    3. Create the VPN endpoint.
    4. Trigger: wait for the engine to come online, then generate the certificate.

  Usage:
    terraform apply -auto-approve
*/

variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  fw_name           = "tf_vpn_cert_fw"
  fw_address        = "192.168.10.1"
  fw_network        = "192.168.10.0/24"
  tf_single_fw_name = local.fw_name
  tf_fw_ref         = smc_single_fw.fw
  tf_node_ref       = local.tf_fw_ref.nodes[0].firewall_node
}

resource "smc_fw_policy" "tf_aaa_policy" {
  name     = "tf_aaa_int"
  template = data.smc_href.default_policy_template.id
  comment  = var.resource_comment
}

# Allow all rule for testing purposes
resource "smc_fw_ipv4_access_rule" "allow_ssh" {
  from_ref = smc_fw_policy.tf_aaa_policy.link.fw_ipv4_access_rules
  name     = "aaa"
  sources { any = true }
  destinations { any = true }
  services { any = true }
  action { action = ["allow"] }
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

# Look up the built-in internal RSA certificate authority.
data "smc_href" "internal_rsa_ca" {
  type = "vpn_certificate_authority"
  name = "Internal RSA CA for Gateways 20201019"
}

resource "smc_single_fw" "fw" {
  name           = local.fw_name
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href

  nodes {
    firewall_node {
      name   = "${local.fw_name}_node1"
      nodeid = 1
    }
  }

  # Interface 0 — management
  physical_interfaces {
    physical_interface {
      interface_id = 0
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 0
          address       = "192.168.1.1"
          network_value = "192.168.1.0/24"
          primary_mgt   = true
        }
      }
    }
  }

  # Interface 1 — VPN endpoint
  physical_interfaces {
    physical_interface {
      interface_id = 1
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 1
          address       = local.fw_address
          network_value = local.fw_network
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
        action.smc_engine_upload_policy.upload.id
      ]
    }
  }
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

action "smc_engine_upload_policy" "upload" {
  config {
    href        = local.tf_fw_ref.link.upload
    policy_name = smc_fw_policy.tf_aaa_policy.name
    wait        = true
  }
}

# Enable the VPN endpoint on interface 1.
# After the endpoint is created, generate the VPN certificate.
# The certificate action waits for the engine to come online before proceeding.
resource "smc_internal_endpoint" "vpn_ep" {
  id        = "${smc_single_fw.fw.link.internal_gateway}#*/internal_endpoint/${local.fw_address}"
  enabled   = true
  ipsec_vpn = true

  lifecycle {
    action_trigger {
      events  = [after_create]
      actions = [action.smc_generate_vpn_certificate.generate_and_sign.id]
    }
  }
}

# Resolve the internal gateway href (implicitly created with the firewall).
data "smc_sub_href" "internal_gateway" {
  from_ref = smc_single_fw.fw.link.internal_gateway
  match    = "${local.fw_name}*"
}

# Setting certificate_authority_href makes this a single generate-and-sign
# call: the SMC signs the certificate request with the internal RSA CA in
# the same request, so there is no separate signing step. The action
# resolves the gateway's generate_certificate link itself and automatically
# waits for the engine to be Online before proceeding.
action "smc_generate_vpn_certificate" "generate_and_sign" {
  config {
    gateway_href               = data.smc_sub_href.internal_gateway.id
    common_name                = "${local.fw_name} VPN certificate"
    organization               = "Forcepoint"
    public_key_algorithm       = "rsa"
    signature_algorithm        = "rsa_sha_512"
    public_key_length          = 2048
    certificate_authority_href = data.smc_href.internal_rsa_ca.id
  }
}
