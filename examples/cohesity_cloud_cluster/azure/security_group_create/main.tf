################################################################################
# Provider Configuration
################################################################################

terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.0"
    }
  }
}

provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
  environment     = var.environment

  # Credentials are interpreted based on auth_method:
  # - Service Principal: client_id + client_secret + tenant_id all required
  # - Managed Identity: client_id optional (only required for multiple user-assigned identities)
  # - Azure CLI: all empty, provider uses 'az login' credentials automatically
  client_id     = var.client_id != "" ? var.client_id : null
  client_secret = var.client_secret != "" ? var.client_secret : null
  tenant_id     = var.tenant_id != "" ? var.tenant_id : null

  # Enable managed identity authentication when specified
  use_msi = var.auth_method == "managed_identity"
}

locals {
  rules               = jsondecode(file(var.rules_file)).rules
  resource_group_name = var.resource_group_name == "" ? azurerm_resource_group.resource_group[0].name : var.resource_group_name

  # Parse user-provided tags into a key-value map
  parsed_tags = { for tag in var.tags : split(":", tag)[0] => split(":", tag)[1] }
}

###############################################################################
# Resource Group (Created if not provided)
###############################################################################

resource "azurerm_resource_group" "resource_group" {
  count    = var.resource_group_name == "" ? 1 : 0
  name     = "${var.resource_name_prefix}-rg"
  location = var.region
}

###############################################################################
# Network Security Group (NSG)
###############################################################################

resource "azurerm_network_security_group" "cohesity_nsg" {
  name                = "${var.resource_name_prefix}-nsg"
  location            = var.region
  resource_group_name = local.resource_group_name
  tags                = local.parsed_tags
}

resource "azurerm_network_security_rule" "cohesity_rule" {
  count                       = length(local.rules)
  name                        = local.rules[count.index].name
  priority                    = local.rules[count.index].priority
  direction                   = local.rules[count.index].direction
  access                      = local.rules[count.index].access
  protocol                    = local.rules[count.index].protocol
  source_port_range           = local.rules[count.index].source_port_range
  destination_port_range      = local.rules[count.index].destination_port_range
  source_address_prefix       = local.rules[count.index].direction == "Outbound" ? var.cidr : lookup(local.rules[count.index], "source_address_prefix", "*")
  destination_address_prefix  = local.rules[count.index].direction == "Inbound" ? var.cidr : lookup(local.rules[count.index], "destination_address_prefix", "*")
  resource_group_name         = local.resource_group_name
  network_security_group_name = azurerm_network_security_group.cohesity_nsg.name
  description                 = lookup(local.rules[count.index], "description", null)
}
