###############################################################################
# Authentication variables
###############################################################################

variable "subscription_id" {
  description = "Azure Subscription ID (Required for all authentication methods)"
  type        = string
}

variable "client_id" {
  description = "Azure Service Principal Client ID (leave empty to use other authentication methods)"
  type        = string
  default     = ""
}

variable "client_secret" {
  description = "Azure Service Principal Client Secret (leave empty to use other authentication methods)"
  type        = string
  default     = ""
}

variable "tenant_id" {
  description = "Azure Tenant ID (leave empty to use other authentication methods)"
  type        = string
  default     = ""
}

variable "use_managed_identity" {
  description = "Set to true to authenticate using Managed Identity"
  type        = bool
  default     = false
}

variable "managed_identity_id" {
  description = "User-assigned Managed Identity ID (optional). Leave empty for system-assigned."
  type        = string
  default     = ""
}

variable "use_azure_cli" {
  description = "Set to true to authenticate using Azure CLI login"
  type        = bool
  default     = false
}

variable "environment" {
  description = "The Azure cloud environment (public, usgovernment, china, german). Defaults to 'public'."
  type        = string
  default     = "public"

  validation {
    condition     = contains(["public", "usgovernment", "china", "german"], var.environment)
    error_message = "The environment must be one of: 'public', 'usgovernment', 'china', or 'german'."
  }
}

###############################################################################
# Security Group Variables
###############################################################################

variable "resource_group_name" {
  description = "Name of the resource group where the security group will be created."
  type        = string
}

variable "region" {
  description = "Azure region for the resources."
  type        = string
}

variable "resource_name_prefix" {
  description = <<EOT
Prefix for the names of all created resources (VMs, NICs, Disks, etc.).
This value is mandatory. It must be greater than 0 and less than or equal to 30 characters long.
It must contain only lowercase letters, numbers, and hyphens, must start with a lowercase letter, and cannot end with a hyphen.
EOT

  type = string

  validation {
    condition = (
      length(var.resource_name_prefix) > 0 &&
      length(var.resource_name_prefix) <= 30 &&
      can(regex("^[a-z][a-z0-9-]*[a-z0-9]$", var.resource_name_prefix))
    )
    error_message = "resource_name_prefix is mandatory, must be 1-30 chars, start with a lowercase letter, contain only lowercase letters, numbers, and hyphens, and not end with a hyphen."
  }
}

variable "tags" {
  type        = list(string)
  description = "List of tags in the format ['key:value', 'key:value']"
}

variable "rules_file" {
  type        = string
  description = "Path to the JSON rules file for security group rules."
  default     = "./rules.json"
}

variable "cidr" {
  description = "CIDR block that will be used as source_address_prefix for Outbound rules and destination_address_prefix for Inbound rules."
  type        = string
  default     = "10.4.0.0/24"
}
