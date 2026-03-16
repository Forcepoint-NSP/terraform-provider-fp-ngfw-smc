variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  engine_name = "tf_single_fw"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "any_network" {
  type = "network"
  name = "Any network"
}

data "smc_href" "active_connection_type" {
  type = "connection_type"
  name = "Active"
}

resource "smc_network" "tf_net_10_0_0_0-8" {
  name         = "Network 10.0.0.0/8"
  comment      = var.resource_comment
  ipv4_network = "10.0.0.0/8"
}

resource "smc_network" "tf_net_172_16_0_0-12" {
  name         = "Network 172.16.0.0/12"
  comment      = var.resource_comment
  ipv4_network = "172.16.0.0/12"
}

resource "smc_network" "tf_net_60_100_0_0-16" {
  name         = "Network 60.100.0.0/16"
  comment      = var.resource_comment
  ipv4_network = "60.100.0.0/16"
}

resource "smc_group" "group" {
  name    = "tf_group_for_engines_routing"
  comment = var.resource_comment
  element = [
    smc_network.tf_net_60_100_0_0-16.id,
    smc_network.tf_net_172_16_0_0-12.id,
    smc_network.tf_net_10_0_0_0-8.id
  ]
}

resource "smc_network" "tf_net_ipv6" {
  name         = "Network 2001:db8:acad:1::/120"
  comment      = var.resource_comment
  ipv6_network = "2001:db8:acad:1::/120"
}

resource "smc_network" "tf_net_ipv6_routing" {
  name         = "Network 2001:db8:acad:2::/120"
  comment      = var.resource_comment
  ipv6_network = "2001:db8:acad:2::/120"
}

resource "smc_network" "tf_net_192_168_101_0-24" {
  name         = "Network 192.168.101.0/24"
  comment      = var.resource_comment
  ipv4_network = "192.168.101.0/24"
}

resource "smc_network" "tf_net_192_168_100_0-24" {
  name         = "Network 168.100.0.0/24"
  comment      = var.resource_comment
  ipv4_network = "192.168.100.00/24"
}

resource "smc_network" "tf_net_100_100_100_0-24" {
  name         = "Network 100.100.100.0/24"
  comment      = var.resource_comment
  ipv4_network = "100.100.100.0/24"
}

resource "smc_network" "tf_net_10_10_10_0-24" {
  name         = "Network 10.10.10.0/24"
  comment      = var.resource_comment
  ipv4_network = "10.10.10.0/24"
}

resource "smc_router" "tf_sample_router" {
  address = "192.168.101.254"
  name    = "Router-192.168.101.254"
  comment = var.resource_comment
}

resource "smc_router" "tf_sample_router2" {
  address = "100.100.100.254"
  name    = "Router-100.100.100.254"
  comment = var.resource_comment
}

resource "smc_netlink" "tf_sample_netlink" {
  name                = "tf_sample_netlink-orange-${local.engine_name}"
  comment             = var.resource_comment
  active_mode_period  = 5
  active_mode_timeout = 1
  connection_type_ref = data.smc_href.active_connection_type.id
  # domain_server_address = []
  gateway_ref = smc_router.tf_sample_router.id
  input_speed = 200000
  network_ref = [smc_network.tf_net_192_168_101_0-24.id]
  # nsp_name = "Orange"
  output_speed         = 200000
  probe_address        = ["8.8.8.8", "8.8.4.4"]
  standby_mode_period  = 3600
  standby_mode_timeout = 30
}

resource "smc_netlink" "tf_sample_netlink2" {
  name                = "tf_sample_netlink-free-${local.engine_name}"
  comment             = var.resource_comment
  active_mode_period  = 5
  active_mode_timeout = 1
  connection_type_ref = data.smc_href.active_connection_type.id
  gateway_ref         = smc_router.tf_sample_router2.id
  input_speed         = 200000
  network_ref         = [smc_network.tf_net_100_100_100_0-24.id]
  # nsp_name = "Free"
  output_speed         = 200000
  probe_address        = ["8.8.8.8", "8.8.4.4"]
  standby_mode_period  = 3600
  standby_mode_timeout = 30
}

resource "smc_single_fw" "tf_single_fw" {
  name           = local.engine_name
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href
  depends_on = [
    smc_router.tf_sample_router,
    smc_router.tf_sample_router2,
    smc_network.tf_net_192_168_101_0-24,
    smc_network.tf_net_100_100_100_0-24,
    smc_network.tf_net_192_168_100_0-24,
    smc_network.tf_net_10_10_10_0-24,
    smc_network.tf_net_ipv6_routing,
    smc_network.tf_net_ipv6,
    smc_group.group,
    smc_netlink.tf_sample_netlink, # needed otherwise destroy failed
    smc_netlink.tf_sample_netlink2 # needed otherwise destroy failed
  ]

  nodes {
    firewall_node {
      name   = "${local.engine_name} node"
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
    physical_interface {
      interface_id           = "1"
      name                   = "Interface 1"
      virtual_engine_vlan_ok = false
      vlan_interfaces {
        interface_id = "1.2000"
        interfaces {
          single_node_interface {
            address       = "192.168.101.1"
            network_value = "192.168.101.0/24"
            nicid         = "1.2000"
            nodeid        = 1
          }
        }
        name = "VLAN 1.2000"
      }
      vlan_interfaces {
        interface_id = "1.100"
        interfaces {
          single_node_interface {
            address       = "100.100.100.100"
            network_value = "100.100.100.0/24"
            nicid         = "1.100"
            nodeid        = 1
          }
        }
        name = "VLAN 1.100"
      }
    }
  }
  physical_interfaces {
    tunnel_interface {
      interface_id = "6483"
      interfaces {
        single_node_interface {
          address       = "2001:db8:acad:1::1"
          network_value = "2001:db8:acad:1::1/128"
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

}

resource "smc_routing_node" "tf_engine_routing" {
  id    = smc_single_fw.tf_single_fw.link.routing
  level = "engine_cluster"
  routing_node {
    level = "interface"
    name  = "Interface 0"
    routing_node {
      level = "network"
      href  = smc_network.tf_net_192_168_100_0-24.id
    }
  }
  routing_node {
    level = "interface"
    name  = "VLAN 1.100"
    routing_node {
      level = "network"
      href  = smc_network.tf_net_100_100_100_0-24.id
      routing_node {
        level = "gateway"
        href  = smc_netlink.tf_sample_netlink2.id
        routing_node {
          level = "any"
          href  = data.smc_href.any_network.id
        }
      }
    }
  }
  routing_node {
    level = "interface"
    name  = "VLAN 1.2000"
    routing_node {
      level = "network"
      href  = smc_network.tf_net_192_168_101_0-24.id
      routing_node {
        level = "gateway"
        href  = smc_netlink.tf_sample_netlink.id
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
    routing_node {
      level = "gateway"
      href  = smc_network.tf_net_ipv6_routing.id
    }
    routing_node {
      href         = smc_group.group.id
      level        = "gateway"
      probe_metric = 2
    }
  }
  routing_node {
    level = "interface"
    name  = "Tunnel Interface 1000"
    routing_node {
      level = "gateway"
      href  = smc_network.tf_net_10_10_10_0-24.id
    }
    routing_node {
      href              = smc_group.group.id
      level             = "gateway"
      probe_metric      = 1
      probe_test        = "next_hop_reachability"
      probe_retry_count = 0
    }
  }
}
