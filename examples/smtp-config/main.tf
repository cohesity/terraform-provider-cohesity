# Provider dependencies.
terraform {
  required_providers {
    cohesity = {
      source  = "cohesity/cohesity"
      version = "2.1.2"
    }
  }
}

provider "cohesity" {
  cluster_vip      = var.cluster_vip
  cluster_username = var.cohesity_username
  cluster_password = var.cohesity_password
  support_password = var.support_password
  cluster_domain   = "LOCAL"
}

resource "cohesity_configure_smtp" "smtp" {
  smtp_server = var.smtp_server
  port = var.port
  sender_email_address = var.sender_email_address
  username = var.username
  password = var.password
  use_ssl = var.use_ssl
}