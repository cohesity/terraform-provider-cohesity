# variables.tf
###############################################################################
# Authentication Variables
###############################################################################
variable "credentials_file_path" {
  description = "(Optional) The path to the GCP credentials file."
  type        = string
  default     = ""
}

variable "impersonate_service_account" {
  description = "(Optional) The email address of the GCP service account to impersonate for all operations. If set, Terraform will act as this service account (requires Service Account Token Creator role on it). Leave empty to not impersonate."
  type        = string
  default     = ""
}
###############################################################################
# Deployment Variables
###############################################################################
variable "project_id" {
  description = "The ID of the GCP project"
  type        = string
}

variable "num_instances" {
  type        = number
  description = "Number of Virtual Machines to create"
  default     = 1
}

variable "add_random_prefix" {
  type        = bool
  description = "Random 8 character alphanumeric string is added to the names of all resources created"
  default     = false
}

variable "region" {
  type        = string
  description = "GCP Region to deploy to"
}

variable "zone" {
  description = "The zone within the region to deploy instances"
  type        = string
}

variable "vnet_name" {
  type        = string
  description = "Name of the existing Virtual Network where instances will be created"
}

variable "subnet_name" {
  type        = string
  description = "Name of the existing Subnet where instances will be created"
}

variable "image_id" {
  type        = string
  description = "ID of the image to use for the Virtual Machine (format: 'projects/<project>/global/images/<image>')"
}

variable "labels" {
  type        = list(string)
  description = "List of labels in the format ['key:value', 'key:value']"
}

variable "goog_partner_solution1" {
  type        = string
  description = "Google Partner Solution label 1 value"
  default     = "isol_psn_0014m00001h33qrqai_cohesitydatacloud"
}

variable "goog_partner_solution2" {
  type        = string
  description = "Google Partner Solution label 2 value"
  default     = "isol_psn_0014m00001h34k5qaa_cohesitydatacloud"
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
  description = "Boolean flag to attach public IPs to each GCP VM."
  type        = bool
  default     = false
}

variable "service_account_email" {
  description = "Optional: Pre-existing service account email to attach to the VM"
  type        = string
  default     = ""
}

variable "network_tags" {
  description = "Optional: List of network tags to apply to the GCP VMs (used for firewall rules, etc.)"
  type        = list(string)
  default     = []
}

variable "customer_managed_kms_key" {
  description = "(Optional) The KMS key self link or resource ID to use for disk encryption. If not set, Google-managed keys are used."
  type        = string
  default     = ""
}

###############################################################################
# Config Variables
###############################################################################
variable "config_id" {
  description = "ID of the configuration to choose from the JSON file"
  type        = string
}
variable "config_file" {
  description = "Path to the JSON configuration file"
  type        = string
  default     = "./configs.json"
}
variable "custom_config" {
  description = "Custom configuration of the VM which overrides the config provided by the config_id"
  type = object({
    InstanceCPUCount    = string
    SSDTierNumDisks     = number
    SSDTierDiskSizeinGB = number
    SSDTierDiskType     = string
    HDDTierNumDisks     = number
    HDDTierDiskSizeinGB = number
    HDDTierDiskType     = string
  })
  default = null
}

variable "machine_family_type" {
  type        = string
  description = <<EOT
The machine family type for the instance.
Example: "n2d-standard", "n2-standard". This represents the base machine family and is combined with InstanceCPUCount to form the machine type.
EOT
}

variable "post_boot_wait" {
  type        = string
  description = "Wait after the VM Boot up for the nexus service to be available"
  default     = "600"
}

###############################################################################
# Cluster Create Variables
###############################################################################

variable "cluster_name" {
  description = "Name given to the Cohesity cluster"
}

variable "dns_server_ips" {
  description = "Comma seperated DNS server IPs"
  type        = string
}

variable "ntp_servers" {
  description = "Comma seperated NTP Server"
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

variable "issue_cluster_create_cmd" {
  description = <<EOT
Whether to issue the cluster creation command via SSH after VM deployment.
Usually set to true (default) as Terraform is expected to run in the same VPC as the VMs, or public IPs are attached for SSH access.
Set to false if you want to create the cluster manually.
EOT
  type        = bool
  default     = true
}
