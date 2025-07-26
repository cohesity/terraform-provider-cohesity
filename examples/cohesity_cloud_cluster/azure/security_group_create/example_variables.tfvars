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
# Security Group Variables
################################################################################

# The prefix to use for naming the security group and related resources
# This will be prepended to resource names, followed by additional identifiers
#
# Example: If you set resource_name_prefix = "cohesity", the security group will be named:
# - "cohesity-nsg"
#
# Restrictions:
# - Must only contain letters (a-z, A-Z), numbers (0-9), and hyphens (-).
# - Cannot start or end with a hyphen.
# - Hyphens cannot be consecutive (e.g., "--").
resource_name_prefix = "replace-with-prefix"

# Name of the resource group where the security group will be created
# - If blank, a new resource group named {resource_name_prefix}-rg will be created
resource_group_name = "replace-with-resource-group-name"

# Azure Region for deployment (e.g., eastus, westus2, westeurope)
# - List of regions: https://azure.microsoft.com/en-us/explore/global-infrastructure/products/
region = "replace-with-region"

# Tags to apply to the security group and related resources
# Format: key-value pairs in a map
# Example: { environment = "production", project = "cohesity-deployment" }
tags = {
  environment = "replace-with-environment",
  owner       = "replace-with-owner"
}

# CIDR block that will be used as source_address_prefix for Outbound rules and
# destination_address_prefix for Inbound rules
# - This should match the subnet CIDR where your Cohesity cluster will be deployed
# - Example: "10.0.0.0/24"
cidr = "replace-with-cidr-block"

# Path to the JSON file containing security group rules
# - The default file (./rules.json) contains a comprehensive set of rules for Cohesity deployments
# - You can customize this file or provide a different file with your own rule set
# - Format is a JSON array of rule objects as shown in the default rules.json
rules_file = "./rules.json"
