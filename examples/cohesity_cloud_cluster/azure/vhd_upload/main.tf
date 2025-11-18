################################################################################
# Provider Configuration
################################################################################

provider "azurerm" {
  features {}

  # Required for all authentication methods
  subscription_id = var.subscription_id

  # Service Principal authentication (if credentials are provided)
  client_id     = var.client_id != "" ? var.client_id : null
  client_secret = var.client_secret != "" ? var.client_secret : null
  tenant_id     = var.tenant_id != "" ? var.tenant_id : null

  # Use Managed Identity if enabled
  use_msi = var.use_managed_identity

  # Use Azure CLI authentication if enabled
  use_cli = var.use_azure_cli

  # Set the Azure cloud environment
  environment = var.environment
}

################################################################################
# Resource Group (Created if not provided)
################################################################################

resource "azurerm_resource_group" "resource_group" {
  count    = var.resource_group_name == "" ? 1 : 0
  name     = "${var.resource_name_prefix}-rg"
  location = var.region
}

################################################################################
# Storage Account (Existing or New)
################################################################################
# Creates a new storage account if the user does not provide one.
resource "azurerm_storage_account" "storage" {
  count                    = var.storage_account_name == "" ? 1 : 0
  name                     = "${var.resource_name_prefix}sa"
  resource_group_name      = local.resource_group_name
  location                 = var.region
  account_tier             = "Premium"
  account_replication_type = "LRS"
}

# Fetch details of an existing storage account if provided.
data "azurerm_storage_account" "existing_storage" {
  count               = var.storage_account_name != "" ? 1 : 0
  name                = var.storage_account_name
  resource_group_name = local.resource_group_name
}

################################################################################
# Storage Container (Existing or New)
################################################################################
# Creates a new container if the user does not provide one.
resource "azurerm_storage_container" "container" {
  count              = var.container_name == "" ? 1 : 0
  name               = "${var.resource_name_prefix}-container"
  storage_account_id = local.storage_account_id
}

################################################################################
# Local Variables
################################################################################

locals {
  resource_group_name  = var.resource_group_name == "" ? azurerm_resource_group.resource_group[0].name : var.resource_group_name
  storage_account_name = var.storage_account_name == "" ? azurerm_storage_account.storage[0].name : var.storage_account_name
  container_name       = var.container_name == "" ? azurerm_storage_container.container[0].name : var.container_name

  storage_account_key = var.storage_account_name == "" ? azurerm_storage_account.storage[0].primary_access_key : data.azurerm_storage_account.existing_storage[0].primary_access_key
  # Remove sensitivity if non sensitive logging is enabled
  nonsensitive_storage_account_key = var.log_storage_account_key ? nonsensitive(local.storage_account_key) : local.storage_account_key
  storage_account_id               = var.storage_account_name == "" ? azurerm_storage_account.storage[0].id : data.azurerm_storage_account.existing_storage[0].id

  vhd_filename = basename(var.local_vhd_path)

  # Define the azure-vhd-utils command as a local variable
  azure_vhd_utils_cmd = <<EOT
    echo "Starting VHD upload using azure-vhd-utils..."
    echo "Uploading: ${var.local_vhd_path} to ${local.storage_account_name}/${local.container_name}/${local.vhd_filename}"

    ${var.azure_vhd_utils_path} upload \
      --localvhdpath "${var.local_vhd_path}" \
      --stgaccountname "${local.storage_account_name}" \
      --stgaccountkey "${local.nonsensitive_storage_account_key}" \
      --containername "${local.container_name}" \
      --blobname "${local.vhd_filename}" \
      --parallelism ${var.vhd_upload_parallelism} \
      --overwrite

    if [ $? -eq 0 ]; then
      echo "VHD upload completed successfully!"
    else
      echo "Error: VHD upload failed!" >&2
      exit 1
    fi
  EOT
}

################################################################################
# Validation logic for resource dependencies
################################################################################

locals {
  validate_storage_and_rg = (
    var.storage_account_name != "" && var.resource_group_name == ""
  ) ? "ERROR: If storage_account_name is provided, resource_group_name must also be provided." : ""

  validate_container_and_storage = (
    var.container_name != "" && var.storage_account_name == ""
  ) ? "ERROR: If container_name is provided, storage_account_name must also be provided." : ""

  validate_container_and_rg = (
    var.container_name != "" && var.resource_group_name == ""
  ) ? "ERROR: If container_name is provided, resource_group_name must also be provided." : ""
}

# Enforce validation using a Terraform trick (null_resource with count = 0)
resource "null_resource" "validation_checks" {
  count = (
    local.validate_storage_and_rg != "" ||
    local.validate_container_and_storage != "" ||
    local.validate_container_and_rg != ""
  ) ? 1 : 0

  provisioner "local-exec" {
    command = <<EOT
      echo "${local.validate_storage_and_rg}"
      echo "${local.validate_container_and_storage}"
      echo "${local.validate_container_and_rg}"
      exit 1
    EOT
  }
}

################################################################################
# Upload VHD using azure-vhd-utils
################################################################################

resource "null_resource" "upload_vhd" {
  provisioner "local-exec" {
    command = local.azure_vhd_utils_cmd
  }

  depends_on = [
    azurerm_storage_account.storage[0],
    azurerm_storage_container.container[0]
  ]
}
