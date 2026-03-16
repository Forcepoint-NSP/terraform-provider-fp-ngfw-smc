variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "ntp_protocol_agent" {
  name = "NTP"
  type = "protocol"
}

data "smc_href" "dns_protocol_agent" {
  name = "DNS"
  type = "protocol"
}

data "smc_sub_href" "udp_dns_pa_parameters" {
  from_ref = data.smc_href.dns_protocol_agent.id
  match    = "pa_parameters"
}

data "smc_sub_href" "udp_dns_pa_parameters_dns_protocol_enforcement" {
  from_ref = data.smc_sub_href.udp_dns_pa_parameters.id
  match    = "dns_protocol_enforcement"
}

data "smc_sub_href" "udp_dns_pa_parameters_dns_ddns_deny" {
  from_ref = data.smc_sub_href.udp_dns_pa_parameters.id
  match    = "dns_ddns_deny"
}

data "smc_sub_href" "udp_dns_pa_parameters_dns_deny_zone_transfer" {
  from_ref = data.smc_sub_href.udp_dns_pa_parameters.id
  match    = "dns_deny_zone_transfer"
}

data "smc_sub_href" "udp_dns_pa_parameters_dns_google_safesearch" {
  from_ref = data.smc_sub_href.udp_dns_pa_parameters.id
  match    = "dns_google_safesearch"
}

data "smc_sub_href" "udp_dns_pa_parameters_dns_bing_safesearch" {
  from_ref = data.smc_sub_href.udp_dns_pa_parameters.id
  match    = "dns_bing_safesearch"
}

data "smc_sub_href" "udp_dns_pa_parameters_dns_duckduckgo_safesearch" {
  from_ref = data.smc_sub_href.udp_dns_pa_parameters.id
  match    = "dns_duckduckgo_safesearch"
}

data "smc_sub_href" "udp_dns_pa_parameters_dns_youtube_restrict" {
  from_ref = data.smc_sub_href.udp_dns_pa_parameters.id
  match    = "dns_youtube_restrict"
}


resource "smc_udp_service" "ntp_udp" {
  min_dst_port       = 321
  name               = "tf_udp_service_no_pa"
  comment            = var.resource_comment
  protocol_agent_ref = data.smc_href.ntp_protocol_agent.id
}

resource "smc_udp_service" "udp_pa" {
  comment      = var.resource_comment
  max_dst_port = 14545
  min_dst_port = 5648
  name         = "tf_dns_udp_pa"
  pa_values {
    boolean_value {
      parameter_ref = data.smc_sub_href.udp_dns_pa_parameters_dns_duckduckgo_safesearch.id
      value         = "1"
    }
  }
  pa_values {
    boolean_value {
      parameter_ref = data.smc_sub_href.udp_dns_pa_parameters_dns_bing_safesearch.id
      value         = "1"
    }
  }
  pa_values {
    boolean_value {
      parameter_ref = data.smc_sub_href.udp_dns_pa_parameters_dns_google_safesearch.id
      value         = "1"
    }
  }
  pa_values {
    string_value {
      parameter_ref = data.smc_sub_href.udp_dns_pa_parameters_dns_youtube_restrict.id
      value         = "moderate"
    }
  }
  pa_values {
    boolean_value {
      parameter_ref = data.smc_sub_href.udp_dns_pa_parameters_dns_protocol_enforcement.id
      value         = "1"
    }
  }
  pa_values {
    boolean_value {
      parameter_ref = data.smc_sub_href.udp_dns_pa_parameters_dns_ddns_deny.id
      value         = "1"
    }
  }
  pa_values {
    boolean_value {
      parameter_ref = data.smc_sub_href.udp_dns_pa_parameters_dns_deny_zone_transfer.id
      value         = "1"
    }
  }
  protocol_agent_ref = data.smc_href.dns_protocol_agent.id
}

resource "smc_udp_service_group" "udp_service_group" {
  name    = "tf_udp_service_group"
  comment = var.resource_comment
  element = [smc_udp_service.ntp_udp.id, smc_udp_service.udp_pa.id]
}