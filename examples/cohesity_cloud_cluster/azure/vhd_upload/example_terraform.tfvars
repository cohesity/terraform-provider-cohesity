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
# Resource and Upload Variables
################################################################################

# Azure region where the resources will be created.
# Example: "East US", "West US", "Central US", "West Europe", "eastus2", etc.
# See all regions: https://azure.microsoft.com/en-us/explore/global-infrastructure/products/
region = "replace-with-region"

# Prefix for naming resources (resource group, storage account, container) that
# this terraform creates.
# - Must be 3-20 characters long
# - Can only contain lowercase letters and numbers
# - Must start with a letter
# Example: If set to "myprefix", resources will be named:
#   Resource Group: myprefix-rg
#   Storage Account: myprefixsa
#   Storage Container: myprefix-container
resource_name_prefix = "replace-with-prefix"

# Name of the existing resource group to use (leave blank to create a new one)
# Example: "my-existing-rg"
resource_group_name = ""

# Name of the existing storage account to use (leave blank to create a new one)
# Example: "mystorageacct"
storage_account_name = ""

# Name of the existing container in the storage account (leave blank to create a new one)
# Example: "vhds"
container_name = ""

# Path to the local VHD file to upload
# Example: "/home/user/images/myimage.vhd"
local_vhd_path = "/path/to/your/vhd-file.vhd"

# Path to the azure-vhd-utils executable
# Example: "/usr/local/bin/azure-vhd-utils"
azure_vhd_utils_path = "/path/to/azure-vhd-utils"
