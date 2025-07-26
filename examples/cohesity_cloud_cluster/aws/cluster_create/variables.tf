################################################################################
# Authentication Variables
################################################################################

variable "region" {
  type        = string
  description = "AWS Region to deploy to"
}

variable "iam_role_arn" {
  description = <<EOT
(Optional) The ARN of the IAM Role to assume for AWS authentication.
If specified, Terraform will use this role for all AWS API calls instead of the default credentials.
Leave blank to use the default AWS authentication mechanism (environment variables, profile, etc).
EOT
  type        = string
  default     = ""
}

variable "profile" {
  description = <<EOT
(Optional) AWS CLI named profile to use for authentication. Set this if you want to use a specific profile from ~/.aws/credentials or ~/.aws/config. If not set, the default profile or other authentication methods will be used.
EOT
  type        = string
  default     = ""
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

variable "security_group_ids" {
  type        = list(string)
  default     = []
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
variable "enforce_imdsv2" {
  description = <<EOT
Whether to enforce AWS Instance Metadata Service Version 2 (IMDSv2).

- true: Enforces usage of IMDSv2 (http_tokens = "required")
        which improves security by requiring session tokens.
- false: Allows IMDSv1 (http_tokens = "optional") for legacy compatibility.

EOT
  type        = bool
  default     = false
}

variable "kms_key_id" {
  description = "(Optional) The ARN of the KMS key to use for EBS volume encryption. If not set, AWS default EBS encryption is used."
  type        = string
  default     = ""
}

variable "encrypt_ebs_volumes" {
  description = "Whether to encrypt EBS volumes. If false, volumes are not encrypted and kms_key_id is ignored. If true, kms_key_id is used if set, otherwise AWS default EBS encryption is used."
  type        = bool
  default     = true
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
Usually set to true (default) as Terraform is expected to run in the same VPC as the VMs, or public IPs are attached for SSH access.
Set to false if you want to create the cluster manually.
EOT
  type        = bool
  default     = true
}
