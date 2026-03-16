variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "any_network" {
  type = "network"
  name = "Any network"
}

resource "smc_network" "tf_sample_network_single_1" {
  name         = "tf_sample_net-192.168.100.0/24"
  comment      = var.resource_comment
  ipv4_network = "192.168.100.0/24"
}

resource "smc_network" "tf_sample_network_single_2" {
  name         = "tf_sample_net-10.0.50.0/24"
  comment      = var.resource_comment
  ipv4_network = "10.0.50.0/24"
}

resource "smc_router" "tf_sample_router_1" {
  address = "192.168.100.254"
  name    = "tf_sample_router-192.168.100.254"
  comment = var.resource_comment
}

resource "smc_router" "tf_sample_router_2" {
  address = "10.0.50.254"
  name    = "tf_sample_router-10.0.50.254"
  comment = var.resource_comment
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_sub_href" "single_fw1_internal_gateway_ref" {
  from_ref = smc_single_fw.tf_single_fw1.link.internal_gateway
  match    = "tf_single_fw1*"
}

data "smc_sub_href" "single_fw1_tunnel_interface_6483_ref" {
  from_ref = smc_single_fw.tf_single_fw1.link.tunnel_interface
  match    = "*6483"
}

data "smc_sub_href" "single_fw1_tunnel_interface_1000_ref" {
  from_ref = smc_single_fw.tf_single_fw1.link.tunnel_interface
  match    = "*1000"
}

data "smc_sub_href" "single_fw2_internal_gateway_ref" {
  from_ref = smc_single_fw.tf_single_fw2.link.internal_gateway
  match    = "tf_single_fw2*"
}

data "smc_sub_href" "single_fw2_tunnel_interface_4563_ref" {
  from_ref = smc_single_fw.tf_single_fw2.link.tunnel_interface
  match    = "*4563"
}

data "smc_sub_href" "single_fw2_tunnel_interface_1234_ref" {
  from_ref = smc_single_fw.tf_single_fw2.link.tunnel_interface
  match    = "*1234"
}

data "smc_sub_href" "single_fw1_ep" {
  depends_on = [smc_internal_endpoint.ep_192_168_100_14]
  from_ref   = smc_single_fw.tf_single_fw1.link.internal_gateway
  match      = "*/internal_endpoint/192.168.100.14"
}

data "smc_sub_href" "single_fw2_ep" {
  depends_on = [smc_internal_endpoint.ep_10_0_50_14]
  from_ref   = smc_single_fw.tf_single_fw2.link.internal_gateway
  match      = "*/internal_endpoint/10.0.50.14"
}

resource "smc_single_fw" "tf_single_fw1" {
  name           = "tf_single_fw1"
  log_server_ref = data.smc_href.log_server.href
  comment        = var.resource_comment
  nodes {
    firewall_node {
      name   = "tf_single_fw1 node 1"
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
  physical_interfaces {
    tunnel_interface {
      interface_id = "6483"
      interfaces {
        single_node_interface {
          address       = "959b:4bb4:a752:a0ef:991b:45e4:a165:e0ee"
          network_value = "959b:4bb4:a752:a0ef:991b:45e4:a165:e0ee/128"
          nicid         = "6483"
          nodeid        = 1
        }
      }
      name = "Tunnel Interface 6483"
    }
  }
  physical_interfaces {
    tunnel_interface {
      interface_id = "1000"
      interfaces {
        single_node_interface {
          address       = "192.168.123.1"
          network_value = "192.168.123.1/32"
          nicid         = "1000"
          nodeid        = 1
        }
      }
      name = "Tunnel Interface 1000"
    }
  }
  depends_on = [smc_network.tf_sample_network_single_1, smc_router.tf_sample_router_1]

}

resource "smc_routing_node" "tf_single_fw" {
  id    = smc_single_fw.tf_single_fw1.link.routing
  level = "engine_cluster"
  routing_node {
    level = "interface"
    name  = "Interface 0"
    routing_node {
      level = "network"
      href  = smc_network.tf_sample_network_single_1.id
      routing_node {
        level = "gateway"
        href  = smc_router.tf_sample_router_1.id
        routing_node {
          level = "any"
          href  = data.smc_href.any_network.id
        }
      }
    }
  }
  routing_node {
    level = "interface"
    name  = "Tunnel Interface 6483"
  }
  routing_node {
    level = "interface"
    name  = "Tunnel Interface 1000"
  }
}

resource "smc_internal_endpoint" "ep_192_168_100_14" {
  id = format("%s#%s", smc_single_fw.tf_single_fw1.link.internal_gateway,
  "*/internal_endpoint/192.168.100.14")
  enabled   = true
  ipsec_vpn = true
}

resource "smc_single_fw" "tf_single_fw2" {
  name           = "tf_single_fw2"
  log_server_ref = data.smc_href.log_server.href
  comment        = var.resource_comment
  nodes {
    firewall_node {
      name   = "tf_single_fw2 node1"
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
          address       = "10.0.50.14"
          network_value = "10.0.50.0/24"
          primary_mgt   = true
        }
      }
    }
  }
  physical_interfaces {
    tunnel_interface {
      interface_id = "4563"
      interfaces {
        single_node_interface {
          address       = "2001:db8:85a3::8a2e:370:7334"
          network_value = "2001:db8:85a3::8a2e:370:7334/128"
          nicid         = "4563"
          nodeid        = 1
        }
      }
      name = "Tunnel Interface 4563"
    }
  }
  physical_interfaces {
    tunnel_interface {
      interface_id = "1234"
      interfaces {
        single_node_interface {
          address       = "172.16.50.1"
          network_value = "172.16.50.1/32"
          nicid         = "1234"
          nodeid        = 1
        }
      }
      name = "Tunnel Interface 1234"
    }
  }
  depends_on = [smc_network.tf_sample_network_single_2, smc_router.tf_sample_router_2]
}

resource "smc_routing_node" "tf_single_fw2" {
  id    = smc_single_fw.tf_single_fw2.link.routing
  level = "engine_cluster"
  routing_node {
    level = "interface"
    name  = "Interface 0"
    routing_node {
      level = "network"
      href  = smc_network.tf_sample_network_single_2.id
      routing_node {
        level = "gateway"
        href  = smc_router.tf_sample_router_2.id
        routing_node {
          level = "any"
          href  = data.smc_href.any_network.id
        }
      }
    }
  }
  routing_node {
    level = "interface"
    name  = "Tunnel Interface 4563"
  }
  routing_node {
    level = "interface"
    name  = "Tunnel Interface 1234"
  }
}

resource "smc_internal_endpoint" "ep_10_0_50_14" {
  id = format("%s#%s", smc_single_fw.tf_single_fw2.link.internal_gateway,
  "*/internal_endpoint/10.0.50.14")
  enabled   = true
  ipsec_vpn = true
}

resource "smc_rbvpn_tunnel" "gre_single_fw1_single_fw2" {
  depends_on     = [smc_internal_endpoint.ep_192_168_100_14, smc_internal_endpoint.ep_10_0_50_14]
  comment        = var.resource_comment
  enabled        = true
  mtu            = 0
  name           = "GRE-tf_single_fw1 to_tf_single_fw2"
  pmtu_discovery = true
  rbvpn_tunnel_side_a {
    endpoint_ref         = data.smc_sub_href.single_fw1_ep.id
    gateway_ref          = data.smc_sub_href.single_fw1_internal_gateway_ref.id
    tunnel_interface_ref = data.smc_sub_href.single_fw1_tunnel_interface_1000_ref.id
  }
  rbvpn_tunnel_side_b {
    ip_address = "10.2.3.1"
  }
  ttl               = 0
  tunnel_encryption = "no_encryption"
  tunnel_mode       = "ipip"
}
