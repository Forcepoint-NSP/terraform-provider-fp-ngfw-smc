variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "tls_profile" {
  type = "tls_profile"
  name = "Stonesoft Update Servers"
}

resource "smc_elasticsearch_cluster" "elasticsearch_cluster" {
  addresses = ["10.3.3.3", "10.3.3.33"]
  name = "tf_elasticsearch_cluster"
  comment = var.resource_comment
  authentication_settings {
    login = "es_login"
    method = "basic"
    password = "superPassword00!!"
  }
  es_enable_cluster_sniffer = true
  es_replica_number = 1
  es_retention_period = 30
  es_shard_number = 0
  indexing_active = true
  port = 9200
  product = "elasticsearch"
  # Not working since attribute [tls_profile] is missing
  # tls_profile = data.smc_href.tls_profile.id
}
