variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_sub_href" "internal_gateway_ref" {
  from_ref = smc_single_fw.tf_single_fw.link.internal_gateway
  match    = "tf_single_fw_vpn_broker_interfaces*"
}

resource "smc_broker_config_file" "vpn_broker_config" {
  name = "tf_vpn_broker_config"
}

resource "smc_vpn_broker_domain" "vpn_broker_domain" {
  name               = "tf_vpn_broker_domain"
  file_ref           = smc_broker_config_file.vpn_broker_config.id
  mac_address_prefix = "00:00:20"
  comment            = var.resource_comment
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw_vpn_broker_interfaces"
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
}

resource "smc_vpn_broker_interface" "vpn_broker_intfx" {
  from_ref = smc_single_fw.tf_single_fw.link.vpn_broker_interface

  adjust_antispoofing             = false
  gateway_ref                     = data.smc_sub_href.internal_gateway_ref.id
  include_prefix_info_option_flag = false
  interface_id                    = "VPN_0"
  interfaces {
    single_node_interface {
      address       = "10.11.12.13"
      network_value = "10.11.12.0/24"
      nicid         = "VPN_0"
      nodeid        = 1
    }
  }
  mac_address_postfix = "11:11:11"
  name                = "VPN Broker Interface 0"
  # todo shared_secret is not displayed correctly in the smc. Need
  # encoding ?
  shared_secret         = "SecretPassword123!"
  syn_mode              = "default"
  vpn_broker_domain_ref = smc_vpn_broker_domain.vpn_broker_domain.id
}
