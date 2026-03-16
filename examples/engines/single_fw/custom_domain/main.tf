/* this test creates a custom domain 'tf_domain', then connects to it
   using a provider alias and creates a single firewall within that domain 
*/
variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_admin_domain" "tf_domain" {
  name    = "tf_domain"
  comment = "test domain created by terraform"
}

data "smc_href" "log_server" {
  provider = smc.tfdomain
  depends_on = [smc_admin_domain.tf_domain]
  type = "log_server"
  name = "*"
}

resource "smc_single_fw" "tf_single_fw" {
  provider = smc.tfdomain
  depends_on = [smc_admin_domain.tf_domain]

  name           = "tf_single_fw"
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
          address       = "192.168.100.14"
          network_value = "192.168.100.00/24"
          primary_mgt   = true
        }
      }
    }
  }
}
