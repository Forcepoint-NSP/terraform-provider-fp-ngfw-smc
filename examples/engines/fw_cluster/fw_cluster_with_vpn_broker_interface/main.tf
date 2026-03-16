variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  engine_name = "tf_fw_cluster_with_vpn_broker_interface"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "Log Server"
}

data "smc_sub_href" "internal_gateway_ref" {
  from_ref = smc_fw_cluster.tf_cluster_example.link.internal_gateway
  match    = "tf_fw_cluster*"
}

resource "smc_broker_config_file" "vpn_broker_config" {
  name           = "tf_vpn_broker_config"
  binary_content = "aGVsbG8K"
  comment        = var.resource_comment
}

resource "smc_vpn_broker_domain" "vpn_broker_domain" {
  name               = "tf_vpn_broker_domain"
  file_ref           = smc_broker_config_file.vpn_broker_config.id
  mac_address_prefix = "00:00:20"
  comment            = var.resource_comment
}

resource "smc_fw_cluster" "tf_cluster_example" {
  name           = local.engine_name
  cluster_mode   = "balancing"
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href
  nodes {
    firewall_node {
      name   = "${local.engine_name} node 1"
      nodeid = 1
    }
  }
  nodes {
    firewall_node {
      name   = "${local.engine_name} node 2"
      nodeid = 2
    }
  }
  physical_interfaces {
    physical_interface {
      interface_id = "0"
      name         = "Interface 0"
      cvi_mode     = "packetdispatch"
      macaddress   = "00:00:5e:00:02:00"

      interfaces {
        node_interface {
          nicid             = "0"
          nodeid            = 1
          address           = "192.168.2.21"
          network_value     = "192.168.2.0/24"
          outgoing          = true
          primary_heartbeat = true
          primary_mgt       = true
        }
      }
      interfaces {
        node_interface {
          nicid             = "0"
          nodeid            = 2
          address           = "192.168.2.22"
          network_value     = "192.168.2.0/24"
          outgoing          = true
          primary_heartbeat = true
          primary_mgt       = true

        }
      }
      interfaces {
        cluster_virtual_interface {
          nicid           = "0"
          address         = "192.168.2.254"
          network_value   = "192.168.2.0/24"
          relayed_by_dhcp = false

        }
      }
    }
  }
}


resource "smc_vpn_broker_interface" "vpn_broker_intfx" {
  from_ref            = smc_fw_cluster.tf_cluster_example.link.vpn_broker_interface
  adjust_antispoofing = true
  gateway_ref         = data.smc_sub_href.internal_gateway_ref.id
  interface_id        = "VPN_12"
  interfaces {
    node_interface {
      address       = "12.12.12.1"
      network_value = "12.12.12.0/24"
      nicid         = "VPN_12"
      nodeid        = 1
    }
  }
  interfaces {
    node_interface {
      address       = "12.12.12.2"
      network_value = "12.12.12.0/24"
      nicid         = "VPN_12"
      nodeid        = 2
    }
  }
  interfaces {
    cluster_virtual_interface {
      address         = "12.12.12.12"
      network_value   = "12.12.12.0/24"
      nicid           = "VPN_12"
      relayed_by_dhcp = false
    }
  }
  mac_address_postfix   = "00:11:11"
  name                  = "VPN Broker Interface 12"
  shared_secret         = "SecretPassword00!"
  vpn_broker_domain_ref = smc_vpn_broker_domain.vpn_broker_domain.id
}
