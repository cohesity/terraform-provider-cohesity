################################################################################
# Provider Configuration
################################################################################

terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.0"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
  }
}

provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
  environment     = var.environment

  # Restricted subscriptions (e.g. Azure Gov) scope automation SPs to only the
  # actions this module needs, so azurerm's default "core" auto-registration
  # of ~35 unrelated resource providers 403s and aborts the apply; this module
  # only needs Microsoft.Resources/Microsoft.Storage, which are pre-registered.
  resource_provider_registrations = "none"

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

  # Primary blob endpoint of the storage account. The provider renders this with
  # the correct suffix for the active Azure environment (e.g.
  # blob.core.windows.net for public, blob.core.usgovcloudapi.net for
  # usgovernment), so the resulting VHD URI is valid in sovereign clouds too.
  # Value already ends with a trailing slash (e.g. https://<acct>.blob.core.windows.net/).
  primary_blob_endpoint = var.storage_account_name == "" ? azurerm_storage_account.storage[0].primary_blob_endpoint : data.azurerm_storage_account.existing_storage[0].primary_blob_endpoint

  vhd_filename = basename(var.local_vhd_path)

  # Concurrent Jenkins runs of the same build share this fixed storage
  # account/container/filename; a shared blob name lets one run's Put Blob
  # reset another's in-progress upload. Give the native (Gov) upload path a
  # per-run unique name instead, so concurrent runs can never collide - at the
  # cost of re-uploading every time and leaving one blob per run (needs
  # periodic cleanup; no reuse-across-runs optimization).
  native_vhd_blob_name = var.use_azurerm_native_upload ? "${uuid()}-${local.vhd_filename}" : local.vhd_filename

  # Upload command for the native (Gov/sovereign) path. azcopy copies bytes
  # verbatim, so the Dynamic VHD build artifact is first converted to a Fixed
  # VHD with qemu-img — Azure managed-disk import (create_option=Import) rejects
  # Dynamic VHDs. azcopy then uploads the Fixed VHD as a Page blob to the
  # environment-correct endpoint (from the storage account URL);
  # --overwrite=true keeps reruns idempotent.
  # azcopy and qemu-img are expected to be pre-installed on the runner.
  azcopy_upload_cmd = <<EOT
    set -e

    command -v azcopy >/dev/null 2>&1 || { echo "ERROR: azcopy not found on PATH; it must be pre-installed on the runner." >&2; exit 1; }
    command -v qemu-img >/dev/null 2>&1 || { echo "ERROR: qemu-img not found on PATH; it must be pre-installed on the runner." >&2; exit 1; }

    # Convert the Dynamic VHD build artifact to a Fixed VHD. force_size preserves
    # the MB-aligned virtual size Azure requires. The Fixed copy is full-size
    # (not sparse), so it is removed on exit via trap.
    SRC_VHD="${var.local_vhd_path}"
    FIXED_VHD="$${SRC_VHD%.vhd}-fixed.vhd"
    trap 'rm -f "$FIXED_VHD"' EXIT
    echo "Converting VHD to Fixed subformat: $SRC_VHD -> $FIXED_VHD"
    qemu-img convert -f vpc -O vpc -o subformat=fixed,force_size "$SRC_VHD" "$FIXED_VHD"

    export AZCOPY_AUTO_LOGIN_TYPE=SPN
    export AZCOPY_SPA_CLIENT_SECRET='${var.client_secret}'
    export AZCOPY_SPA_APPLICATION_ID='${var.client_id}'
    export AZCOPY_TENANT_ID='${var.tenant_id}'

    BLOB_URL="${local.primary_blob_endpoint}${local.container_name}/${local.native_vhd_blob_name}"

    echo "Starting VHD upload using azcopy..."
    echo "Uploading: $FIXED_VHD -> $BLOB_URL"

    azcopy copy \
      "$FIXED_VHD" \
      "$BLOB_URL" \
      --blob-type=PageBlob \
      --overwrite=true \
      --cap-mbps=0 \
      --log-level=INFO

    echo "VHD upload completed successfully!"
  EOT

  # Define the azure-vhd-utils command as a local variable
  # Upload without --overwrite so incomplete page blobs can be resumed.
  # If the blob is already complete, azure-vhd-utils asks for --overwrite; treat
  # that as success (ensure-if-complete) so re-runs skip a full re-upload.
  azure_vhd_utils_cmd = <<EOT

    set -u
    echo "Starting VHD upload using azure-vhd-utils..."
    echo "Uploading: ${var.local_vhd_path} to ${local.storage_account_name}/${local.container_name}/${local.vhd_filename}"

    LOG=$(mktemp)
    RCFILE="$${LOG}.rc"
    trap 'rm -f "$${LOG}" "$${RCFILE}"' EXIT

    run_upload() {
      overwrite_flag="$${1:-}"
      set +e
      {
        ${var.azure_vhd_utils_path} upload \
          --localvhdpath "${var.local_vhd_path}" \
          --stgaccountname "${local.storage_account_name}" \
          --stgaccountkey "${local.nonsensitive_storage_account_key}" \
          --containername "${local.container_name}" \
          --blobname "${local.vhd_filename}" \
          --parallelism ${var.vhd_upload_parallelism} \
          $${overwrite_flag} \
          2>&1
        echo $? > "$${RCFILE}"
      } | tee "$${LOG}"
      read rc < "$${RCFILE}"
      set -e
    }

    run_upload
    if [ "$${rc}" -eq 0 ]; then
      echo "VHD upload completed successfully!"
      exit 0
    fi

    if grep -q "use the --overwrite option" "$${LOG}"; then
      echo "VHD already present and complete in Azure; skipping re-upload."
      exit 0
    fi

    if grep -q "Last modified time of the VHD file" "$${LOG}" && \
       grep -q "does not match" "$${LOG}"; then
      echo "Local VHD mtime does not match Azure blob; retrying with --overwrite..."
      run_upload --overwrite
      if [ "$${rc}" -eq 0 ]; then
        echo "VHD upload completed successfully after overwrite retry!"
        exit 0
      fi
    fi

    if grep -q "Upload Incomplete: Some blocks of the VHD failed to upload" "$${LOG}"; then
      echo "Some VHD pages failed to upload; retrying with --overwrite..."
      run_upload --overwrite
      if [ "$${rc}" -eq 0 ]; then
        echo "VHD upload completed successfully after overwrite retry!"
        exit 0
      fi
    fi

    echo "Error: VHD upload failed!" >&2
    exit 1
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
# Used when use_azurerm_native_upload = false (default / commercial). The
# azure-vhd-utils binary hardcodes the commercial 'core.windows.net' storage
# suffix, so this path is not usable in sovereign clouds.

resource "null_resource" "upload_vhd" {
  count = var.use_azurerm_native_upload ? 0 : 1

  # Re-run only when the target blob name or source content changes. Moving the
  # same build to another local directory must not re-upload it.
  triggers = {
    vhd_filename = local.vhd_filename
    vhd_md5      = filemd5(var.local_vhd_path)
  }

  provisioner "local-exec" {
    command = local.azure_vhd_utils_cmd
  }

  depends_on = [
    azurerm_storage_account.storage[0],
    azurerm_storage_container.container[0]
  ]
}

################################################################################
# Upload VHD using azcopy (Gov / sovereign)
################################################################################
# Used when use_azurerm_native_upload = true. azcopy copies bytes verbatim and
# is NOT VHD-aware, so azcopy_upload_cmd first converts the Dynamic VHD build
# artifact to a Fixed VHD (qemu-img) — Azure managed-disk import rejects Dynamic
# VHDs with a "Dynamic VHD type" error. azcopy then uploads the Fixed VHD as a
# Page blob, resolving the correct blob endpoint for usgovernment/china/etc.
# from the storage account URL; --overwrite=true makes reruns idempotent.
resource "null_resource" "upload_vhd_azcopy" {
  count = var.use_azurerm_native_upload ? 1 : 0

  triggers = {
    blob_name = local.native_vhd_blob_name
  }

  provisioner "local-exec" {
    command = local.azcopy_upload_cmd
  }

  depends_on = [
    azurerm_storage_account.storage,
    azurerm_storage_container.container,
    data.azurerm_storage_account.existing_storage
  ]
}
