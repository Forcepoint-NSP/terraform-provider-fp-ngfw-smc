variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_host" "tf_host_example" {
  name      = "tf_host_example"
  address   = "192.168.1.44"
  comment   = var.resource_comment
  secondary = ["212.20.1.1", "123.6.5.22"]
}

output "host_href" {
  # you can access to the links of the host like this
  value = smc_host.tf_host_example.link.self
}
