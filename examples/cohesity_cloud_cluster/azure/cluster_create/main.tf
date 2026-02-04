################################################################################
# Provider Configuration
################################################################################

terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.0"
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

provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
  environment     = var.environment

  # Credentials are interpreted based on auth_method:
  # - Service Principal: client_id + client_secret + tenant_id all required
  # - Managed Identity: client_id optional (only required for multiple user-assigned identities)
  # - Azure CLI: all empty, provider uses 'az login' credentials automatically
  client_id     = var.client_id != "" ? var.client_id : null
  client_secret = var.client_secret != "" ? var.client_secret : null
  tenant_id     = var.tenant_id != "" ? var.tenant_id : null

  # Enable managed identity authentication when specified
  use_msi = var.auth_method == "managed_identity"
}

################################################################################
# Data Sources
################################################################################

# Fetch details of the existing subnet
data "azurerm_subnet" "example" {
  name                 = local.subnet_name
  virtual_network_name = local.virtual_network_name
  resource_group_name  = local.subnet_resource_group_name
}

# Validate if the resource group exists if provided
data "azurerm_resource_group" "existing_rg" {
  count = var.resource_group_name == "" ? 0 : 1
  name  = var.resource_group_name
}

################################################################################
# Random String for Resource Names
################################################################################

# Generate an 8-character alphanumeric random string for resource names
resource "random_string" "unique_id" {
  length  = 8
  special = false
  upper   = false
}

################################################################################
# Local Variables
################################################################################

locals {
  random_prefix = random_string.unique_id.result

  # Define resource name prefix based on user input
  resource_name_prefix = var.add_random_prefix == false ? var.resource_name_prefix : "${var.resource_name_prefix}-${local.random_prefix}"

  # Set the resource group name for resource creation
  resource_group_name = var.resource_group_name == "" ? azurerm_resource_group.resource_group[0].name : var.resource_group_name

  # Extract subnet details from the Subnet ID
  subnet_resource_group_name = element(split("/", var.subnet_id), 4)
  virtual_network_name       = element(split("/", var.subnet_id), 8)
  subnet_name                = element(split("/", var.subnet_id), 10)

  # Extract Address Prefix and Subnet Gateway
  # Address Prefix example "10.100.128.0/21"
  address_prefix = data.azurerm_subnet.example.address_prefix != "" ? data.azurerm_subnet.example.address_prefix : data.azurerm_subnet.example.address_prefixes[0]
  # Extract subnet mask
  subnet_mask = cidrnetmask(local.address_prefix)
  # First usable IP is the subnet gateway
  subnet_gateway = cidrhost(local.address_prefix, 1)

  # Parse user-provided tags into a key-value map
  parsed_tags = { for tag in var.tags : split(":", tag)[0] => split(":", tag)[1] }

  # Collect private IPs of the created NICs
  private_ips = [
    for nic in azurerm_network_interface.nic :
    nic.private_ip_address
  ]

  # Collect public IPs of the created public IPs (if any)
  public_ips = var.attach_public_ip ? [for ip in azurerm_public_ip.public_ip : ip.ip_address] : []

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

# Create a new resource group if enabled
resource "azurerm_resource_group" "resource_group" {
  count    = var.resource_group_name == "" ? 1 : 0
  name     = "${local.resource_name_prefix}-rg"
  location = var.region

  tags = local.parsed_tags
}

# Create public IPs for VMs if enabled
resource "azurerm_public_ip" "public_ip" {
  count               = var.attach_public_ip ? var.num_instances : 0
  name                = "${local.resource_name_prefix}-vm-${count.index}-public-ip"
  location            = var.region
  resource_group_name = local.resource_group_name
  allocation_method   = "Static"

  tags = local.parsed_tags
}

# Create network interfaces for VMs
resource "azurerm_network_interface" "nic" {
  count               = var.num_instances
  name                = "${local.resource_name_prefix}-vm-${count.index}-nic"
  location            = var.region
  resource_group_name = local.resource_group_name
  accelerated_networking_enabled = true

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = var.attach_public_ip ? azurerm_public_ip.public_ip[count.index].id : null
  }

  tags = local.parsed_tags
}

# Associate Network Security Group with NICs if provided
resource "azurerm_network_interface_security_group_association" "attach_nsg" {
  count = var.network_security_group_id != null ? var.num_instances : 0

  network_interface_id      = azurerm_network_interface.nic[count.index].id
  network_security_group_id = var.network_security_group_id
}

# Create managed OS disks
resource "azurerm_managed_disk" "os_disk" {
  count                  = var.num_instances
  name                   = "${local.resource_name_prefix}-vm-${count.index}-os-disk"
  location               = var.region
  resource_group_name    = local.resource_group_name
  storage_account_type   = "Premium_LRS"
  create_option          = "Import"
  source_uri             = var.vhd_uri
  storage_account_id     = var.storage_account_id
  hyper_v_generation     = "V1"
  os_type                = "Linux"
  disk_encryption_set_id = var.customer_managed_disk_encryption_id != "" ? var.customer_managed_disk_encryption_id : null
  tags                   = local.parsed_tags
}

# Create managed disks for SSD tier
resource "azurerm_managed_disk" "ssd_tier_data_disk" {
  for_each             = local.ssd_disk_map
  name                 = "${local.resource_name_prefix}-${each.key}"
  location             = var.region
  resource_group_name  = local.resource_group_name
  storage_account_type = local.config.SSDTierDiskType
  disk_size_gb         = local.config.SSDTierDiskSizeinGB
  disk_iops_read_write = local.config.SSDTierDiskType == "PremiumV2_LRS" ? local.config.SSDTierDiskIops : null
  disk_mbps_read_write = local.config.SSDTierDiskType == "PremiumV2_LRS" ? local.config.SSDTierDiskThroughputinMBps : null
  create_option        = "Empty"
  zone                 = var.availability_zone
  disk_encryption_set_id = var.customer_managed_disk_encryption_id != "" ? var.customer_managed_disk_encryption_id : null
  tags                 = local.parsed_tags
}

# Create managed disks for HDD tier
resource "azurerm_managed_disk" "hdd_tier_data_disk" {
  for_each             = local.hdd_disk_map
  name                 = "${local.resource_name_prefix}-${each.key}"
  location             = var.region
  resource_group_name  = local.resource_group_name
  storage_account_type = local.config.HDDTierDiskType
  disk_size_gb         = local.config.HDDTierDiskSizeinGB
  disk_iops_read_write = local.config.HDDTierDiskType == "PremiumV2_LRS" ? local.config.HDDTierDiskIops : null
  disk_mbps_read_write = local.config.HDDTierDiskType == "PremiumV2_LRS" ? local.config.HDDTierDiskThroughputinMBps : null
  create_option        = "Empty"
  zone                 = var.availability_zone
  disk_encryption_set_id = var.customer_managed_disk_encryption_id != "" ? var.customer_managed_disk_encryption_id : null
  tags                 = local.parsed_tags
}

# Create Virtual Machines
resource "azurerm_virtual_machine" "azure_vm" {
  count               = var.num_instances
  name                = "${local.resource_name_prefix}-vm-${count.index}"
  location            = var.region
  resource_group_name = local.resource_group_name
  network_interface_ids = [
    azurerm_network_interface.nic[count.index].id
  ]
  delete_os_disk_on_termination    = true
  delete_data_disks_on_termination = true
  vm_size                          = local.config.InstanceType
  zones                            = var.availability_zone != null ? [var.availability_zone] : null

  storage_os_disk {
    name              = "${local.resource_name_prefix}-vm-${count.index}-os-disk"
    caching           = "None"
    create_option     = "Attach"
    os_type           = "Linux"
    managed_disk_type = "Premium_LRS"
    managed_disk_id   = azurerm_managed_disk.os_disk[count.index].id
  }

  # Attach SSD Tier Data Disks
  dynamic "storage_data_disk" {
    for_each = { for k, v in local.ssd_disk_map : k => v if v.vm_index == count.index }
    content {
      name            = azurerm_managed_disk.ssd_tier_data_disk[storage_data_disk.key].name
      lun             = storage_data_disk.value.disk_index
      caching         = "None"
      create_option   = "Attach"
      disk_size_gb    = local.config.SSDTierDiskSizeinGB
      managed_disk_id = azurerm_managed_disk.ssd_tier_data_disk[storage_data_disk.key].id
    }
  }

  # Attach HDD Tier Data Disks
  dynamic "storage_data_disk" {
    for_each = { for k, v in local.hdd_disk_map : k => v if v.vm_index == count.index }
    content {
      name            = azurerm_managed_disk.hdd_tier_data_disk[storage_data_disk.key].name
      lun             = storage_data_disk.value.disk_index + local.config.SSDTierNumDisks
      caching         = "None"
      create_option   = "Attach"
      disk_size_gb    = local.config.HDDTierDiskSizeinGB
      managed_disk_id = azurerm_managed_disk.hdd_tier_data_disk[storage_data_disk.key].id
    }
  }

  dynamic "identity" {
    for_each = length(var.user_assigned_managed_identities_to_attach) > 0 ? [1] : []
    content {
      type         = "UserAssigned"
      identity_ids = var.user_assigned_managed_identities_to_attach
    }
  }

  tags = local.parsed_tags
}

################################################################################
# Commands Based on Cluster Creation
################################################################################

# Print command for manual cluster creation
resource "null_resource" "print_command" {
  depends_on = [azurerm_virtual_machine.azure_vm]
  provisioner "local-exec" {
    command = "echo Run the following command on any one of the VMs created: ${local.cluster_create_cmd}"
  }
}

# Post boot commands for cluster creation
resource "null_resource" "post_boot_commands" {
  depends_on = [azurerm_virtual_machine.azure_vm]
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
