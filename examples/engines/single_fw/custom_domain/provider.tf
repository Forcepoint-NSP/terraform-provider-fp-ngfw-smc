terraform {
  required_providers {
    smc = {
      source  = "forcepoint-nsp/fp-ngfw-smc"
      version = "0.0.1"
    }
  }
}

variable "api_key" {
  type    = string
  default = null
}

variable "url" {
  type    = string
  default = "http://localhost:8082"
}

variable "api_version" {
  type    = string
  default = "7.4"
}

provider "smc" {
  url         = var.url
  api_key     = var.api_key
  api_version = var.api_version
}

provider "smc" {
  alias       = "tfdomain"
  domain      = "tf_domain"
  url         = var.url
  api_key     = var.api_key
  api_version = var.api_version
}
