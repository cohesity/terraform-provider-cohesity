################################################################################
# Authentication Variables
################################################################################

# Choose your authentication method:
#   "service_principal" - Use client ID and secret
#   "managed_identity"  - Use identity attached to Azure VM (for in-Azure runners eg Azure VM)
#   "azure_cli"         - Use 'az login' credentials (for local development)
auth_method = ""

# Azure Subscription ID (required for all methods)
# Find at: Azure Portal > Subscriptions > Your Subscription > Subscription ID
subscription_id = "replace-with-your-subscription-id"

################################################################################
# Authentication Credentials
################################################################################
# Fill in based on your chosen auth_method:
#
# SERVICE PRINCIPAL (all three required):
#   - client_id: Application (client) ID from App registrations
#   - client_secret: Client secret value
#   - tenant_id: Directory (tenant) ID
#   How to create: https://learn.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal
#
# MANAGED IDENTITY (usually leave blank):
#   - client_id: Only needed if VM has MULTIPLE user-assigned identities
#     Leave blank for system-assigned or single user-assigned identity (auto-detected)
#     Find at: Azure Portal > Managed Identities > Your Identity > Client ID
#   - client_secret: Not used (leave blank)
#   - tenant_id: Not used (leave blank)
#
# AZURE CLI (none needed):
#   - Leave all three blank

client_id     = ""     # See notes above
client_secret = ""     # Service Principal only
tenant_id     = ""     # Service Principal only

# Azure Environment (Valid values: public, usgovernment, china, german)
# Most users should leave this as "public"
environment = "public"

################################################################################
# Deployment Variables
################################################################################

# Number of Virtual Machines to create
num_instances = 3

# The prefix to use for naming all Azure resources (e.g., Virtual Machines,
# NICs, OS Disks, and Data Disks). This prefix will be prepended to all
# resource names, followed by additional identifiers.
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

# Name of existing resource group to use (leave blank to create a new one)
# - If blank, a new resource group named {resource_name_prefix}-rg will be created
resource_group_name = "replace-with-resource-group-name"

# Azure Region for deployment (e.g., eastus, westus2, westeurope)
# - List of regions: https://azure.microsoft.com/en-us/explore/global-infrastructure/products/
region = "replace-with-region"

# Specify the availability zone (optional). Example: "1", "2", or "3".
# Uncomment to use.
# availability_zone = "1"

# ID of the existing Virtual Network
# Example: /subscriptions/<subscription_id>/resourceGroups/<resource_group_name>/providers/Microsoft.Network/virtualNetworks/<vnet_name>
vnet_id = "replace-with-your-vnet-id"

# ID of the existing Subnet
# Example: /subscriptions/<subscription_id>/resourceGroups/<resource_group_name>/providers/Microsoft.Network/virtualNetworks/<vnet_name>/subnets/<subnet_name>
subnet_id = "replace-with-your-subnet-id"

# ID of the Network Security Group to attach to the NIC (optional)
# Example: /subscriptions/<subscription_id>/resourceGroups/<resource_group_name>/providers/Microsoft.Network/networkSecurityGroups/<nsg_name>
# Comment out if not needed
# network_security_group_id = "replace-with-your-nsg-id"

# Replace with the VHD URL of the Cohesity VM VHD uploaded to the storage account
vhd_uri = "replace-with-vhd-url"
# Replace with the storage account id where the VHD was uploaded
storage_account_id = "replace-with-storage-account-id"

# Tags to apply to the resources
tags = [
  "replace-with-key1:replace-with-value1",
  "replace-with-key2:replace-with-value2"
]

# Set to true to add a random 8-character string to all resource names (optional)
add_random_prefix = false

# Set to true to attach public IPs to all Azure VMs (set to true if you need SSH access from outside the VNet)
attach_public_ip = false

# (Optional) The Azure Disk Encryption Set (DES) resource ID to use for customer-managed disk encryption.
# If not set, Azure-managed keys are used.
# Example:
# customer_managed_disk_encryption_id = "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/my-rg/providers/Microsoft.Compute/diskEncryptionSets/my-des"
customer_managed_disk_encryption_id = ""

# List of user-assigned managed identities to attach to the Azure VMs (optional)
# Example: [
#   "/subscriptions/<subscription_id>/resourceGroups/<resource_group_name>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identity_name1>",
#   "/subscriptions/<subscription_id>/resourceGroups/<resource_group_name>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identity_name2>"
# ]
user_assigned_managed_identities_to_attach = []

################################################################################
# Config Variables
################################################################################

# Configuration ID to use for VM deployment (see configs.json for available IDs)
config_id = "3"

# Path to the JSON configuration file (default: ./configs.json)
config_file = "./configs.json"

# Optional: Specify a custom configuration to override the config_id
# Uncomment and fill the block below to use custom configuration
# custom_config = {
#   InstanceType                = "Standard_D16s_v5"
#   SSDTierNumDisks             = 2
#   SSDTierDiskSizeinGB         = 2047
#   SSDTierDiskType             = "Premium_LRS"
#   SSDTierDiskIops             = 0
#   SSDTierDiskThroughputinMBps = 0
#   HDDTierNumDisks             = 2
#   HDDTierDiskSizeinGB         = 2048
#   HDDTierDiskType             = "Premium_LRS"
#   HDDTierDiskIops             = 0
#   HDDTierDiskThroughputinMBps = 0
# }

###############################################################################
# Cluster Create Variables
###############################################################################

# The name of the Cohesity cluster. Specify a name up to 63 alphanumeric
# characters such as Cohesity1MyCompany. Hyphens are allowed but cannot be
# the first or last character. This is the name of the Cohesity cluster as
# displayed in the Cohesity UI and the Cohesity CLI.
cluster_name = "replace-with-cluster-name"

# The IP addresses of the Domain Name System (DNS) servers that the Cohesity
# cluster should use. Separate multiple IPs with commas.
dns_server_ips = "replace-with-dns-server-ips"

# Cohesity recommends using the external Google Public Network Time Protocol
# (NTP) server and specifying multiple servers (time1.google.com,
# time2.google.com, time3.google.com, time4.google.com).
ntp_servers = "replace-with-ntp-servers"

# The domain name(s) for the Cohesity cluster. (comma-separated if multiple)
domain_names = "replace-with-domain-names"

# A private IPv4 subnet for the internal overlay network of the kubernetes
# cluster to run the apps infrastructure.
# Update accordingly.
apps_subnet = "192.168.0.0"

# A private IPv4 subnet mask for the internal overlay network of the
# kubernetes cluster to run the apps infrastructure.
# Update accordingly.
apps_subnet_mask = "255.255.0.0"

# Whether to issue the cluster creation command via SSH after VM deployment.
# Usually set to true (default) as Terraform is expected to run in the same
# VPC/VNet as the VMs, or public IPs are attached for SSH access.
# Set to false if you want to create the cluster manually.
issue_cluster_create_cmd = true
