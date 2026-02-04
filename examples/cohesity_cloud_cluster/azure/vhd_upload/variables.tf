###############################################################################
# Authentication Variables
###############################################################################

variable "auth_method" {
  description = "Authentication method to use: 'service_principal', 'managed_identity', or 'azure_cli'"
  type        = string

  validation {
    condition     = contains(["service_principal", "managed_identity", "azure_cli"], var.auth_method)
    error_message = "auth_method must be one of: 'service_principal', 'managed_identity', or 'azure_cli'."
  }
}

variable "subscription_id" {
  description = "Azure Subscription ID (Required for all authentication methods)"
  type        = string
}

variable "client_id" {
  description = <<-EOT
    Client ID for authentication:
    - Service Principal: Application (client) ID (required)
    - Managed Identity: User-assigned identity client ID (only required if VM has multiple user-assigned identities)
    - Azure CLI: Not used (leave empty)
  EOT
  type        = string
  default     = ""
}

variable "client_secret" {
  description = "Azure Service Principal Client Secret (required when auth_method = 'service_principal', leave empty otherwise)"
  type        = string
  default     = ""
  sensitive   = true
}

variable "tenant_id" {
  description = "Azure Tenant ID (required when auth_method = 'service_principal', leave empty otherwise)"
  type        = string
  default     = ""
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
  description = "Set to false to not log the storage account key"
  type        = bool
  default     = true
}

variable "vhd_upload_parallelism" {
  description = "Number of parallel threads for VHD upload. Default is 2."
  type        = number
  default     = 2
}
