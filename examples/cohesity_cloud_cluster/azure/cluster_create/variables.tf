################################################################################
# Authentication Variables
################################################################################

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
    - Managed Identity: User-assigned identity client ID (optional, only required if VM has multiple user-assigned identities)
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

################################################################################
# Deployment Variables
################################################################################

variable "num_instances" {
  type        = number
  description = "Number of Virtual Machines to create"
  default     = 1
}

variable "resource_group_name" {
  description = "Name of the existing resource group where resource are created (leave empty to create a new one)"
  type        = string
  default     = ""
}

variable "add_random_prefix" {
  type        = bool
  description = "Random 8 character alphanumeric string is added to the names of all resources created"
  default     = false
}

variable "region" {
  type        = string
  description = "Azure Region to deploy to"
}

variable "availability_zone" {
  description = <<EOT
The Availability Zone where the Virtual Machine will be deployed. Example:
'1', '2', or '3'
EOT
  type        = string
  default     = null
}

variable "vnet_id" {
  type        = string
  description = "ID of the existing Virtual Network"
}

variable "subnet_id" {
  type        = string
  description = "ID of the existing Subnet"
}

variable "network_security_group_id" {
  description = <<EOT
ID of the Network Security Group to attach to the network interface. Leave
empty to skip attaching an NSG.
EOT
  type        = string
  default     = null
}

variable "vhd_uri" {
  type        = string
  description = "The URI of the VHD file to use for creating the OS disk"
}

variable "storage_account_id" {
  type        = string
  description = "The ID of the Azure Storage Account where the VHD is stored"
}

variable "tags" {
  type        = list(string)
  description = "List of tags in the format ['key:value', 'key:value']"
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


variable "attach_public_ip" {
  description = "Boolean flag to attach public IPs to each Azure VM."
  type        = bool
  default     = false
}

variable "user_assigned_managed_identities_to_attach" {
  description = "List of user-assigned managed identities to attach to the Azure VMs. Leave empty for no managed identities to attach."
  type        = list(string)
  default     = []
}


################################################################################
# Config Variables
################################################################################

variable "config_id" {
  description = "ID of the configuration to choose from the JSON file"
  type        = string
  default     = "3"
}

variable "config_file" {
  description = "Path to the JSON configuration file"
  type        = string
  default     = "./configs.json"
}

variable "custom_config" {
  description = <<EOT
Custom configuration of the VM which overrides the config provided by the
config_id
EOT
  type = object({
    InstanceType                = string
    SSDTierNumDisks             = number
    SSDTierDiskSizeinGB         = number
    SSDTierDiskType             = string
    SSDTierDiskIops             = number
    SSDTierDiskThroughputinMBps = number
    HDDTierNumDisks             = number
    HDDTierDiskSizeinGB         = number
    HDDTierDiskType             = string
    HDDTierDiskIops             = number
    HDDTierDiskThroughputinMBps = number
  })
  default = null
}

################################################################################
# Cluster Create Variables
################################################################################

variable "cluster_name" {
  description = "Name given to the Cohesity cluster"
  type        = string
}

variable "dns_server_ips" {
  description = "Comma separated DNS server IPs"
  type        = string
}

variable "ntp_servers" {
  description = "Comma separated NTP Servers"
  type        = string
}

variable "domain_names" {
  description = "Domain Name"
  type        = string
}

variable "apps_subnet" {
  description = "A private IPv4 subnet for the internal overlay network of the kubernetes cluster to run the apps infrastructure."
  type        = string
  default     = "192.168.0.0"
}

variable "apps_subnet_mask" {
  description = "A private IPv4 subnet mask for the internal overlay network of the kubernetes cluster to run the apps infrastructure."
  type        = string
  default     = "255.255.0.0"
}

variable "post_boot_wait" {
  type        = string
  description = "Wait after the VM Boot up for the nexus service to be available"
  default     = "600"
}

variable "issue_cluster_create_cmd" {
  description = <<EOT
Whether to issue the cluster creation command via SSH after VM deployment.
Usually set to true (default) as Terraform is expected to run in the same VNet as the VMs, or public IPs are attached for SSH access.
Set to false if you want to create the cluster manually.
EOT
  type        = bool
  default     = true
}

variable "customer_managed_disk_encryption_id" {
  description = "(Optional) The Azure Disk Encryption Set (DES) resource ID to use for customer-managed disk encryption. If not set, Azure-managed keys are used."
  type        = string
  default     = ""
}
