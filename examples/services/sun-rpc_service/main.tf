variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_rpc_service" "tf_rpc_service" {
  name           = "tf_rpc_service"
  comment        = var.resource_comment
  program_number = "14"
  rpc_version    = "2"
  transport      = "both"
}

data "smc_href" "rpc_service_sample" {
  name = "Sample"
  type = "rpc_service"
}

resource "smc_rpc_service_group" "tf_rpc_service_group" {
  name    = "tf_rpc_service_group"
  comment = var.resource_comment
  element = [
    smc_rpc_service.tf_rpc_service.id,
    data.smc_href.rpc_service_sample.id
  ]
}