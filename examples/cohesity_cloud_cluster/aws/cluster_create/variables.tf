################################################################################
# Authentication Variables
################################################################################

variable "region" {
  type        = string
  description = "AWS Region to deploy to"
}

variable "iam_role_arn" {
  type        = string
  description = "AWS IAM role to use"
}

variable "use_iam_role" {
  type        = bool
  description = "Whether to use IAM role for Authentication or IAM user credentials"
  default     = false
}

################################################################################
# Deployment Variables
################################################################################

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

variable "subnet_id" {
  type        = string
  description = "ID of the existing Subnet"
}

variable "vpc_id" {
  type        = string
  description = "ID of the vpc where you want the security group created"
  default = ""
}

variable "use_existing_sg" {
  description = "Whether to use an existing Security Group"
  type        = bool
}

variable "existing_sg_id" {
  description = "ID of the existing Security Group (if using one)"
  type        = string
}

variable "sg_name" {
  description = "Name of the new Security Group (if creating one)"
  type        = string
}

variable "rules_file" {
  description = "Path to the JSON rules file for creating new sec grp"
  type        = string
  default     = "./rules.json"
}

variable "security_group_ids" {
  type        = list(string)
  description = <<EOT
List of security group IDs to attach to the instances.
Each ID must correspond to an existing security group in AWS.
EOT
}

variable "image_id" {
  type        = string
  description = "ID of the AMI to use for the Virtual Machine"
}

variable "tags" {
  type        = list(string)
  description = "List of tags in the format ['key:value', 'key:value']"
}

variable "attach_public_ip" {
  type        = bool
  description = "Boolean flag to attach public IPs to each AWS VM."
  default     = false
}

variable "resource_name_prefix" {
  description = <<EOT
Prefix for the names of all created resources (VMs, NICs, Disks, etc.).
EOT
  type        = string
  default     = ""

  # Validation rules for non-null values
  validation {
    condition     = length(var.resource_name_prefix) <= 30
    error_message = <<EOT
The resource_name_prefix must less than or equal to 30 characters long.
EOT
  }

  validation {
    condition = var.resource_name_prefix == "" || (
      can(regex("^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*$", var.resource_name_prefix))
    )
    error_message = <<EOT
The resource_name_prefix must either be an empty string (in which case a
random string will be used as prefix) or contain only letters, numbers,
and hyphens. Hyphens cannot be consecutive, or appear at the start or end.
EOT
  }
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

variable "post_boot_wait" {
  type        = string
  description = "Wait after the VM Boot up for the nexus service to be available"
  default     = "600"
}

################################################################################
# Cluster Create Variables
################################################################################

variable "create_cluster" {
  description = "Whether to issue cluster creation request or not"
  type        = bool
}

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
  description = "Apps subnet"
  type        = string
}

variable "apps_subnet_mask" {
  description = "Apps subnet mask"
  type        = string
}
