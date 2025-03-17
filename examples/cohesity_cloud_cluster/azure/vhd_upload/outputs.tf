output "storage_account_id" {
  value       = local.storage_account_id
  description = "The ID of the Storage Account where the VHD was uploaded"
}

output "vhd_uri" {
  value       = "https://${var.storage_account_name == "" ? azurerm_storage_account.storage[0].name : var.storage_account_name}.blob.core.windows.net/${local.container_name}/${local.vhd_filename}"
  description = "The URL of the uploaded VHD file"
}
