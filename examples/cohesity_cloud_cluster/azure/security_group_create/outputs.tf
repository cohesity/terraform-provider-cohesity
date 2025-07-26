################################################################################
# Outputs
################################################################################

output "network_security_group_id" {
  description = "The ID of the created network security group."
  value       = azurerm_network_security_group.cohesity_nsg.id
}
