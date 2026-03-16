variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "active_connection_type" {
  type = "connection_type"
  name = "Active"
}
resource "smc_network" "tf_sample_network" {
  name         = "tf_sample_network"
  comment      = var.resource_comment
  broadcast    = false
  ipv4_network = "10.20.30.0/24"
}

resource "smc_router" "tf_example_router" {
  address = "46.200.30.80"
  name    = "tf_example_router"
  comment = var.resource_comment
  third_party_monitoring {
    netflow   = false
    snmp_trap = false
  }
}

resource "smc_netlink" "tf_sample_netlink" {
  name                = "tf_sample_netlink"
  comment             = var.resource_comment
  active_mode_period  = 5
  active_mode_timeout = 1
  connection_type_ref = data.smc_href.active_connection_type.id
  # domain_server_address = []
  gateway_ref = smc_router.tf_example_router.id
  input_speed = 200000
  network_ref = [smc_network.tf_sample_network.id]
  # nsp_name = ""
  output_speed         = 200000
  probe_address        = ["10.20.30.29", "10.20.30.34", "8.8.8.8", "8.8.4.4"]
  standby_mode_period  = 3600
  standby_mode_timeout = 30
}


# todo full example with smc_outbound_multilink
# resource "smc_outbound_multilink" "h0384_out" {
#   multilink_member {
#     ip_range = "185.218.32.254"
#     netlink_ref = "http://localhost:8082/7.4/elements/netlink/268570525"
#     netlink_role = "active"
#     network_ref = "http://localhost:8082/7.4/elements/network/268495623"
#   }
#   multilink_member {
#     ip_range = "192.168.45.30"
#     netlink_ref = "http://localhost:8082/7.4/elements/netlink/268570527"
#     netlink_role = "standby"
#     network_ref = "http://localhost:8082/7.4/elements/network/1495"
#   }
#   multilink_method = "rtt"
#   multilink_qos_class = []
#   name = "H0384-OUT"
#   retries = 2
#   timeout =3600
# }
