terraform {
  required_version = ">= 1.11"
  required_providers {
    cohesity = {
      source  = "stegraab/cohesity"
      version = "2.1.3-stegra.1"
    }
  }
}

# Supply COHESITY_API_KEY through the execution environment.
provider "cohesity" {
  endpoint           = "https://backup.example.com"
  ca_certificate_pem = file("root-ca.pem")
}

resource "cohesity_web_server_certificate" "main" {
  certificate_pem     = file("server-chain.pem")
  private_key_pem_wo  = file("server-key.pem")
  private_key_version = 1
  restart_iris        = true
}
