################################################################################
# Outputs
################################################################################

output "instance_info" {
  value = {
    for i in range(length(oci_core_instance.vm)) :
    oci_core_instance.vm[i].id => {
      private_ip = oci_core_instance.vm[i].private_ip
      public_ip  = var.attach_public_ip ? oci_core_instance.vm[i].public_ip : null
    }
  }
}
