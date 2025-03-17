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
# Others
################################################################################

# Azure region where the resources will be created.
# Example: "East US", "West US", "Central US", "West Europe" etc
region = "West US"

# Prefix for naming resources (resource group, storage account, container) that
# this terraform creates.
# This only applies when resources are created
# - Must be 3-20 characters long
# - Can only contain lowercase letters and numbers
# - Must start with a letter
# Example: If `resource_name_prefix` is set to "myprefix", Terraform may create:
# - Resource Group:         myprefix-rg
# - Storage Account:        myprefixssa
# - Storage Container:      myprefix-container
resource_name_prefix = "myprefix"

# Name of the existing resource group
# Leave empty to create a new one
resource_group_name = ""

# Name of the existing storage account
# Leave empty to create a new one
storage_account_name = ""

# Name of the existing container in the storage account
# Leave empty to create a new one
container_name = ""

# Path to the local VHD file to upload
local_vhd_path = "/path/to/your/vhd-file.vhd"

# Path to the azure-vhd-utils executable
azure_vhd_utils_path = "/path/to/azure-vhd-utils"
