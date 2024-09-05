variable "cohesity_username" {
  description = "Cohesity IRIS username"
  type        = string
}

variable "cohesity_password" {
  description = "Cohesity IRIS user password"
  type        = string
}

variable "support_password" {
  description = "Cohesity support user password"
  type        = string
}

variable "credential_file" {
  description = "The file path of the file containing credentials to access the GCP project"
  type        = string
}
variable "project" {
  description = "The GCP project ID"
  type        = string
}
variable "region" {
  description = "The GCP region"
  type        = string
  default     = "us-west1"
}

variable "zone" {
  description = "The GCP zone"
  type        = string
  default     = "us-west1-a"
}
variable "cluster_name" {
  description = "The name of the Cluster instance"
  type        = string
}
variable "num_instances" {
  description = "The number of instances to create"
  type        = number

  validation {
    condition     = var.num_instances > 0
    error_message = "The number of instances must be greater than 0."
  }
}
variable "machine_type" {
  description = "The type of the VM instance"
  type        = string
  default     = "n1-standard-32"
}
variable "disk_image" {
  description = "The image for the boot disk"
  type        = string
  default     = "debian-cloud/debian-10"
}
variable "labels" {
  description = "A map of labels to assign to the instance"
  type        = map(string)
  default = {
  }
}
variable "network" {
  description = "GCP network name"
  type        = string
}

variable "subnetwork" {
  description = "GCP network name"
  type        = string
}

variable "allot_external_ip" {
  description = "Whether to attach public IPs or not"
  type        = bool
}
variable "dns-server-ips" {
  description = "Comma seperated DNS server IPs"
  type        = string
}

variable "ntp-servers" {
  description = "Comma seperated NTP servers"
  type        = string
}

variable "domain-names" {
  description = "Domain Name"
  type        = string
}
variable "add-apps" {
  description = "whether to allow apps or not, defaults to false"
  type        = bool
  default     = false
}
variable "apps-subnet" {
  description = "Apps subnet"
  type        = string
  default     = ""
}

variable "apps-subnet-mask" {
  description = "Apps subnet mask"
  type        = string
  default     = ""
}
variable "add_gcp_external_target" {
  description = "whether to add a gcp-external-target"
  type        = bool
}
variable "gcp_external_target" {
  description = "gcp external target information"
  type = object({
    project_id                   = string
    client_email_id              = string
    client_private_key_file_path = string
    bucket_name                  = string
    tier_type                    = string
  })

}
