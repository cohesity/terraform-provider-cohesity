provider "google" {
  credentials = file(var.credential_file)
  project     = var.project
  region      = var.region
}
provider "cohesity" {
  type            = "ssh"
  ssh_host        = google_compute_instance.control_vm.network_interface[0].access_config[0].nat_ip
  ssh_user        = var.ssh_username
  ssh_password    = var.ssh_password
  ssh_timeout     = "60m"
}
resource "google_compute_instance" "control_vm" {

  name         = var.instance_name
  machine_type = var.machine_type
  zone         = var.zone

  boot_disk {
    initialize_params {
      image = var.disk_image
    }
  }

  network_interface {
    network    = var.network
    subnetwork = var.subnetwork
    access_config {
    }
  }
  labels = merge(var.labels, { expiry = local.expiry_date })
  tags = ["ssh", "http"]
  provisioner "local-exec" {
    command = "sleep 20"
  }

}

resource "cohesity_gcp_cluster" "cvm_commands" {
  depends_on = [ google_compute_instance.control_vm ]
  cluster_config_file = "gcp-resources/cluster_config.json"
  gcp_key_file = "gcp-resources/GCPKey.json"
  gcp_bucket_config_file = "gcp-resources/gcp_bucket_config"
}
locals {
  cluster_config = jsondecode(file("gcp-resources/cluster_config.json"))
  vm_name = local.cluster_config.cohesity_cluster_name
  expiry_date = formatdate("MM-DD-YYYY", timeadd(timestamp(),"336h"))
}
data "google_compute_instance" "cluster" {
  depends_on = [ cohesity_gcp_cluster.cvm_commands ]
  name = format("%s%s",local.vm_name,"-vm-1")
  zone = "us-west1-a"
}
output "cluster_ip" {
  value = data.google_compute_instance.cluster.network_interface[0].network_ip
}
