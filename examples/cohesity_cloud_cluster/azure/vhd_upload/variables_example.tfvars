################################################################################
# Authentication Variables
################################################################################

# Supported Azure Authentication Methods:
#
# 1. Service Principal Authentication
#    - Set these four variables below. Leave others blank/false.
#    - How to create: https://learn.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal
#    - Where to find values:
#        subscription_id: Azure Portal > Subscriptions > Your Subscription > Subscription ID
#        client_id: Azure Portal > Azure Active Directory > App registrations > Your App > Application (client) ID
#        client_secret: Azure Portal > App registrations > Your App > Certificates & secrets > New client secret
#        tenant_id: Azure Portal > Azure Active Directory > Overview > Directory (tenant) ID
#    - Example:
#        subscription_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#        client_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#        client_secret = "your-app-secret"
#        tenant_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#
# 2. Managed Identity Authentication (for use within Azure VMs with Managed Identity enabled)
#    - Set use_managed_identity = true
#    - If using a user-assigned managed identity, set managed_identity_id to the
#      Client ID of the identity (Azure Portal > Managed Identities > Your Identity > Client ID)
#    - Leave all other authentication variables blank
#    - Example:
#        use_managed_identity = true
#        managed_identity_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" # (optional)
#
# 3. Azure CLI Authentication (for local development, uses 'az login' credentials)
#    - Set use_azure_cli = true
#    - All other authentication variables can be left blank
#    - Example:
#        use_azure_cli = true
#
# Only set variables for the authentication method you are using. Leave others blank or false.

# (Required for all authentication methods) Your Azure Subscription ID
subscription_id = "replace-with-your-subscription-id"

# (Service Principal only) Application (client) ID
client_id = "replace-with-your-client-id"
# (Service Principal only) Application secret
client_secret = "replace-with-your-client-secret"
# (Service Principal only) Directory (tenant) ID
tenant_id = "replace-with-your-tenant-id"

# (Managed Identity only) Set to true to use Managed Identity
use_managed_identity = false
# (Managed Identity only) User-assigned Managed Identity Client ID (optional, leave blank for system-assigned)
managed_identity_id = ""

# (Azure CLI only) Set to true to use Azure CLI authentication
use_azure_cli = false

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
