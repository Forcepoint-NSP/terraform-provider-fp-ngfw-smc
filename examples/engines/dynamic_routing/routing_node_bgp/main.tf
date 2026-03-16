variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "any_network" {
  type = "network"
  name = "Any network"
}

resource "smc_network" "tf_sample_network192" {
  name         = "tf_net-192.168.100.0/24"
  ipv4_network = "192.168.100.0/24"
  comment      = var.resource_comment
}

resource "smc_network" "tf_sample_network172" {
  name         = "tf_net-172.17.17.0/24"
  ipv4_network = "172.17.17.0/24"
  comment      = var.resource_comment
}

resource "smc_router" "tf_sample_router" {
  address = "192.168.100.254"
  name    = "tf_sample_router"
  comment = var.resource_comment
}

resource "smc_network" "tf_sample_network10" {
  name         = "tf_net-10.1.1.0/24"
  ipv4_network = "10.1.1.0/24"
  comment      = var.resource_comment
}

resource "smc_router" "tf_sample_router2" {
  address = "192.168.101.254"
  name    = "tf_sample_router2"
  comment = var.resource_comment
}

resource "smc_network" "network1" {
  name         = "tf_network1"
  ipv4_network = "192.168.10.0/24"
  comment      = var.resource_comment
}

resource "smc_network" "network2" {
  name         = "tf_network2"
  ipv4_network = "172.16.0.0/20"
  comment      = var.resource_comment
}

resource "smc_bgp_profile" "bgp_profile" {
  aggregation_entry {
    mode   = "as_set_and_summary"
    subnet = smc_network.network1.id
  }
  bmp_entry {
    bmp_address                = "192.168.10.1"
    bmp_connect_through_master = true
    bmp_port                   = 179
  }
  bmp_entry {
    bmp_address                = "10.100.100.14"
    bmp_connect_through_master = false
    bmp_port                   = 1179
  }
  distance_entry {
    distance = 220
    subnet   = smc_network.network1.id
  }
  distance_entry {
    distance = 255
    subnet   = smc_network.network2.id
  }
  external = 40
  internal = 220
  local    = 250
  name     = "tf_bgp_profile"
  port     = 179
  comment  = var.resource_comment
  redistribution_entry {
    enabled     = false
    filter_type = "none"
    type        = "kernel"
  }
  redistribution_entry {
    enabled     = true
    filter_type = "none"
    metric      = 120
    type        = "static"
  }
  redistribution_entry {
    enabled     = false
    filter_type = "none"
    type        = "connected"
  }
  redistribution_entry {
    enabled     = false
    filter_type = "none"
    type        = "ospfv2"
  }
}

resource "smc_ip_access_list" "ip_access_list" {
  comment = var.resource_comment
  entries {
    ip_access_list_entry {
      action = "deny"
      subnet = "172.16.16.0/21"
    }
  }
  entries {
    ip_access_list_entry {
      action = "permit"
      subnet = "192.168.100.0/24"
    }
  }
  name = "tf_ip_access_list"
}

resource "smc_as_path_access_list" "as_path_access_list" {
  comment = var.resource_comment
  entries {
    as_path_access_list_entry {
      action     = "permit"
      expression = "^65001_"
    }
  }
  entries {
    as_path_access_list_entry {
      action     = "deny"
      expression = "_65002$"
    }
  }
  name = "tf_as_path_access_list"
}

resource "smc_bgp_connection_profile" "bgp_connection_profile" {
  connect            = 150
  md5_password       = "TheSuperPassword00!!"
  name               = "tf_bgp_connection_profile"
  session_hold_timer = 190
  session_keep_alive = 70
  comment            = var.resource_comment
}

resource "smc_bgp_peering" "bgp_peering" {
  bfd_enabled               = false
  comment                   = var.resource_comment
  connected_check           = "disabled"
  connection_profile        = smc_bgp_connection_profile.bgp_connection_profile.id
  default_originate         = false
  dont_capability_negotiate = false
  inbound_aspath_filter     = smc_as_path_access_list.as_path_access_list.id
  inbound_ip_filter         = smc_ip_access_list.ip_access_list.id
  local_as_option           = "not_set"
  max_prefix_option         = "not_enabled"
  name                      = "tf_bgp_peering"
  next_hop_self             = true
  orf_option                = "disabled"
  outbound_aspath_filter    = smc_as_path_access_list.as_path_access_list.id
  outbound_ip_filter        = smc_ip_access_list.ip_access_list.id
  override_capability       = false
  remove_private_as         = false
  route_reflector_client    = true
  send_community            = "no"
  soft_reconfiguration      = true
  ttl_option                = "disabled"
}

resource "smc_autonomous_system" "autonomous_system_580" {
  as_number = "580"
  name      = "tf_autonomous_system_580"
  comment   = var.resource_comment
}

resource "smc_route_map" "route_map" {
  name    = "tf_route_map"
  comment = var.resource_comment
}

resource "smc_autonomous_system" "autonomous_system_578" {
  as_number = "578"
  name      = "tf_autonomous_system_578"
  comment   = var.resource_comment
}

resource "smc_external_bgp_peer" "external_bgp_peer" {
  name          = "tf_external_bgp_peer"
  neighbor_as   = smc_autonomous_system.autonomous_system_580.id
  neighbor_ip   = "10.10.14.10"
  neighbor_port = 179
  comment       = var.resource_comment
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw"
  log_server_ref = data.smc_href.log_server.href
  comment        = var.resource_comment
  depends_on = [
    smc_router.tf_sample_router,
    smc_network.tf_sample_network192,
    smc_router.tf_sample_router2,
    smc_network.tf_sample_network10,
    smc_network.tf_sample_network172,
    smc_external_bgp_peer.external_bgp_peer,
    smc_bgp_peering.bgp_peering
  ]

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
          network_value = "192.168.100.0/24"
          primary_mgt   = true
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
          address       = "172.17.17.14"
          network_value = "172.17.17.0/24"
        }
      }
    }
  }
  dynamic_routing {
    bgp {
      announced_ne_setting {
        announced_ne_ref = smc_network.tf_sample_network10.id
        announced_rm_ref = smc_route_map.route_map.id
      }
      bgp_as_ref      = smc_autonomous_system.autonomous_system_580.id
      bgp_profile_ref = smc_bgp_profile.bgp_profile.id
      enabled         = true
    }
    ospfv2 {
      enabled = false
    }
  }

}

resource "smc_routing_node" "tf_single_fw" {
  id    = smc_single_fw.tf_single_fw.link.routing
  level = "engine_cluster"
  routing_node {
    level = "interface"
    name  = "Interface 0"
    routing_node {
      level = "network"
      href  = smc_network.tf_sample_network192.id
      routing_node {
        level = "gateway"
        href  = smc_router.tf_sample_router.id
        routing_node {
          level = "any"
          href  = data.smc_href.any_network.id
        }
      }
    }
  }
  routing_node {
    level = "interface"
    name  = "Interface 1"
    routing_node {
      level = "network"
      href  = smc_network.tf_sample_network172.id
      routing_node {
        level = "gateway"
        href  = smc_bgp_peering.bgp_peering.id
        routing_node {
          level = "any"
          href  = smc_external_bgp_peer.external_bgp_peer.id
        }
      }
    }
  }
}
