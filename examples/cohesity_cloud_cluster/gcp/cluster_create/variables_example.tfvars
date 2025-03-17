################################################################################
# Authentication Variables
################################################################################

# Replace these placeholders with the actual Service Principal credentials
credentials_file_path = "replace-with-the-credentials-file"

################################################################################
# Deployment Variables
################################################################################

# Project under which the resources will be created
project_id = "replace-with-project-id"

# Number of Virtual Machines to create
num_instances = 1

# GCP Region for deployment
# Example: "us-west1", "europe-west3", "asia-east1" etc
region = "replace-with-region"

# GCP zone for deployment
# Example: "us-west1-a", "us-west1-b", "us-west1-c" etc
zone = "replace-with-zone"

# Replace with the ID of the existing Virtual Network
vnet_name = "replace-with-your-vnet-name"

# Replace with the ID of the existing Subnet
subnet_name = "replace-with-your-subnet-name"

# Replace with the ID of the Cohesity VM image that was shared to you
image_id = "replace-with-shared-image-id"

# Tags to apply to the resources
tags = [
  "replace-with-key1:replace-with-value1",
  "replace-with-key2:replace-with-value2"
]

# The prefix to use for naming all GCP resources (e.g., Virtual Machines,
# NICs, Boot Disks, and Data Disks).
# This prefix will be prepended to all resource names, followed by additional
# identifiers.
#
# Example: If you set resource_name_prefix = "cohesity", the resources will be
# named as follows:
# - Virtual Machine: cohesity-vm-0
# - NIC: cohesity-vm-0-nic-0
# - Boot Disk: cohesity-vm-0-boot-disk
# - SSD Data Disk: cohesity-vm-0-ssd-disk-0
# - HDD Data Disk: cohesity-vm-0-hdd-disk-0
#
# Restrictions:
# - Must only contain lowecase letters (a-z), numbers (0-9), and hyphens (-).
# - Must not exceed 63 characters.
# - Must start with a lowecase letter.
# - Must not end with a hyphen.
resource_name_prefix = "replace-with-prefix"

# Set the following to true to attach public IPs to all GCP VMs created.
#
# This can be set to true if ssh connectivity via private IP is not present.
attach_public_ip = false

################################################################################
# Config Variables
################################################################################

# Configuration ID to use for VM deployment. Choose a config id to deploy from
# configs.json.
config_id = "<replace-with-config-id>"

# The machine family type for the instance.
# Example: "n2d-standard", "n2-standard".
machine_family_type = "replace-with-machine-type"

# Optional: Specify a custom configuration to override the config_id
# Uncomment and fill the block below to use custom configuration
# custom_config = {
# "InstanceCPUCount" : "16",
# "SSDTierNumDisks" : 2,
# "SSDTierDiskSizeinGB" : 1023,
# "SSDTierDiskType" : "pd-balanced",
# "HDDTierNumDisks" : 2,
# "HDDTierDiskSizeinGB" : 1024,
# "HDDTierDiskType" : "pd-balanced"
# }

################################################################################
# Cluster Create Variables
################################################################################

# Set to true if a cluster should be created. sshpass must be installed and SSH
# access to the network where the VM is created from where terraforn is being
# run must be present.
create_cluster = true

cluster_name     = "replace-with-cluster-name"
dns_server_ips   = "replace-with-dns-server-ips"
ntp_servers      = "replace-with-ntp-servers"
domain_names     = "replace-with-domain-names"
apps_subnet      = "replace-with-apps-subnet"
apps_subnet_mask = "replace-with-apps-subnet-mask"
