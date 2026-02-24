################################################################################
# Deployment Variables
################################################################################

variable "region" {
  type        = string
  description = "OCI Region to deploy to"
}

variable "compartment_id" {
  type        = string
  description = "OCID of the compartment where resources will be created"
}

variable "availability_domain" {
  type        = string
  description = "Availability domain for the instances"
}

variable "num_instances" {
  type        = number
  description = "Number of Virtual Machines to create"
  default     = 1
  validation {
    condition     = var.num_instances > 0
    error_message = "num_instances must be greater than 0."
  }
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

variable "add_random_prefix" {
  type        = bool
  description = "Random 8 character alphanumeric string is added to the names of all resources created"
  default     = false
}

variable "subnet_id" {
  type        = string
  description = "OCID of the existing Subnet"
}

variable "image_id" {
  type        = string
  description = "OCID of the image to use for the Virtual Machine"
}

variable "tags" {
  type        = list(string)
  description = "List of tags in the format ['key:value', 'key:value']"
}

variable "attach_public_ip" {
  type        = bool
  description = "Boolean flag to attach public IPs to each OCI VM."
  default     = false
}

variable "nsg_ids" {
  description = "List of Network Security Group (NSG) OCIDs to attach to created resources. Optional."
  type        = list(string)
  default     = []
}

################################################################################
# Config Variables
################################################################################

variable "config_id" {
  description = "ID of the configuration to choose from the JSON file"
  type        = string
  default     = "2"
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

variable "post_boot_wait" {
  type        = string
  description = "Wait after the VM Boot up for the nexus service to be available"
  default     = "600"
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

variable "issue_cluster_create_cmd" {
  description = <<EOT
Whether to issue the cluster creation command via SSH after VM deployment.
Usually set to true (default) as Terraform is expected to run in the same VCN as the VMs, or public IPs are attached for SSH access.
Set to false if you want to create the cluster manually.
EOT
  type        = bool
  default     = true
}
