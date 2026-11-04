###############################################################################
# Outputs
###############################################################################

# Output the instance IDs and private IPs

output "instance_info" {
  value = {
    for i in range(length(azurerm_virtual_machine.azure_vm)) :
    azurerm_virtual_machine.azure_vm[i].id => {
      private_ip = azurerm_network_interface.nic[i].private_ip_address
      public_ip  = var.attach_public_ip ? azurerm_public_ip.public_ip[i].ip_address : null
    }
  }
}

output "system_managed_identity_principal_ids" {
  description = "Principal (object) IDs of the system-assigned managed identities, one per VM. Empty when SMI is disabled."
  value = var.enable_system_managed_identity ? [
    for vm in azurerm_virtual_machine.azure_vm : vm.identity[0].principal_id
  ] : []
}

output "system_managed_identity_role_assignments" {
  description = "Map of SMI role assignment IDs created (key: vm<N>-role<M>). Empty when SMI is disabled."
  value = var.enable_system_managed_identity ? {
    for k, ra in azurerm_role_assignment.smi_roles : k => {
      role_definition_name = ra.role_definition_name
      scope                = ra.scope
      principal_id         = ra.principal_id
    }
  } : {}
}
