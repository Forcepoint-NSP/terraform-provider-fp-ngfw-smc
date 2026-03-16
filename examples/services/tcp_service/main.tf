variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "ssm_tcp_proxy" {
  name = "SSM TCP Proxy"
  type = "protocol"
}

data "smc_href" "https_default_tls_inspection_policy" {
  name = "Default HTTPS Inspection Exceptions"
  type = "tls_inspection_policy"
}

data "smc_href" "imaps_protocol" {
  name = "IMAPS"
  type = "protocol"
}

data "smc_sub_href" "imaps_protocol_pa_parameters" {
  from_ref = data.smc_href.imaps_protocol.id
  match    = "pa_parameters"
}

data "smc_sub_href" "imaps_tls_inspection" {
  from_ref = data.smc_sub_href.imaps_protocol_pa_parameters.id
  match    = "tls_inspection"
}

data "smc_sub_href" "impas_tls_policy" {
  from_ref = data.smc_sub_href.imaps_protocol_pa_parameters.id
  match    = "tls_policy"
}


resource "smc_tcp_service" "tf_tcp_service" {
  max_dst_port       = 1554
  max_src_port       = 5642
  min_dst_port       = 789
  min_src_port       = 165
  name               = "tf_tcp_service"
  comment            = var.resource_comment
  protocol_agent_ref = data.smc_href.ssm_tcp_proxy.id
}


resource "smc_tcp_service" "tf_tcp_service_pa" {
  min_dst_port = 456
  name         = "tf_tcp_service_pa"
  comment      = var.resource_comment
  pa_values {
    string_value {
      parameter_ref = data.smc_sub_href.imaps_tls_inspection.id
      value         = "Yes"
    }
  }
  pa_values {
    tls_inspection_policy_value {
      parameter_ref             = data.smc_sub_href.impas_tls_policy.id
      tls_inspection_policy_ref = data.smc_href.https_default_tls_inspection_policy.id
    }
  }
  protocol_agent_ref = data.smc_href.imaps_protocol.id
}

resource "smc_tcp_service_group" "tcp_service_group" {
  name    = "tf_tcp_service_group"
  comment = var.resource_comment
  element = [smc_tcp_service.tf_tcp_service.id, smc_tcp_service.tf_tcp_service_pa.id]
}
