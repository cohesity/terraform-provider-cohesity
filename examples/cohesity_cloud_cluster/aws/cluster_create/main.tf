################################################################################
# Provider Configuration
################################################################################

provider "aws" {
  region = var.region

  # Conditionally assume role if use_iam_role is true
  dynamic "assume_role" {
    for_each = var.use_iam_role ? [1] : []
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

  # Load security group rules
  rules = jsondecode(file(var.rules_file)).rules

  # Select the configuration using the `config_id` variable
  selected_config = lookup(
    { for config in local.configurations : config.Id => config },
    var.config_id,
    null
  )

  # Fallback to `custom_config` if provided; otherwise, use the selected configuration
  config = var.custom_config != null ? var.custom_config : local.selected_config

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

# Security Group (if creating a new one)
resource "aws_security_group" "cohesity_sg" {
  count       = var.use_existing_sg ? 0 : 1
  name        = var.sg_name
  description = "Security Group created for Cohesity cluster Deployment"
  vpc_id      = var.vpc_id

  dynamic "ingress" {
    for_each = local.rules
    content {
      description = ingress.value["description"]
      protocol    = ingress.value["protocol"]
      from_port   = ingress.value["from_port"]
      to_port     = ingress.value["to_port"]
      cidr_blocks = ingress.value["cidr_blocks"]
    }
  }

  tags = merge(local.parsed_tags, {
    Name = var.sg_name
  })
}

# Select Security Group (either existing or new)
locals {
  security_group_id = var.use_existing_sg ? var.existing_sg_id : aws_security_group.cohesity_sg[0].id
}

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
  security_groups             = [local.security_group_id]

  tags = merge(local.parsed_tags, {
    Name = "${local.resource_name_prefix}-vm-${count.index}"
  })
}

# Create EBS Volumes (SSD Tier)
resource "aws_ebs_volume" "ssd_tier_data_disk" {
  count             = var.num_instances * local.config.SSDTierNumDisks
  availability_zone = data.aws_subnet.subnet.availability_zone
  size              = local.config.SSDTierDiskSizeinGB
  type              = local.config.SSDTierDiskType
  iops              = local.config.SSDTierDiskIops
  throughput        = local.config.SSDTierDiskThroughputinMBps

  tags = merge(local.parsed_tags, {
    Name = "${local.resource_name_prefix}-ssd-tier-disk-${count.index}"
  })
}

# Attach EBS Volumes (SSD Tier)
resource "aws_volume_attachment" "attach_ssd" {
  count       = var.num_instances * local.config.SSDTierNumDisks
  instance_id = aws_instance.vm[floor(count.index / local.config.SSDTierNumDisks)].id
  volume_id   = aws_ebs_volume.ssd_tier_data_disk[count.index].id

  # Generate a unique device name for each SSD disk attached to the same instance
  device_name = "/dev/sd${element(["f", "g", "h", "i", "j", "k", "l", "m", "n", "o"], count.index % local.config.SSDTierNumDisks)}"
}

# Create EBS Volumes (HDD Tier)
resource "aws_ebs_volume" "hdd_tier_data_disk" {
  count             = var.num_instances * local.config.HDDTierNumDisks
  availability_zone = data.aws_subnet.subnet.availability_zone
  size              = local.config.HDDTierDiskSizeinGB
  type              = local.config.HDDTierDiskType
  iops              = local.config.HDDTierDiskIops
  throughput        = local.config.HDDTierDiskThroughputinMBps

  tags = merge(local.parsed_tags, {
    Name = "${local.resource_name_prefix}-hdd-tier-disk-${count.index}"
  })
}

# Attach EBS Volumes (HDD Tier)
resource "aws_volume_attachment" "attach_hdd" {
  count       = var.num_instances * local.config.HDDTierNumDisks
  instance_id = aws_instance.vm[floor(count.index / local.config.HDDTierNumDisks)].id
  volume_id   = aws_ebs_volume.hdd_tier_data_disk[count.index].id

  # Generate a unique device name for each HDD disk attached to the same instance
  device_name = "/dev/sd${element(["p", "q", "r", "s", "t", "u", "v", "w", "x", "y"], count.index % local.config.HDDTierNumDisks)}"
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
  count      = var.create_cluster ? 1 : 0

  provisioner "local-exec" {
    command = <<EOT
    echo "Waiting for 5 minutes for VM to be ready..."
    sleep ${var.post_boot_wait}  # Wait for 5 minutes before starting the loop

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
