output "storage_account_id" {
  value       = local.storage_account_id
  description = "The ID of the Storage Account where the VHD was uploaded"
}

output "vhd_uri" {
  # local.primary_blob_endpoint is environment-aware (public, usgovernment, etc.)
  # and already includes the trailing slash, so it works in sovereign clouds.
  # For the native/Gov path, local.native_vhd_blob_name is unique per run;
  # for the azure-vhd-utils/commercial path, it equals the plain vhd_filename.
  value       = "${local.primary_blob_endpoint}${local.container_name}/${local.native_vhd_blob_name}"
  description = "The URL of the uploaded VHD file"
  # Referencing both upload resources here is safe even though only one has
  # count = 1 for a given use_azurerm_native_upload value: Terraform allows
  # depends_on to reference a 0-count resource, it just contributes no edges.
  depends_on = [null_resource.upload_vhd, null_resource.upload_vhd_azcopy]
}
