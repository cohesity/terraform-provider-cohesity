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
