# Provider dependencies.
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.40.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "~> 0.7" # Specify the version of the time provider
    }
    null = {
        source  = "hashicorp/null"
        version = "~> 3.2.2"
    }
    local = {
        source = "hashicorp/local"
        version = "~> 2.5.1"
     }
  }
}