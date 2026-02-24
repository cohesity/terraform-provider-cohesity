###############################################################################
# Authentication Overview
###############################################################################

# You can use one of the following authentication methods to allow Terraform to
# create resources in your OCI account. Terraform uses the OCI SDK credential
# chain in the following priority order to determine which method to use:
#  1) Environment variables
#  2) OCI Config Files
#  3) Instance Principals
#
# Methods:
# 1) Environment variables (API key authentication)
#    export OCI_CLI_USER=ocid1.user.oc1..xxxxx
#    export OCI_CLI_FINGERPRINT=xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx
#    export OCI_CLI_TENANCY=ocid1.tenancy.oc1..xxxxx
#    export OCI_CLI_REGION=us-ashburn-1
#    export OCI_CLI_KEY_FILE=~/.oci/oci_api_key.pem
#
# 2) OCI Config Files
# Terraform does not call the OCI CLI. Installing it is optional and useful to
# run 'oci setup config' to create/update the shared files.
#    Files: ~/.oci/config
#    Default: run 'oci setup config' or create entries manually.
#    See OCI docs https://docs.oracle.com/en-us/iaas/Content/API/Concepts/cliconcepts.htm for details.
#
# 3) Instance Principals
#    When running terraform on an OCI compute instance with an instance principal
#    attached, if above two methods are not used, terraform will discover the credentials
#    automatically by leveraging OCI Instance Metadata Service.
#
# Note:
#  - Whichever method is chosen, the identity must have the permissions needed
#    to create/modify/delete the OCI resources this terraform module creates,
#     for example compute instances, block volumes, etc.
#  - Instance Principals method auth works only on OCI compute instance.

###############################################################################
# Deployment Variables
###############################################################################

# OCI region for deployment.
# Example: "us-sanjose-1", "us-chicago-1"
# To list available regions:
#   oci iam region list
region = "replace-with-region"

# OCID of the compartment where resources will be created.
# Example: "ocid1.compartment.oc1..xxxxx"
# To list available compartments:
compartment_id = "replace-with-compartment-id"

# Availability domain for the instances.
# Example: "JCMe:US-SANJOSE-1-AD-1"
# To list available availability domains:
#   oci iam availability-domain list
availability_domain = "replace-with-availability-domain"

# Number of compute instances to create in the cluster.
# Example: 3
num_instances = 3

# The prefix to use for naming all resources (e.g., Virtual Machines, OS Disks,
# and Data Disks). This prefix will be prepended to all resource names, followed
# by additional identifiers.
#
# Example: If you set resource_name_prefix = "cohesity", the resources will be
# named as follows:
# - Virtual Machine: cohesity-vm-0
# - OS Disk: cohesity-vm-0-os-disk
# - SSD Data Disk: cohesity-vm-0-ssd-disk-0
# - HDD Data Disk: cohesity-vm-0-hdd-disk-0
#
# Restrictions:
# - Only letters (a-z, A-Z), numbers (0-9), and hyphens (-)
# - Max 30 characters
# - Cannot start or end with a hyphen
# - No consecutive hyphens (e.g., "--")
resource_name_prefix = "replace-with-prefix"

# Set to true to add a random 8-character string to all resource names (optional)
add_random_prefix = false

# OCID of the subnet where compute instances will be deployed.
# Example: "ocid1.subnet.oc1.iad.xxxxx"
# To list subnets:
#   oci network subnet list --compartment-id <compartment-ocid>
subnet_id = "replace-with-subnet-id"

# List of network security group (NSG) IDs to attach to the compute instances.
# Example: ["ocid1.networksecuritygroup.oc1..xxxxx"]
# Use the NSG ID output from the security_group_create terraform
# module or find with:
#   oci network nsg list --compartment-id <compartment-ocid>
nsg_ids = ["replace-with-nsg-id"]

# Image OCID for the compute instances (should be the Cohesity-provided image).
# Example: "ocid1.image.oc1..xxxxx"
# To list images:
#   oci compute image list --compartment-id <compartment-ocid>
image_id = "replace-with-image-id"

# Tags to apply to all OCI resources.
# Format: ["key1:value1", "key2:value2"]
# Example: ["environment:dev", "owner:your-name"]
tags = [
  "key1:value1",
  "key2:value2"
]

# Set to true to attach public IPs to all compute VMs created.
attach_public_ip = false

###############################################################################
# Config Variables
#
# These control the VM configuration. You can use a predefined config from a
# JSON file or override with a custom config.
###############################################################################

# ID of the configuration to choose from the JSON file.
# Example: "1", "2", "3", etc. (see configs.json for available IDs)
config_id = "replace-with-config-id"

# Path to the JSON configuration file. Keep this as is unless you have a
# custom location for your configs.
# Example: "./configs.json"
config_file = "./configs.json"

# Custom configuration of the VM which overrides the config provided by the
# config_id. If you want to use a custom setup, uncomment below and the config will be overridden with the custom config.
#
# Example:
# custom_config = {
#   InstanceType                = "VM.Standard2.8"
#   SSDTierNumDisks             = 1
#   SSDTierDiskSizeinGB         = 1023
#   HDDTierNumDisks             = 1
#   HDDTierDiskSizeinGB         = 1024
# }

###############################################################################
# Cluster Create Variables
###############################################################################

# The name of the Cohesity cluster. Specify a name up to 63 alphanumeric
# characters such as Cohesity1MyCompany. Hyphens are allowed but cannot be
# the first or last character. This is the name of the Cohesity cluster as
# displayed in the Cohesity UI and the Cohesity CLI.
cluster_name = "replace-with-cluster-name"

# The IP addresses of the Domain Name System (DNS) servers that the Cohesity
# cluster should use. Separate multiple IPs with commas.
dns_server_ips = "replace-with-dns-server-ips"

# Cohesity recommends using the external Google Public Network Time Protocol
# (NTP) server and specifying multiple servers (time1.google.com,
# time2.google.com, time3.google.com, time4.google.com).
ntp_servers = "replace-with-ntp-servers"

# The domain name(s) for the Cohesity cluster. (comma-separated if multiple)
domain_names = "replace-with-domain-names"

# A private IPv4 subnet for the internal overlay network of the kubernetes
# cluster to run the apps infrastructure.
# Update accordingly.
apps_subnet = "192.168.0.0"

# A private IPv4 subnet mask for the internal overlay network of the
# kubernetes cluster to run the apps infrastructure.
# Update accordingly.
apps_subnet_mask = "255.255.0.0"

# Whether to issue the cluster creation command via SSH after VM deployment.
# Usually set to true (default) as Terraform is expected to run in the same VPC
# as the VMs, or public IPs are attached for SSH access.
# Set to false if you want to create the cluster manually.
issue_cluster_create_cmd = true
