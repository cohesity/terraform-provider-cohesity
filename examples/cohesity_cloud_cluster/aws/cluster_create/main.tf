################################################################################
# Provider Configuration
################################################################################

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
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

provider "aws" {
  region = var.region

  # Conditionally assume role if iam_role_arn is not empty
  dynamic "assume_role" {
    for_each = var.iam_role_arn != "" ? [1] : []
    content {
      role_arn = var.iam_role_arn
    }
  }
}

################################################################################
# Data Sources
################################################################################

# Fetch details of the existing subnet
data "aws_subnet" "subnet" {
  id = var.subnet_id
}

################################################################################
# Random String for Resource Name Prefix
################################################################################

# Generate an 8-character alphanumeric random string for resource names
resource "random_string" "random_prefix" {
  length  = 8
  special = false
  upper   = false
}

################################################################################
# Local Variables
################################################################################

locals {
  random_prefix = random_string.random_prefix.result

  # Define resource name prefix based on user input
  resource_name_prefix = var.add_random_prefix == false ? var.resource_name_prefix : "${var.resource_name_prefix}-${local.random_prefix}"

  # Extract subnet mask
  subnet_mask = cidrnetmask(data.aws_subnet.subnet.cidr_block)
  # First usable IP is the subnet gateway
  subnet_gateway = cidrhost(data.aws_subnet.subnet.cidr_block, 1)

  # Parse user-provided tags into a key-value map
  parsed_tags = { for tag in var.tags : split(":", tag)[0] => split(":", tag)[1] }

  # Collect private IPs of the created instances
  private_ips = [
    for instance in aws_instance.vm :
    instance.private_ip
  ]

  # Collect public IPs of the created instances (if any)
  public_ips = var.attach_public_ip ? [
    for instance in aws_instance.vm :
    instance.public_ip
  ] : []

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

  # Device name suffixes for SSD tier disks
  # Supports up to 20 disks: f-o (10), then aa-aj (10)
  # Note: We skip p-z to avoid conflicts with HDD which uses that range
  ssd_device_suffixes = concat(
    ["f", "g", "h", "i", "j", "k", "l", "m", "n", "o"], # 0-9
    ["aa", "ab", "ac", "ad", "ae", "af", "ag", "ah", "ai", "aj"] # 10-19
  )

  # Device name suffixes for HDD tier disks
  # Starts at 'p' to avoid conflicts with SSD (which uses f-o, then skips to aa+)
  # Supports up to 20 disks: p-y (10), then z, then ca-ci (9)
  hdd_device_suffixes = concat(
    ["p", "q", "r", "s", "t", "u", "v", "w", "x", "y"], # 0-9
    ["z", "ca", "cb", "cc", "cd", "ce", "cf", "cg", "ch", "ci"] # 10-19
  )

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

# Create Public IPs if attach_public_ip is true
resource "aws_eip" "public_ip" {
  count    = var.attach_public_ip ? var.num_instances : 0
  instance = aws_instance.vm[count.index].id

  tags = merge(local.parsed_tags, {
    Name = "${local.resource_name_prefix}-eip-${count.index}"
  })
}

# Create Instances
resource "aws_instance" "vm" {
  count                       = var.num_instances
  ami                         = var.image_id
  instance_type               = local.config.InstanceType
  subnet_id                   = var.subnet_id
  associate_public_ip_address = var.attach_public_ip
  vpc_security_group_ids      = var.security_group_ids
  iam_instance_profile        = var.instance_profile_name != "" ? var.instance_profile_name : null

  root_block_device {
    volume_type = "gp3"
    encrypted   = var.encrypt_ebs_volumes
    kms_key_id  = var.encrypt_ebs_volumes ? (var.kms_key_id != "" ? var.kms_key_id : null) : null
    tags = merge(local.parsed_tags, {
      Name = "${local.resource_name_prefix}-vm-${count.index}-os-disk"
    })
  }

  metadata_options {
    http_tokens   = var.enforce_imdsv2 ? "required" : "optional"
    http_endpoint = "enabled"
  }

  tags = merge(local.parsed_tags, {
    Name = "${local.resource_name_prefix}-vm-${count.index}"
  })
}

# Create EBS Volumes (SSD Tier)
resource "aws_ebs_volume" "ssd_tier_data_disk" {
  for_each          = local.ssd_disk_map
  availability_zone = data.aws_subnet.subnet.availability_zone
  size              = local.config.SSDTierDiskSizeinGB
  type              = local.config.SSDTierDiskType
  iops              = local.config.SSDTierDiskIops
  throughput        = local.config.SSDTierDiskThroughputinMBps
  encrypted         = var.encrypt_ebs_volumes
  kms_key_id        = var.encrypt_ebs_volumes ? (var.kms_key_id != "" ? var.kms_key_id : null) : null

  tags = merge(local.parsed_tags, {
    Name = "${local.resource_name_prefix}-${each.key}"
  })
}

# Attach EBS Volumes (SSD Tier)
resource "aws_volume_attachment" "attach_ssd" {
  for_each    = local.ssd_disk_map
  instance_id = aws_instance.vm[each.value.vm_index].id
  volume_id   = aws_ebs_volume.ssd_tier_data_disk[each.key].id

  # Generate a unique device name for each SSD disk attached to the same instance
  device_name  = "/dev/sd${element(local.ssd_device_suffixes, each.value.disk_index)}"
  force_detach = true
}

# Create EBS Volumes (HDD Tier)
resource "aws_ebs_volume" "hdd_tier_data_disk" {
  for_each          = local.hdd_disk_map
  availability_zone = data.aws_subnet.subnet.availability_zone
  size              = local.config.HDDTierDiskSizeinGB
  type              = local.config.HDDTierDiskType
  iops              = local.config.HDDTierDiskIops
  throughput        = local.config.HDDTierDiskThroughputinMBps
  encrypted         = var.encrypt_ebs_volumes
  kms_key_id        = var.encrypt_ebs_volumes ? (var.kms_key_id != "" ? var.kms_key_id : null) : null

  tags = merge(local.parsed_tags, {
    Name = "${local.resource_name_prefix}-${each.key}"
  })
}

# Attach EBS Volumes (HDD Tier)
resource "aws_volume_attachment" "attach_hdd" {
  for_each    = local.hdd_disk_map
  instance_id = aws_instance.vm[each.value.vm_index].id
  volume_id   = aws_ebs_volume.hdd_tier_data_disk[each.key].id

  # Generate a unique device name for each HDD disk attached to the same instance
  device_name  = "/dev/sd${element(local.hdd_device_suffixes, each.value.disk_index)}"
  force_detach = true
}

################################################################################
# Commands Based on Cluster Creation
################################################################################

# Print command for manual cluster creation
resource "null_resource" "print_command" {
  depends_on = [aws_volume_attachment.attach_ssd, aws_volume_attachment.attach_hdd]
  provisioner "local-exec" {
    command = "echo Run the following command on any one of the VMs created: ${local.cluster_create_cmd}"
  }
}

# Post boot commands for cluster creation
resource "null_resource" "post_boot_commands" {
  depends_on = [aws_volume_attachment.attach_ssd, aws_volume_attachment.attach_hdd]
  count      = var.issue_cluster_create_cmd ? 1 : 0

  provisioner "local-exec" {
    command = <<EOT
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

# Execute cluster create command
resource "null_resource" "execute_cluster_create_command" {
  depends_on = [null_resource.post_boot_commands]
  count      = var.issue_cluster_create_cmd ? 1 : 0

  provisioner "local-exec" {
    command = local.sshpass_cluster_create_cmd
  }
}
