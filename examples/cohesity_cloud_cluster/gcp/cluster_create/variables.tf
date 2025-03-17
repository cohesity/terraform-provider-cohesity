# variables.tf
###############################################################################
# Authentication Variables
###############################################################################
variable "credentials_file_path" {
  description = "The path to the GCP credentials file"
  type        = string
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

  validation {
    condition     = contains(data.google_compute_zones.available_zones.names, var.zone)
    error_message = "Invalid zone '${var.zone}' for region '${var.region}'. Available zones are: ${join(", ", data.google_compute_zones.available_zones.names)}."
  }
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

variable "tags" {
  type        = list(string)
  description = "List of tags in the format ['key:value', 'key:value']"
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
      can(regex("^[a-z]([-a-z0-9]*[a-z0-9])?", var.resource_name_prefix))
    )
    error_message = <<EOT
The resource_name_prefix must either be an empty string (in which case a
random string will be used as prefix) or contain only lowecase letters, numbers,
and hyphens. Must start with lowecase letter and cannot end with a Hyphen.
EOT
  }
}

variable "attach_public_ip" {
  description = "Boolean flag to attach public IPs to each GCP VM."
  type        = bool
  default     = false
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
variable "create_cluster" {
  description = "Whether to issue cluster creation request or not"
  type        = bool
}

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
  description = "Apps subnet"
  type        = string
}

variable "apps_subnet_mask" {
  description = "Apps subnet mask"
  type        = string
}
