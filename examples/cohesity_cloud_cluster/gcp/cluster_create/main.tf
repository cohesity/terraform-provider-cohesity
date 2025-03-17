################################################################################
# Provider Configuration
################################################################################

provider "google" {
  project     = var.project_id
  region      = var.region
  credentials = file(var.credentials_file_path)
}

################################################################################
# Data Sources
################################################################################

# Fetch details of an existing subnet
data "google_compute_subnetwork" "network_details" {
  name    = var.subnet_name
  region  = var.region
  project = var.project_id
}

# Fetch details of an existing vnet
data "google_compute_network" "vnet_details" {
  name    = var.vnet_name
  project = var.project_id
}

# Fetch available zones in the given region
data "google_compute_zones" "available_zones" {
  region  = var.region
  project = var.project_id
}

################################################################################
# Random String for Resource Name Prefix
################################################################################

# Generate an 8-character alphanumeric random string for resource names
resource "random_string" "resource_name_prefix" {
  length  = 8
  special = false
  upper   = false
}

################################################################################
# Local Variables
################################################################################

locals {
  random_prefix = random_string.resource_name_prefix.result

  # Define resource name prefix based on user input
  resource_name_prefix = var.add_random_prefix == false ? var.resource_name_prefix : "${var.resource_name_prefix}-${local.random_prefix}"

  # Extract Address Prefix and Subnet Gateway
  # Address Prefix example "10.100.128.0/21"
  address_prefix = data.google_compute_subnetwork.network_details.ip_cidr_range
  # Extract subnet mask
  subnet_mask = cidrnetmask(local.address_prefix)
  # First usable IP is the subnet gateway
  subnet_gateway = cidrhost(local.address_prefix, 1)

  # Parse user-provided tags into a key-value map
  parsed_tags = { for tag in var.tags : split(":", tag)[0] => split(":", tag)[1] }

  # Collect private IPs of the created NICs
  private_ips = [
    for instance in google_compute_instance.vm :
    instance.network_interface[0].network_ip
  ]

  # Collect public IPs of the created public IPs (if any)
  public_ips = var.attach_public_ip ? google_compute_address.public_ip.*.address : []

  # Set instance_ip based on attach_public_ip flag
  # This is the IP to which cluster create command is issued via ssh.
  instance_ip = var.attach_public_ip ? local.public_ips[0] : local.private_ips[0]

  # Convert the list of private IPs into a comma-separated string
  private_ips_string = join(",", local.private_ips)

  # Convert the list of public IPs into a comma-separated string (if applicable)
  public_ips_string = join(",", local.public_ips)

  # Load VM configurations from the specified JSON file
  configurations = jsondecode(file(var.config_file)).Configurations

  # Select the configuration using the `config_id` variable
  selected_config = lookup(
    { for config in local.configurations : config.Id => config },
    var.config_id,
    null
  )

  # Fallback to `custom_config` if provided; otherwise, use the selected configuration
  config = var.custom_config != null ? var.custom_config : local.selected_config

  # Construct the machine type by combining machine_family_type and InstanceCPUCount
  machine_type = "${var.machine_family_type}-${local.config.InstanceCPUCount}"

  # Determine fault tolerance level based on the number of instances
  metadata_fault_tolerance = var.num_instances > 1 ? 1 : 0

  # Construct iris_cli command prefixes
  iris_cli_cmd_prefix = join(" ", [
    "iris_cli",
    "--v=10",
    "--username=admin",
    "--password=admin",
    "--skip_password_prompt=true",
    "--skip_force_password_change=true"
  ])

  # Construct cluster command
  cluster_create_cmd = join(" ", [
    local.iris_cli_cmd_prefix,
    "cluster cloud-create cluster-size=nextGen enable-software-encryption=true",
    "name=${var.cluster_name}",
    "node-ips=${local.private_ips_string}",
    "hostname=${local.private_ips[0]}",
    "subnet-gateway=${local.subnet_gateway}",
    "subnet-mask=${local.subnet_mask}",
    "dns-server-ips=${var.dns_server_ips}",
    "ntp-servers=${var.ntp_servers}",
    "domain-names=${var.domain_names}",
    "metadata-fault-tolerance=${local.metadata_fault_tolerance}",
    "apps-subnet=${var.apps_subnet}",
    "apps-subnet-mask=${var.apps_subnet_mask}"
  ])

  # Construct cluster status command
  cluster_status_cmd = join(" ", [
    local.iris_cli_cmd_prefix,
    "cluster status"
  ])

  # Construct the SSH-based cluster creation command
  sshpass_cmd_prefix         = "sshpass -p 'Cohe$1ty' ssh -o StrictHostKeyChecking=no cohesity@${local.instance_ip}"
  sshpass_cluster_create_cmd = "${local.sshpass_cmd_prefix} ${local.cluster_create_cmd}"
  sshpass_cluster_status_cmd = "${local.sshpass_cmd_prefix} ${local.cluster_status_cmd}"
}

################################################################################
# Resource Definitions
################################################################################

# Create Public IPs for VMs if enabled
resource "google_compute_address" "public_ip" {
  count  = var.attach_public_ip ? var.num_instances : 0
  name   = "${local.resource_name_prefix}-vm-${count.index}-public-ip"
  region = var.region
  labels = local.parsed_tags
}

# Create Boot Disks
resource "google_compute_disk" "boot_disk" {
  count  = var.num_instances
  name   = "${local.resource_name_prefix}-vm-${count.index}-boot-disk"
  type   = "pd-standard"
  zone   = var.zone
  labels = local.parsed_tags
  image  = var.image_id
}

# Create SSD Tier Data Disks
resource "google_compute_disk" "ssd_tier_data_disk" {
  count  = var.num_instances * local.config.SSDTierNumDisks
  name   = "${local.resource_name_prefix}-vm-${floor(count.index / local.config.SSDTierNumDisks)}-ssd-tier-${count.index % local.config.SSDTierNumDisks}"
  type   = local.config.SSDTierDiskType
  zone   = var.zone
  size   = local.config.SSDTierDiskSizeinGB
  labels = local.parsed_tags
}

# Create HDD Tier Data Disks
resource "google_compute_disk" "hdd_tier_data_disk" {
  count  = var.num_instances * local.config.HDDTierNumDisks
  name   = "${local.resource_name_prefix}-vm-${floor(count.index / local.config.HDDTierNumDisks)}-hdd-tier-${count.index % local.config.HDDTierNumDisks}"
  type   = local.config.HDDTierDiskType
  zone   = var.zone
  size   = local.config.HDDTierDiskSizeinGB
  labels = local.parsed_tags
}

# Create Virtual Machines
resource "google_compute_instance" "vm" {
  count        = var.num_instances
  name         = "${local.resource_name_prefix}-vm-${count.index}"
  machine_type = local.machine_type
  zone         = var.zone
  labels       = local.parsed_tags
  # allow_stopping_for_update = true   (In case machinetype/numberofcpus needs to be changed, set this to true)

  boot_disk {
    source = google_compute_disk.boot_disk[count.index].self_link
  }

  network_interface {
    network    = var.vnet_name
    subnetwork = var.subnet_name
    dynamic "access_config" {
      for_each = var.attach_public_ip ? [1] : []
      content {
        # Ephemeral public IP if enabled
        nat_ip = var.attach_public_ip ? google_compute_address.public_ip[count.index].address : null
      }
    }
  }

  # Attach SSD tier disks
  dynamic "attached_disk" {
    for_each = [for i in range(local.config.SSDTierNumDisks) : google_compute_disk.ssd_tier_data_disk[count.index * local.config.SSDTierNumDisks + i]]
    content {
      source = attached_disk.value.self_link
    }
  }

  # Attach HDD tier disks
  dynamic "attached_disk" {
    for_each = [for i in range(local.config.HDDTierNumDisks) : google_compute_disk.hdd_tier_data_disk[count.index * local.config.HDDTierNumDisks + i]]
    content {
      source = attached_disk.value.self_link
    }
  }
}

################################################################################
# Commands Based on Cluster Creation
################################################################################

# Print command for manual cluster creation
resource "null_resource" "print_command" {
  depends_on = [google_compute_instance.vm]
  provisioner "local-exec" {
    command = "echo Run the following command on any one of the VMs created: ${local.cluster_create_cmd}"
  }
}

# Post boot commands for cluster creation
resource "null_resource" "post_boot_commands" {
  depends_on = [google_compute_instance.vm]
  count      = var.create_cluster ? 1 : 0

  provisioner "local-exec" {
    command = <<EOT
    echo "Waiting for 10 minutes for VM to be ready..."
    sleep ${var.post_boot_wait}  # Wait for 10 minutes before starting the loop

    while true; do
      output=$(${local.sshpass_cluster_status_cmd})

      echo "Command Output: $output"

      if echo "$output" | grep -q "is not part of a cluster"; then
        echo "VM services are ready"
        break
      fi

      echo "VM services are not ready yet. Retrying in 10 seconds..."
      sleep 10
    done
    EOT
  }
}

# Execute cluster create command
resource "null_resource" "execute_cluster_create_command" {
  depends_on = [null_resource.post_boot_commands]
  count      = var.create_cluster ? 1 : 0

  provisioner "local-exec" {
    command = local.sshpass_cluster_create_cmd
  }
}
