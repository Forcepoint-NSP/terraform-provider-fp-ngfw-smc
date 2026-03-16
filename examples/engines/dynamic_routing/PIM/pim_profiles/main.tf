variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_pim_ipv4_profile" "pim_profile" {
  hello_interval = 45
  joined_prune   = 80
  name           = "tf_pim_profile"
  comment        = var.resource_comment
  pim_multicast_group_entry {
    mapping              = "10.0.1.100"
    mode                 = "pim_sm"
    multicast_ip_network = "224.0.1.0/24"
  }
  pim_multicast_group_entry {
    mapping              = "source1.example.com"
    mode                 = "pim_ssm"
    multicast_ip_network = "232.1.1.0/24"
  }
  pim_multicast_group_entry {
    mapping              = ""
    mode                 = "pim_sm"
    multicast_ip_network = "239.255.0.0/16"
  }
  pim_multicast_group_entry {
    mapping              = "10.0.2.50"
    mode                 = "pim_ssm"
    multicast_ip_network = "232.10.10.0/24"
  }
  smart_multicast_antispoofing = true
  # spt_switch_interval          = 90
  # spt_switch_threshold         = 10
  # spt_switch_threshold_unit    = "packets"
}
