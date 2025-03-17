################################################################################
# Authentication Variables
################################################################################

# Region for deployment
# Example: "us-west-1"
region = "replace-with-your-region"

################################################################################
# Deployment Variables
################################################################################

# Number of Virtual Machines to create
# Replace with number VMs to create
num_instances = 3

# The prefix to use for naming all resources (e.g., Virtual Machines,
# OS Disks, and Data Disks).
# This prefix will be prepended to all resource names, followed by additional
# identifiers.
#
# Example: If you set resource_name_prefix = "cohesity", the resources will be
# named as follows:
# - Virtual Machine: cohesity-vm-0
# - OS Disk: cohesity-vm-0-os-disk
# - SSD Data Disk: cohesity-vm-0-ssd-disk-0
# - HDD Data Disk: cohesity-vm-0-hdd-disk-0
#
# Restrictions:
# - Must only contain letters (a-z, A-Z), numbers (0-9), and hyphens (-).
# - Must not exceed 30 characters.
# - Cannot start or end with a hyphen.
# - Hyphens cannot be consecutive (e.g., "--").
resource_name_prefix = "replace-with-prefix"

# Replace with the ID of the existing Subnet
# Example: "subnet-abc1234def5678"
subnet_id = "replace-with-subnet-id"

# Provide the ID of the Network Security Group. Optional.
# Example: ["sg-0a1b2c3d4e5f6g7h", "sg-abc1234def5678"]
security_group_ids = ["replace-with-security-group-id"]

# Replace with the AMI ID of the Cohesity VM image that was shared to you
# Example: "ami-abc1234def5678"
image_id = "replace-with-ami-id"

# Tags to apply to the resources
tags = [
  "replace-with-key1:replace-with-value1",
  "replace-with-key2:replace-with-value2"
]

# Set the following to true to attach public IPs to all VMs created.
#
# This can be set to true if ssh connectivity via private IP is not present.
attach_public_ip = false

################################################################################
# Config Variables
################################################################################

# Configuration ID to use for VM deployment. Choose a config id to deploy from
# configs.json.
config_id = "2"

# Optional: Specify a custom configuration to override the config_id
# Uncomment and fill the block below to use custom configuration
# custom_config = {
# "InstanceType" = "m6i.4xlarge",
# "SSDTierDiskType" = "gp3",
# "SSDTierNumDisks" = 2,
# "SSDTierDiskSizeinGB" = 1023,
# "SSDTierDiskIops" = 3000,
# "SSDTierDiskThroughputinMBps" = 200,
# "HDDTierDiskType" = "gp3",
# "HDDTierNumDisks" = 2,
# "HDDTierDiskSizeinGB" = 1024,
# "HDDTierDiskIops" = 3000,
# "HDDTierDiskThroughputinMBps" = 400
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
