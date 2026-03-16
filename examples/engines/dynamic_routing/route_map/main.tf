variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_router" "tf_example_router" {
  address = "46.200.30.80"
  name    = "tf_example_router"
  third_party_monitoring {
    netflow   = false
    snmp_trap = false
  }
  comment = var.resource_comment
}

resource "smc_community_access_list" "community_access_list" {
  comment = var.resource_comment
  entries {
    community_access_list_entry {
      action    = "deny"
      community = "178"
    }
  }
  entries {
    community_access_list_entry {
      action    = "permit"
      community = "1022"
    }
  }
  name = "tf_community_access_list"
  type = "standard"
}


resource "smc_route_map" "route_map" {
  name    = "tf_route_map"
  comment = var.resource_comment
}

resource "smc_route_map_rule" "rule_1" {
  from_ref    = smc_route_map.route_map.link.route_map_rules
  action      = "permit"
  finish      = false
  is_disabled = false
  match_condition {
    metric = 10
    rank   = 1
    type   = "metric"
  }
  name = "route_map_rule_1"
  rank = 1.0
  route_entry_settings {
    as_path_type                  = "dont_modify"
    community_type                = "dont_modify"
    extended_community_entry_type = "rt"
    extended_community_type       = "dont_modify"
    ipv4_next_hop {
      next_hop_peer_address = false
      next_hop_ref          = smc_router.tf_example_router.id
    }
    ipv6_global_next_hop {
      next_hop_peer_address = false
    }
    ipv6_local_next_hop {
      next_hop_peer_address = false
    }
  }
}

resource "smc_route_map_rule" "rule_2" {
  from_ref    = smc_route_map_rule.rule_1.link.add_after
  action      = "deny"
  finish      = true
  is_disabled = false
  match_condition {
    access_list_ref = smc_community_access_list.community_access_list.id
    rank            = 1
    type            = "element"
  }
  name = "rule_2"
  rank = 2.0
  route_entry_settings {
    as_path_type                  = "dont_modify"
    community_type                = "dont_modify"
    extended_community_entry_type = "rt"
    extended_community_type       = "dont_modify"
    ipv4_next_hop {
      next_hop_peer_address = false
    }
    ipv6_global_next_hop {
      next_hop_peer_address = false
    }
    ipv6_local_next_hop {
      next_hop_peer_address = false
    }
  }
}
