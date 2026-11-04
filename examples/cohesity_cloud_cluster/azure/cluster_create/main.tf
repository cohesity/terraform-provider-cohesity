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

  # Restricted subscriptions (e.g. Azure Gov) scope automation SPs to only the
  # actions this module needs, so azurerm's default "core" auto-registration
  # of ~35 unrelated resource providers 403s and aborts the apply; this module
  # only needs Microsoft.Resources/Storage/Network/Compute(/Authorization when
  # system managed identity is enabled), which are pre-registered.
  resource_provider_registrations = "none"

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

  # LUNs assigned during disk attachment
  ssd_luns = [for i in range(local.config.SSDTierNumDisks) : i]
  hdd_luns = [for i in range(local.config.HDDTierNumDisks) : i + local.config.SSDTierNumDisks]

  # Build the per-node disk-identifier -> tier mapping passed to iris_cli
  per_node_disk_identifiers = join(",", concat(
    [for lun in local.ssd_luns : "/dev/disk/azure/scsi1/lun${lun}"],
    [for lun in local.hdd_luns : "/dev/disk/azure/scsi1/lun${lun}"]
  ))

  per_node_disk_tiers = join(",", concat(
    [for _ in local.ssd_luns : "PCIeSSD"],
    [for _ in local.hdd_luns : "SATA-HDD"]
  ))

  # Compute VM identity type based on which identity mechanisms are enabled.
  # Supports four states: none, SystemAssigned only, UserAssigned only, or both.
  has_smi = var.enable_system_managed_identity
  has_umi = length(var.user_assigned_managed_identities_to_attach) > 0
  identity_type = local.has_smi && local.has_umi ? "SystemAssigned, UserAssigned" : (
    local.has_smi ? "SystemAssigned" : (
      local.has_umi ? "UserAssigned" : ""
    )
  )

  # Cross-product of VMs × role assignments for SMI RBAC setup
  smi_role_assignments = var.enable_system_managed_identity ? flatten([
    for vm_index in range(var.num_instances) : [
      for ra_index, ra in var.system_managed_identity_role_assignments : {
        key                  = "vm${vm_index}-role${ra_index}"
        vm_index             = vm_index
        role_definition_name = ra.role_definition_name
        scope                = ra.scope
      }
    ]
  ]) : []

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

  # Nodes created while num_instances has this value are tagged with it.
  # The tag is ignored on update, so earlier nodes keep their own value.
  # Nodes created before this tag existed have none and count as existing.
  deploy_id = "deploy-${var.num_instances}"
  new_node_ips = [
    for i, vm in azurerm_virtual_machine.azure_vm :
    local.private_ips[i] if try(vm.tags["cohesity-internal-deploy-id"], "") == local.deploy_id
  ]
  existing_node_ips = [
    for i, vm in azurerm_virtual_machine.azure_vm :
    local.private_ips[i] if try(vm.tags["cohesity-internal-deploy-id"], "") != local.deploy_id
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
  # Must match the VM's zone; without this the OS disk is regional and cannot
  # attach to a zonal VM (required when using PremiumV2_LRS data disks).
  zone                   = var.availability_zone
  disk_encryption_set_id = var.customer_managed_disk_encryption_id != "" ? var.customer_managed_disk_encryption_id : null
  tags                   = local.parsed_tags

  lifecycle {
    # Prevent replacement of existing cluster nodes when VHD URI changes.
    # This enables add-node scenarios where new nodes use a different VHD
    # (matching the upgraded cluster version) while existing nodes remain untouched.
    # When adding nodes to an upgraded cluster, update vhd_uri to match
    # the current cluster version before increasing num_instances.
    ignore_changes = [source_uri]
  }
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
      lun             = local.ssd_luns[storage_data_disk.value.disk_index]
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
      lun             = local.hdd_luns[storage_data_disk.value.disk_index]
      caching         = "None"
      create_option   = "Attach"
      disk_size_gb    = local.config.HDDTierDiskSizeinGB
      managed_disk_id = azurerm_managed_disk.hdd_tier_data_disk[storage_data_disk.key].id
    }
  }

  dynamic "identity" {
    for_each = local.identity_type != "" ? [1] : []
    content {
      type         = local.identity_type
      identity_ids = local.has_umi ? var.user_assigned_managed_identities_to_attach : []
    }
  }

  tags = merge(local.parsed_tags, {
    "cohesity-internal-deploy-id" = local.deploy_id
  })

  lifecycle {
    # cohesity-internal-deploy-id must keep the value from the apply that created the VM.
    ignore_changes = [tags["cohesity-internal-deploy-id"]]
  }
}

################################################################################
# System Managed Identity — RBAC Role Assignments
################################################################################

resource "azurerm_role_assignment" "smi_roles" {
  for_each = { for ra in local.smi_role_assignments : ra.key => ra }

  scope                = each.value.scope
  role_definition_name = each.value.role_definition_name
  principal_id         = azurerm_virtual_machine.azure_vm[each.value.vm_index].identity[0].principal_id
}

################################################################################
# Commands Based on Cluster Creation
################################################################################

# Print cloud-create or cloud-join command
resource "null_resource" "print_command" {
  depends_on = [azurerm_virtual_machine.azure_vm]

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
  depends_on = [azurerm_virtual_machine.azure_vm]
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
