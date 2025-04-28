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

resource "cohesity_configure_snmp" "snmp" {
  server = var.server
  snmp_version = var.snmp_version
  agent_port = var.agent_port
  trap_port = var.trap_port

  trap_user {
    user_name     = var.trap_user["user_name"]
    auth_password = var.trap_user["auth_password"]
    auth_protocol = var.trap_user["auth_protocol"]
    priv_password = var.trap_user["priv_password"]
    priv_protocol = var.trap_user["priv_protocol"]
  }

  read_user {
    user_name     = var.read_user["user_name"]
    auth_password = var.read_user["auth_password"]
    auth_protocol = var.read_user["auth_protocol"]
    priv_password = var.read_user["priv_password"]
    priv_protocol = var.read_user["priv_protocol"]
  }
}