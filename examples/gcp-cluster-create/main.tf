provider "google" {
  credentials = file(var.credential_file)
  project     = var.project
  region      = var.region
}
provider "cohesity" {
  ssh {
    ssh_host        = google_compute_instance.control_vm.network_interface[0].access_config[0].nat_ip
    ssh_user        = var.ssh_username
    ssh_password    = var.ssh_password
    ssh_timeout     = "60m"
  }
}
# Create the Control VM
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
  labels = var.labels
  tags = ["ssh", "http"]
  provisioner "local-exec" {
    command = "sleep 20"
  }

}
# Use the Control VM to deploy the cluster
resource "cohesity_gcp_cluster" "cvm_commands" {
  depends_on = [ google_compute_instance.control_vm ]
  cluster_config_file = "gcp-resources/cluster_config.json"
  gcp_key_file = "gcp-resources/GCPKey.json"
  gcp_bucket_config_file = "gcp-resources/gcp_bucket_config"
}
locals {
  cluster_config = jsondecode(file("gcp-resources/cluster_config.json"))
  cluster_name = local.cluster_config.cohesity_cluster_name
  cluster_gcp_zone = local.cluster_config.gcp_zone
  num_nodes = local.cluster_config.gcp_num_vms
}
# Get information of all the nodes on the cluster
data "google_compute_instance" "cluster" {
  count = num_nodes
  depends_on = [ cohesity_gcp_cluster.cvm_commands ]
  name = format("%s-vm-%s",local.cluster_name,count.index+1)
  zone = local.cluster_gcp_zone
}
output "cluster_private_ips" {
  description = "A comma-separated list of all private IP addresses of the GCP nodes"
  value = join(",",[for i in data.google_compute_instance.cluster : i.network_interface[0].network_ip])
}
output "instance_ids" {
  description = "A comma-separated list of all the instance ids of the GCP nodes"
  value = join(",", [for instance in data.google_compute_instance.cluster : instance.instance_id])
}
