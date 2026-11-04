###############################################################################
# Provider Configuration
###############################################################################

terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 7.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.5"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
  }
}

provider "google" {
  project     = var.project_id
  region      = var.region
  credentials = var.credentials_file_path != "" ? file(var.credentials_file_path) : null
  impersonate_service_account = var.impersonate_service_account != "" ? var.impersonate_service_account : null

  # Sovereign / dedicated cloud universe.
  # Null for commercial (defaults to "googleapis.com").
  universe_domain = var.universe_domain != "" ? var.universe_domain : null
}

###############################################################################
# Data Sources
###############################################################################

# Fetch details of an existing subnet.
# In a Shared VPC deployment the subnet lives in the host project, so it is
# looked up in host_project_id when set (falls back to project_id otherwise).
data "google_compute_subnetwork" "network_details" {
  name    = var.subnet_name
  region  = var.region
  project = local.network_project
}

# Fetch available zones in the given region
data "google_compute_zones" "available_zones" {
  region  = var.region
  project = var.project_id
}

###############################################################################
# Random String for Resource Name Prefix
###############################################################################

# Generate an 8-character alphanumeric random string for resource names
resource "random_string" "random_prefix" {
  length  = 8
  special = false
  upper   = false
}

###############################################################################
# Local Variables
###############################################################################

locals {
  random_prefix = random_string.random_prefix.result

  # Project that owns the VPC network and subnet. For a Shared VPC deployment
  # this is the host project; otherwise it is the service project (project_id).
  network_project = var.host_project_id != "" ? var.host_project_id : var.project_id

  # Define resource name prefix based on user input
  resource_name_prefix = var.add_random_prefix == false ? var.resource_name_prefix : "${var.resource_name_prefix}-${local.random_prefix}"

  # Extract Address Prefix and Subnet Gateway
  # Address Prefix example "10.100.128.0/21"
  address_prefix = data.google_compute_subnetwork.network_details.ip_cidr_range
  # Extract subnet mask
  subnet_mask = cidrnetmask(local.address_prefix)
  # First usable IP is the subnet gateway
  subnet_gateway = cidrhost(local.address_prefix, 1)

  # Parse user-provided labels into a key-value map
  parsed_labels = merge(
    { for label in var.labels : split(":", label)[0] => split(":", label)[1] },
    {
      "goog-partner-solution1" = var.goog_partner_solution1
      "goog-partner-solution2" = var.goog_partner_solution2
    }
  )

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

  # Per-node device names assigned during disk attachment
  ssd_device_names = [for i in range(local.config.SSDTierNumDisks) : "ssd-${i}"]
  hdd_device_names = [for i in range(local.config.HDDTierNumDisks) : "hdd-${i}"]

  # Build the per-node disk-identifier -> tier mapping passed to iris_cli
  per_node_disk_identifiers = join(",", concat(
    local.ssd_device_names,
    local.hdd_device_names
  ))

  per_node_disk_tiers = join(",", concat(
    [for _ in local.ssd_device_names : "PCIeSSD"],
    [for _ in local.hdd_device_names : "SATA-HDD"]
  ))

  # Construct iris_cli command prefixes
  iris_cli_cmd_prefix = join(" ", [
    "iris_cli",
    "--v=10",
    "--username=admin",
    "--password=admin",
    "--skip_password_prompt=true",
    "--skip_force_password_change=true"
  ])

  # Construct cluster create command
  cluster_create_cmd = join(" ", [
    local.iris_cli_cmd_prefix,
    "cluster cloud-create cluster-size=${var.cluster_size} enable-software-encryption=true",
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
    "apps-subnet-mask=${var.apps_subnet_mask}",
  ])
  cluster_create_cmd_with_disk_mapping = join(" ", [
    local.cluster_create_cmd,
    "disk-identifiers=${local.per_node_disk_identifiers}",
    "disk-tiers=${local.per_node_disk_tiers}",
  ])

  # Nodes created while num_instances has this value are labeled with it.
  # The label is ignored on update, so earlier nodes keep their own value.
  # Nodes created before this label existed have none and count as existing.
  deploy_id = "deploy-${var.num_instances}"
  new_node_ips = [
    for instance in google_compute_instance.vm :
    instance.network_interface[0].network_ip if try(instance.labels["cohesity-internal-deploy-id"], "") == local.deploy_id
  ]
  existing_node_ips = [
    for instance in google_compute_instance.vm :
    instance.network_interface[0].network_ip if try(instance.labels["cohesity-internal-deploy-id"], "") != local.deploy_id
  ]
  is_create    = length(local.existing_node_ips) == 0
  is_expansion = !local.is_create && length(local.new_node_ips) > 0

  cluster_join_cmd = join(" ", [
    "iris_cli",
    "--v=10",
    "--username=admin",
    "--password=REPLACE_WITH_IRIS_PASSWORD",
    "--skip_password_prompt=true",
    "--skip_force_password_change=true",
    "cluster cloud-join",
    "node-ips=${join(",", local.new_node_ips)}",
  ])
  cluster_join_cmd_with_disk_mapping = join(" ", [
    local.cluster_join_cmd,
    "disk-identifiers=${local.per_node_disk_identifiers}",
    "disk-tiers=${local.per_node_disk_tiers}",
  ])

  # Construct the cloud-create help command for the target image.
  cluster_cloud_create_help_cmd = join(" ", [
    local.iris_cli_cmd_prefix,
    "cluster cloud-create help show-hidden",
  ])

  # Construct cluster status command
  cluster_status_cmd = join(" ", [
    local.iris_cli_cmd_prefix,
    "cluster status"
  ])

  # Construct the SSH-based cluster creation command
  sshpass_cmd_prefix                           = "sshpass -p 'Cohe$1ty' ssh -o StrictHostKeyChecking=no cohesity@${local.instance_ip}"
  sshpass_cluster_create_cmd                   = "${local.sshpass_cmd_prefix} ${local.cluster_create_cmd}"
  sshpass_cluster_create_cmd_with_disk_mapping = "${local.sshpass_cmd_prefix} ${local.cluster_create_cmd_with_disk_mapping}"
  sshpass_cluster_cloud_create_help_cmd        = "${local.sshpass_cmd_prefix} ${local.cluster_cloud_create_help_cmd}"
  sshpass_cluster_status_cmd                   = "${local.sshpass_cmd_prefix} ${local.cluster_status_cmd}"

  # Boot disk type: use config.BootDiskType if present, else default to pd-standard
  boot_disk_type = contains(keys(local.config), "BootDiskType") ? local.config.BootDiskType : "pd-standard"

  # Create a HDD disk map with a stable composite key
  hdd_disks = flatten([
    for vm_index in range(var.num_instances) : [
      for disk_index in range(local.config.HDDTierNumDisks) : {
        key        = "hdd-tier-${vm_index}-${disk_index}" # composite key: VM index and disk index
        vm_index   = vm_index
        disk_index = disk_index
      }
    ]
  ])
  hdd_disk_map = { for d in local.hdd_disks : d.key => d }

  # Create a SSD disk map with a stable composite key
  ssd_disks = flatten([
    for vm_index in range(var.num_instances) : [
      for disk_index in range(local.config.SSDTierNumDisks) : {
        key        = "ssd-tier-${vm_index}-${disk_index}" # composite key: VM index and disk index
        vm_index   = vm_index
        disk_index = disk_index
      }
    ]
  ])
  ssd_disk_map = { for d in local.ssd_disks : d.key => d }

  # SSD Tier disk IOPS: use config.SSDTierDiskIops if present, else null
  ssd_tier_disk_iops = contains(keys(local.config), "SSDTierDiskIops") ? local.config.SSDTierDiskIops : null
  # SSD Tier disk Throughput: use config.SSDTierDiskThroughputinMBps if present, else null
  ssd_tier_disk_throughput = contains(keys(local.config), "SSDTierDiskThroughputinMBps") ? local.config.SSDTierDiskThroughputinMBps : null
  # HDD Tier disk IOPS: use config.HDDTierDiskIops if present, else null
  hdd_tier_disk_iops = contains(keys(local.config), "HDDTierDiskIops") ? local.config.HDDTierDiskIops : null
  # HDD Tier disk Throughput: use config.HDDTierDiskThroughputinMBps if present, else null
  hdd_tier_disk_throughput = contains(keys(local.config), "HDDTierDiskThroughputinMBps") ? local.config.HDDTierDiskThroughputinMBps : null

  # Determine if gVNIC is required based on machine family type
  # C4*, C3*, and N4* families require gVNIC according to GCP documentation
  # Extract the family prefix (e.g., "c4" from "c4-standard", "n4d" from "n4d-standard")
  machine_family_prefix = split("-", var.machine_family_type)[0]
  requires_gvnic = can(regex("^(c4|c3|n4)", local.machine_family_prefix))
}

###############################################################################
# Resource Definitions
###############################################################################

# Create Public IPs for VMs if enabled
resource "google_compute_address" "public_ip" {
  count  = var.attach_public_ip ? var.num_instances : 0
  name   = "${local.resource_name_prefix}-vm-${count.index}-public-ip"
  region = var.region
  labels = local.parsed_labels
}

# Create Boot Disks
resource "google_compute_disk" "boot_disk" {
  count  = var.num_instances
  name   = "${local.resource_name_prefix}-vm-${count.index}-os-disk"
  type   = local.boot_disk_type
  zone   = var.zone
  labels = local.parsed_labels
  image  = var.image_id
  dynamic "disk_encryption_key" {
    for_each = var.customer_managed_kms_key != "" ? [1] : []
    content {
      kms_key_self_link = var.customer_managed_kms_key
    }
  }
  dynamic "guest_os_features" {
    for_each = local.requires_gvnic ? [1] : []
    content {
      type = "GVNIC"
    }
  }

  lifecycle {
    # Prevent replacement of existing cluster nodes when image changes.
    # This enables add-node scenarios where new nodes use a different image
    # (matching the upgraded cluster version) while existing nodes remain untouched.
    # When adding nodes to an upgraded cluster, update image_id to match
    # the current cluster version before increasing num_instances.
    ignore_changes = [image]
  }
}

# Create SSD Tier Data Disks
resource "google_compute_disk" "ssd_tier_data_disk" {
  for_each          = local.ssd_disk_map
  name              = "${local.resource_name_prefix}-${each.key}"
  type              = local.config.SSDTierDiskType
  zone              = var.zone
  size              = local.config.SSDTierDiskSizeinGB
  labels            = local.parsed_labels
  provisioned_iops  = local.ssd_tier_disk_iops
  provisioned_throughput = local.ssd_tier_disk_throughput
  dynamic "disk_encryption_key" {
    for_each = var.customer_managed_kms_key != "" ? [1] : []
    content {
      kms_key_self_link = var.customer_managed_kms_key
    }
  }
}

# Create HDD Tier Data Disks
resource "google_compute_disk" "hdd_tier_data_disk" {
  for_each          = local.hdd_disk_map
  name              = "${local.resource_name_prefix}-${each.key}"
  type              = local.config.HDDTierDiskType
  zone              = var.zone
  size              = local.config.HDDTierDiskSizeinGB
  labels            = local.parsed_labels
  provisioned_iops  = local.hdd_tier_disk_iops
  provisioned_throughput = local.hdd_tier_disk_throughput
  dynamic "disk_encryption_key" {
    for_each = var.customer_managed_kms_key != "" ? [1] : []
    content {
      kms_key_self_link = var.customer_managed_kms_key
    }
  }
}

# Create Virtual Machines
resource "google_compute_instance" "vm" {
  count        = var.num_instances
  name         = "${local.resource_name_prefix}-vm-${count.index}"
  machine_type = local.machine_type
  zone         = var.zone
  labels = merge(local.parsed_labels, {
    "cohesity-internal-deploy-id" = local.deploy_id
  })
  tags = var.network_tags
  #  Uncomment below line in case machine_type, min_cpu_platform or other instance-related attributes needs to be changed.
  # allow_stopping_for_update = true

  boot_disk {
    source = google_compute_disk.boot_disk[count.index].self_link
  }

  network_interface {
    # Reference the subnetwork by its self_link,
    # and derive the network from the subnetwork read itself.
    network    = data.google_compute_subnetwork.network_details.network
    subnetwork = data.google_compute_subnetwork.network_details.self_link
    nic_type   = local.requires_gvnic ? "GVNIC" : null
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
    for_each = { for k, v in google_compute_disk.ssd_tier_data_disk : k => v if local.ssd_disk_map[k].vm_index == count.index }
    content {
      source      = attached_disk.value.self_link
      device_name = local.ssd_device_names[local.ssd_disk_map[attached_disk.key].disk_index]
    }
  }

  # Attach HDD tier disks
  dynamic "attached_disk" {
    for_each = { for k, v in google_compute_disk.hdd_tier_data_disk : k => v if local.hdd_disk_map[k].vm_index == count.index }
    content {
      source      = attached_disk.value.self_link
      device_name = local.hdd_device_names[local.hdd_disk_map[attached_disk.key].disk_index]
    }
  }

  dynamic "service_account" {
    for_each = var.service_account_email != "" ? [1] : []
    content {
      email  = var.service_account_email
      scopes = ["https://www.googleapis.com/auth/cloud-platform"]
    }
  }

  lifecycle {
    # cohesity-internal-deploy-id must keep the value from the apply that created the VM.
    ignore_changes = [labels["cohesity-internal-deploy-id"]]
  }
}

###############################################################################
# Commands Based on Cluster Creation
###############################################################################

# Print cloud-create or cloud-join command
resource "null_resource" "print_command" {
  depends_on = [google_compute_instance.vm]

  triggers = {
    num_instances = tostring(var.num_instances)
    node_ips      = join(",", local.private_ips)
  }

  provisioner "local-exec" {
    command = (
      local.is_expansion ?
      format("%s %s; %s %s",
        "echo If cluster version is 7.5 or later, run the following command on any existing cluster node. Replace REPLACE_WITH_IRIS_PASSWORD with the Iris password:",
        local.cluster_join_cmd_with_disk_mapping,
        "echo Else, run the following command:",
        local.cluster_join_cmd,
      ) :
      local.is_create ?
      format("%s %s; %s %s",
        "echo If cluster version is 7.5 or later, run the following command on any one of the VMs created:",
        local.cluster_create_cmd_with_disk_mapping,
        "echo Else, run the following command:",
        local.cluster_create_cmd,
      ) :
      "echo No new nodes were created, so there is no cluster command to run."
    )
  }
}

# Post boot commands for cluster creation
resource "null_resource" "post_boot_commands" {
  depends_on = [google_compute_instance.vm]
  count      = var.issue_cluster_create_cmd ? 1 : 0

  provisioner "local-exec" {
    command = <<EOT
    if [ "${local.is_create}" != "true" ]; then
      echo "Existing nodes found; skipping cluster create."
      exit 0
    fi
    echo "Waiting for a few minutes for VM to be ready..."
    sleep ${var.post_boot_wait}

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

# Execute cluster create command, checking cloud-create help for disk-identifiers
# support and falling back to the command without disk args on older images.
resource "null_resource" "execute_cluster_create_command" {
  depends_on = [null_resource.post_boot_commands]
  count      = var.issue_cluster_create_cmd ? 1 : 0

  provisioner "local-exec" {
    command = <<EOT
    if [ "${local.is_create}" != "true" ]; then
      echo "Existing nodes found; skipping cluster create."
      exit 0
    fi
    help_cmd_output=$(${local.sshpass_cluster_cloud_create_help_cmd} 2>&1)

    if echo "$help_cmd_output" | grep -q "disk-identifiers"; then
      ${local.sshpass_cluster_create_cmd_with_disk_mapping}
    else
      ${local.sshpass_cluster_create_cmd}
    fi
    EOT
  }
}
