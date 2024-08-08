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

variable "instance_name" {
  description = "The name of the VM instance"
  type        = string
  default     = "terraform-instance"
}

variable "machine_type" {
  description = "The type of the VM instance"
  type        = string
  default     = "n1-standard-1"
}

variable "disk_image" {
  description = "The image for the boot disk"
  type        = string
  default     = "debian-cloud/debian-10"
}

variable "labels" {
  description = "A map of labels to assign to the instance"
  type        = map(string)
  default     = {
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

variable "ssh_username" {
  description = "The SSH username"
  type        = string
}

variable "ssh_password" {
  description = "The SSH password"
  type        = string
  sensitive   = true
}
