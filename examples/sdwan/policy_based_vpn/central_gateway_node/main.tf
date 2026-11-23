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
  mobile_vpn_topology_mode = "none"
  name                     = "tf_sample_vpn"
  nat                      = false
  vpn_profile              = data.smc_href.suite_b_gcm_128.href
  comment                  = var.resource_comment
}

resource "smc_gateway_node" "tf_sample_gateway_1_node" {
  from_ref   = smc_vpn.tf_sample_vpn.link.central_gateway_node
  comment    = "tf_sample_gateway_node_1"
  node_usage = "central"
  gateway    = data.smc_sub_href.internal_gateway_1_ref.id
}

resource "smc_gateway_node" "tf_sample_gateway_2_node" {
  from_ref   = smc_vpn.tf_sample_vpn.link.central_gateway_node
  comment    = "tf_sample_gateway_node_2"
  node_usage = "central"
  gateway    = data.smc_sub_href.internal_gateway_2_ref.id
}

# Pre-shared key on the tunnel between the two gateways.
#
# The SMC generates the logical tunnel itself once both gateways are in the
# VPN topology; smc_gateway_tunnel adopts it by name and sets its key. The
# generated name is "Gateway Tunnel <gw A> - Primary-<gw B> - Primary" and the
# order of A and B is not stable, so match on one gateway with a wildcard
# instead of spelling out both names. If nothing matches, the error is a
# misleading "jmespath search failed: SyntaxError".
#
# The VPN profile referenced by smc_vpn must have pre_shared_key_for_ike = true;
# the built-in Suite-B-GCM-128 used here does. The key is kept in the state:
# use a local state or a protected backend.
variable "preshared_key" {
  type      = string
  sensitive = true
  # Example value only: override it (TF_VAR_preshared_key). The engine reads a
  # key starting with 0x as hexadecimal and 0s as Base64, so a random text key
  # must not start with either (see README.md).
  default = "TheSuperPreShareKey00!"
}

resource "smc_gateway_tunnel" "tf_sample_tunnel" {
  depends_on = [
    smc_gateway_node.tf_sample_gateway_1_node,
    smc_gateway_node.tf_sample_gateway_2_node,
    smc_internal_endpoint.ep_212_20_1_1,
    smc_internal_endpoint.ep_212_20_30_1,
  ]
  id            = "${smc_vpn.tf_sample_vpn.link.gateway_tunnel}#Gateway Tunnel *tf_single_fw2 - *"
  enabled       = true
  preshared_key = var.preshared_key
}
