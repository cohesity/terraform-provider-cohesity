# Provider dependencies.
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.40.0"
    }
    cohesity = {
      source  = "terraform-providers/cohesity"
      version = "0.1.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "~> 0.7" # Specify the version of the time provider
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2.2"
    }
    local = {
      source  = "hashicorp/local"
      version = "~> 2.5.1"
    }
  }
}

# Init the google provider.
provider "google" {
  credentials = file(var.credential_file)
  project     = var.project
  region      = var.region
}

provider "time" {
}

locals {
  instance_names           = [for i in range(var.num_instances) : format("%s-vm-%03d", var.cluster_name, i)]
  hostname                 = var.allot_external_ip ? google_compute_instance.nodes[0].network_interface[0].access_config[0].nat_ip : google_compute_instance.nodes[0].network_interface[0].network_ip
  private_ips              = join(",", [for i in google_compute_instance.nodes : i.network_interface[0].network_ip])
  public_ips               = join(",", [for i in google_compute_address.external_ip : i.address])
  metadata_fault_tolerance = var.num_instances > 1 ? 1 : 0
}
resource "google_compute_address" "external_ip" {
  count = var.allot_external_ip ? var.num_instances : 0
  name  = local.instance_names[count.index]
}

resource "google_compute_disk" "volume_ssd" {
  count = var.num_instances
  name  = format("%s-ssd", local.instance_names[count.index])
  size  = 1024
  type  = "pd-ssd"
  zone  = var.zone
}

resource "google_compute_disk" "volume_hdd" {
  count = var.num_instances
  name  = format("%s-hdd", local.instance_names[count.index])
  size  = 1024
  type  = "pd-standard"
  zone  = var.zone
}

resource "google_compute_instance" "nodes" {
  depends_on   = [google_compute_address.external_ip, google_compute_disk.volume_hdd, google_compute_disk.volume_ssd]
  count        = var.num_instances
  name         = local.instance_names[count.index]
  machine_type = var.machine_type
  zone         = var.zone


  boot_disk {
    initialize_params {
      image = var.disk_image
    }
  }
  attached_disk {
    source = google_compute_disk.volume_ssd[count.index].id
  }
  attached_disk {
    source = google_compute_disk.volume_hdd[count.index].id
  }

  network_interface {
    network    = var.network
    subnetwork = var.subnetwork
    access_config {
      nat_ip = google_compute_address.external_ip[count.index].address
    }
  }
  labels = var.labels
  tags   = ["ssh", "http"]

}
resource "time_sleep" "post_create_wait" {
  count           = 1
  depends_on      = [google_compute_instance.nodes]
  create_duration = "300s"
  triggers = {
    nodes_count = "${length(google_compute_instance.nodes.*.id)}"
  }
}
data "google_compute_subnetwork" "subnetwork" {
  name   = var.subnetwork
  region = var.region
}
locals {
  # Extract the CIDR prefix length (e.g., /24) from the CIDR notation (e.g., 10.10.0.0/24)
  cidr_prefix        = split("/", data.google_compute_subnetwork.subnetwork.ip_cidr_range)[1]
  subnet_mask        = trimspace(data.local_file.subnet_mask.content)
  subnetwork_gateway = cidrhost(data.google_compute_subnetwork.subnetwork.ip_cidr_range, 1)

}
resource "null_resource" "cidr_to_mask" {
  depends_on = [data.google_compute_subnetwork.subnetwork]
  provisioner "local-exec" {
    command = "./cidr_to_subnet_mask.sh ${local.cidr_prefix} ./subnet_mask.txt"
  }

  # This triggers the null_resource when the CIDR changes
  # triggers = {
  #   cidr = data.google_compute_subnetwork.subnetwork.ip_cidr_range
  # }
}
# Read the subnet mask from the file
data "local_file" "subnet_mask" {
  depends_on = [null_resource.cidr_to_mask]
  filename   = "${path.module}/subnet_mask.txt"
}

provider "cohesity" {
  cluster_vip      = var.allot_external_ip ? google_compute_instance.nodes[0].network_interface[0].access_config[0].nat_ip : google_compute_instance.nodes[0].network_interface[0].network_ip
  cluster_username = var.cohesity_username
  cluster_password = var.cohesity_password
  support_password = var.support_password
  cluster_domain   = "LOCAL"
}

resource "cohesity_ngce_cluster" "cluster" {
  depends_on = [time_sleep.post_create_wait]

  name                     = var.cluster_name
  node_ips                 = [for instance in google_compute_instance.nodes : instance.network_interface[0].network_ip]
  hostname                 = local.hostname
  subnet_gateway           = local.subnetwork_gateway
  subnet_mask              = local.subnet_mask
  dns_server_ips           = var.dns-server-ips
  ntp_servers              = var.ntp-servers
  domain_names             = var.domain-names
  metadata_fault_tolerance = local.metadata_fault_tolerance
  add_apps                 = var.add-apps
  apps_subnet              = var.apps-subnet
  apps_subnet_mask         = var.apps-subnet-mask
}

resource "cohesity_gcp_external_target" "gcp_external_target" {
  depends_on                   = [cohesity_cluster.cluster]
  client_private_key_file_path = var.gcp_external_target.client_private_key_file_path
  bucket_name                  = var.gcp_external_target.bucket_name
  project_id                   = var.gcp_external_target.project_id
  client_email                 = var.gcp_external_target.client_email_id
  tier_type                    = var.gcp_external_target.tier_type
}

output "instance_ids" {
  description = "A comma-separated list of all the instance ids of the GCP nodes"
  value       = join(",", [for instance in google_compute_instance.nodes : instance.instance_id])
}

output "private_ips" {
  description = "A comma-separated list of all private IP addresses of the GCP nodes"
  value       = local.private_ips
}

output "public_ips" {
  description = "A comma-separated list of all public IP addresses of the GCP nodes"
  value       = local.public_ips
}