variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "any_network" {
  type = "network"
  name = "Any network"
}

resource "smc_network" "tf_sample_network_single_1" {
  name         = "tf_sample_net-212.20.1.0/24"
  comment      = var.resource_comment
  ipv4_network = "212.20.1.0/24"
}

resource "smc_network" "tf_sample_network_single_2" {
  name         = "tf_sample_net-212.20.30.0/24"
  comment      = var.resource_comment
  ipv4_network = "212.20.30.0/24"
}

resource "smc_router" "tf_sample_router_1" {
  address = "212.20.1.254"
  name    = "tf_sample_router-212.20.1.254"
  comment = var.resource_comment
}

resource "smc_router" "tf_sample_router_2" {
  address = "212.20.30.254"
  name    = "tf_sample_router-212.20.30.254"
  comment = var.resource_comment
}

data "smc_href" "suite_b_gcm_128" {
  name = "Suite-B-GCM-128"
  type = "vpn_profile"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_sub_href" "internal_gateway_1_ref" {
  from_ref = smc_single_fw.tf_single_fw1.link.internal_gateway
  match    = "tf_single_fw1*"
}

data "smc_sub_href" "internal_gateway_2_ref" {
  from_ref = smc_single_fw.tf_single_fw2.link.internal_gateway
  match    = "tf_single_fw2*"
}

resource "smc_single_fw" "tf_single_fw1" {
  depends_on = [
    smc_network.tf_sample_network_single_1,
    smc_router.tf_sample_router_1
  ]
  name           = "tf_single_fw1"
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
          address       = "192.168.1.1"
          network_value = "192.168.1.0/24"
        }
      }
    }
  }
  physical_interfaces {
    physical_interface {
      interface_id = 1
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 1
          address       = "212.20.1.1"
          network_value = "212.20.1.0/24"
        }
      }
    }
  }
  physical_interfaces {
    physical_interface {
      interface_id = 2
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 2
          address       = "172.16.18.10"
          network_value = "172.16.18.0/24"
          primary_mgt   = true
        }
      }
    }
  }
}

resource "smc_single_fw" "tf_single_fw2" {
  depends_on = [
    smc_network.tf_sample_network_single_2,
    smc_router.tf_sample_router_2
  ]
  name           = "tf_single_fw2"
  log_server_ref = data.smc_href.log_server.href
  comment        = var.resource_comment
  nodes {
    firewall_node {
      name   = "myfwnode2"
      nodeid = 1
    }
  }

  physical_interfaces {
    physical_interface {
      interface_id = 1
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 1
          address       = "192.168.2.1"
          network_value = "192.168.2.0/24"
        }
      }
    }
  }
  physical_interfaces {
    physical_interface {
      interface_id = 0
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 0
          address       = "212.20.30.1"
          network_value = "212.20.30.0/24"
          primary_mgt   = true
        }
      }
    }
  }
}

resource "smc_internal_endpoint" "ep_212_20_30_1" {
  id = format("%s#%s", smc_single_fw.tf_single_fw2.link.internal_gateway,
  "*/internal_endpoint/212.20.30.1")
  enabled   = true
  ipsec_vpn = true
}

resource "smc_internal_endpoint" "ep_212_20_1_1" {
  id = format("%s#%s", smc_single_fw.tf_single_fw1.link.internal_gateway,
  "*/internal_endpoint/212.20.1.1")
  enabled   = true
  ipsec_vpn = true
}

resource "smc_vpn" "tf_sample_vpn" {
  mobile_vpn_topology_mode = "None"
  name                     = "tf_sample_vpn"
  nat                      = false
  vpn_profile              = data.smc_href.suite_b_gcm_128.href
  comment                  = var.resource_comment
}

resource "smc_gateway_node" "tf_sample_gateway_1_node" {
  from_ref   = smc_vpn.tf_sample_vpn.link.central_gateway_node
  name       = "tf_sample_gateway_node_1"
  node_usage = "central"
  gateway    = data.smc_sub_href.internal_gateway_1_ref.id
}

resource "smc_gateway_node" "tf_sample_gateway_2_node" {
  from_ref   = smc_vpn.tf_sample_vpn.link.central_gateway_node
  name       = "tf_sample_gateway_node_2"
  node_usage = "central"
  gateway    = data.smc_sub_href.internal_gateway_2_ref.id
}

# resource "smc_gateway_tunnel" "tf_sample_tunnel" {
#   depends_on = [smc_internal_endpoint.ep_212_20_1_1, smc_internal_endpoint.ep_212_20_1_1]
#   # todo: problem: the generated name is either
#   # - Gateway Tunnel tf_single_fw1 - Primary-tf_single_fw2 - Primary
#   # - Gateway Tunnel tf_single_fw2 - Primary-tf_single_fw1 - Primary
#   id      = "${smc_vpn.tf_sample_vpn.link.gateway_tunnel}#Gateway Tunnel*"
#   name    = "tf_sample_tunnel"
#   enabled = false
#   preshared_key = "TheSuperPreShareKey00!"
# }
