terraform {
  required_providers {
    smc = {
      source  = "forcepoint-nsp/fp-ngfw-smc"
      version = "0.0.2"
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
  default = null
}

variable "trusted_cert" {
  type    = string
  default = null
}

variable "verify_ssl" {
  type    = bool
  default = false
}

provider "smc" {
  url          = var.url
  api_key      = var.api_key
  api_version  = var.api_version
  trusted_cert = var.trusted_cert
  verify_ssl   = var.verify_ssl
}
