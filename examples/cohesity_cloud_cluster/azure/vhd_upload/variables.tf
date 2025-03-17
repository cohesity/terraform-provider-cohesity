###############################################################################
# Authentication Variables
###############################################################################

variable "subscription_id" {
  description = "Azure Subscription ID (Required for all authentication methods)"
  type        = string
}

# Service Principal Authentication
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

# Managed Identity Authentication
variable "use_managed_identity" {
  description = "Set to true to authenticate using Managed Identity"
  type        = bool
  default     = false
}

# Azure CLI Authentication
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
# Required Variables
###############################################################################

variable "region" {
  description = "Azure region where the storage account will be created"
  type        = string
}

variable "resource_name_prefix" {
  description = <<EOT
  Prefix for naming resources like resource group, storage account, and container.
  Must be between 3-20 characters, contain only lowercase letters & numbers,
  and cannot start with a number.
  EOT
  type        = string

  validation {
    condition     = length(var.resource_name_prefix) >= 3 && length(var.resource_name_prefix) <= 20
    error_message = "The resource_name_prefix must be between 3 and 20 characters."
  }

  validation {
    condition     = can(regex("^[a-z][a-z0-9]*$", var.resource_name_prefix))
    error_message = "The resource_name_prefix must contain only lowercase letters and numbers, and cannot start with a number."
  }
}

variable "resource_group_name" {
  description = "Name of the existing resource group (leave empty to create a new one)"
  type        = string
  default     = ""
}

variable "storage_account_name" {
  description = "Name of the existing storage account (leave empty to create a new one)"
  type        = string
  default     = ""
}

variable "container_name" {
  description = "Name of the existing container in the storage account (leave empty to create a new one)"
  type        = string
  default     = ""
}

variable "local_vhd_path" {
  description = "Path to the local VHD file to upload"
  type        = string
}

variable "azure_vhd_utils_path" {
  description = "Path to the azure-vhd-utils executable"
  type        = string
}

variable "log_storage_account_key" {
  description = "Set to true to allow Terraform to log the storage account key (not recommended for security reasons)"
  type        = bool
  default     = false
}

variable "vhd_upload_parallelism" {
  description = "Number of parallel threads for VHD upload. Default is 2."
  type        = number
  default     = 2
}
