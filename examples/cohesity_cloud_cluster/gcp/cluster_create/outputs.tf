###############################################################################
# Outputs
###############################################################################

output "instance_info" {
  value = {
    for i in range(length(google_compute_instance.vm)) :
    google_compute_instance.vm[i].id => {
      private_ip = google_compute_instance.vm[i].network_interface[0].network_ip
      public_ip  = var.attach_public_ip ? google_compute_address.public_ip[i].address : null
    }
  }
}
