variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
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
  comment = var.resource_comment
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
