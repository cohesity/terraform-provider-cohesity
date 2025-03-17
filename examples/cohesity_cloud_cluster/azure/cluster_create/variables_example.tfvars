################################################################################
# Authentication Variables
################################################################################

# Required Azure Subscription ID
subscription_id = "your-subscription-id"

# (Optional) Service Principal Authentication - Leave empty if using another method.
client_id     = "your-client-id"
client_secret = "your-client-secret"
tenant_id     = "your-tenant-id"

# (Optional) Managed Identity Authentication. Leave as false if using another method.
use_managed_identity = false

# (Optional) Azure CLI Authentication. Leave as false if using another method.
use_azure_cli = false

# Azure Environment (Valid values: public, usgovernment, china, german)
environment = "public"

################################################################################
# Deployment Variables
################################################################################

# Number of Virtual Machines to create
num_instances = 1

# The prefix to use for naming all Azure resources (e.g., Virtual Machines,
# NICs, OS Disks, and Data Disks).
# This prefix will be prepended to all resource names, followed by additional
# identifiers.
#
# Example: If you set resource_name_prefix = "cohesity", the resources will be
# named as follows:
# - Resource group: cohesity-rg
# - Virtual Machine: cohesity-vm-0
# - NIC: cohesity-vm-0-nic-0
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

# Name of existing resource group where resources such as Virtual Machines, NICs, OS
# Disks, and Data Disks will be created (leave empty to create a new one).
# If left empty name of resource group created would be {resource_name_prefix}-rg
resource_group_name = "replace-with-resource-group-name"

# Azure Region for deployment
# Example: "East US", "West US", "Central US", "West Europe" etc
region = "replace-with-region"

# Specify the availability zone. Optional. Example: "1", "2", or "3".
# Uncomment if deployment is needed in an availability zone
# availability_zone = "1"

# Replace with the ID of the existing Virtual Network
# Example: "/subscriptions/<subscription_id>/resourceGroups/<resource_group_name>/providers/Microsoft.Network/virtualNetworks/<vnet_name>"
vnet_id = "replace-with-your-vnet-id"

# Replace with the ID of the existing Subnet
# Example: "/subscriptions/<subscription_id>/resourceGroups/<resource_group_name>/providers/Microsoft.Network/virtualNetworks/<vnet_name>/subnets/<subnet_name>"
subnet_id = "replace-with-your-subnet-id"

# Provide the ID of the Network Security Group. Optional.
# Comment out if no NSG is needed
# Example: "/subscriptions/<subscription_id>/resourceGroups/<resource_group_name>/providers/Microsoft.Network/networkSecurityGroups/<nsg_name>"
network_security_group_id = "replace-with-your-nsg-id"

# Replace with the VHD URL of the Cohesity VM VHD uploaded to the storage account
vhd_uri = "replace-with-vhd-url"
# Replace with the storage account id where the VHD was uploaded
storage_account_id = "replace-with-storage-account-id"

# Tags to apply to the resources
tags = [
  "replace-with-key1:replace-with-value1",
  "replace-with-key2:replace-with-value2"
]

# Set the following to true to attach public IPs to all Azure VMs created.
#
# This can be set to true if ssh connectivity via private IP is not present.
attach_public_ip = false

################################################################################
# Config Variables
################################################################################

# Configuration ID to use for VM deployment. Choose a config id to deploy from
# configs.json.
config_id = "3"

# Optional: Specify a custom configuration to override the config_id
# Uncomment and fill the block below to use custom configuration
# custom_config = {
#   InstanceType                = "Standard_D16s_v5",
#   SSDTierNumDisks             = 2,
#   SSDTierDiskSizeinGB         = 2047,
#   SSDTierDiskType             = "Premium_LRS",
#   SSDTierDiskIops             = 0,
#   SSDTierDiskThroughputinMBps = 0,
#   HDDTierNumDisks             = 2,
#   HDDTierDiskSizeinGB         = 2048,
#   HDDTierDiskType             = "Premium_LRS",
#   HDDTierDiskIops             = 0,
#   HDDTierDiskThroughputinMBps = 0
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
