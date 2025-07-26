###############################################################################
# Authentication Variables
###############################################################################

# Path to the GCP Service Account credentials JSON file.
# If not provided or left empty, Application Default Credentials (ADC) will
# be used.
# To set up ADC, run `gcloud auth application-default login` before invoking
# Terraform.
#
# To generate GCP Service Account credentials JSON file:
#   1. Go to the Google Cloud Console > IAM & Admin > Service Accounts.
#   2. Create a new service account or select an existing one.
#   3. Grant it the required roles (e.g., Compute Admin, Service Account User).
#   4. Click "Manage keys" > "Add Key" > "Create new key" (choose JSON).
#   5. Download the JSON file and provide its absolute path below.
credentials_file_path = "replace-with-the-credentials-file"

###############################################################################
# Deployment Variables
###############################################################################

# The GCP Project ID where resources will be created.
# You can find this in the Google Cloud Console dashboard or by running:
#   gcloud projects list
project_id = "replace-with-project-id"

# Number of Virtual Machines to create in the cluster.
# Example: 3
num_instances = 3

# GCP Region for deployment (e.g., "us-west1", "europe-west3", "asia-east1").
# To list available regions, run:
#   gcloud compute regions list
region = "replace-with-region"

# GCP Zone for deployment (e.g., "us-west1-a", "us-west1-b").
# Must be within the selected region.
# To list available zones, run:
#   gcloud compute zones list --filter="region:(<your-region>)"
zone = "replace-with-zone"

# Name of the existing Virtual Network (VPC) in your project.
# To list VPC networks, run:
#   gcloud compute networks list
vnet_name = "replace-with-your-vnet-name"

# Name of the existing Subnet within the VPC.
# To list subnets, run:
#   gcloud compute networks subnets list --regions=<your-region>
subnet_name = "replace-with-your-subnet-name"

# The full ID of the Cohesity VM image to use.
# Format: "projects/<project>/global/images/<image>"
# You may need to request access to the image from Cohesity.
# To list available images, run:
#   gcloud compute images list --project=<project>
image_id = "replace-with-shared-image-id"

# Labels to apply to the resources created.
# Format: ["key1:value1", "key2:value2"]
# These help with resource organization and cost tracking.
labels = [
  "replace-with-key1:replace-with-value1",
  "replace-with-key2:replace-with-value2"
]

# (Optional) Network tags to apply to the VMs (for firewall rules, etc.).
# These tags must match those used in your firewall rules.
# To list existing tags, check your firewall rules in the GCP Console.
# Example: ["internal", "prod"]
network_tags = [
  "replace-with-network-tag1",
  "replace-with-network-tag2"
]

# Prefix for naming all GCP resources (VMs, disks, NICs, etc.).
# Restrictions:
#   - Only lowercase letters (a-z), numbers (0-9), and hyphens (-).
#   - Must start with a lowercase letter and not end with a hyphen.
#   - Max 30 characters.
# Example: If you set resource_name_prefix = "cohesity", the resources will be
# named as follows:
# - Virtual Machine: cohesity-vm-0
# - NIC: cohesity-vm-0-nic-0
# - Boot Disk: cohesity-vm-0-boot-disk
# - SSD Data Disk: cohesity-vm-0-ssd-disk-0
# - HDD Data Disk: cohesity-vm-0-hdd-disk-0
resource_name_prefix = "replace-with-prefix"

# If true, a random 8-character string will be appended to the
# resource_name_prefix. This helps avoid name collisions when deploying
# multiple clusters.
add_random_prefix = false

# Set to true to attach public IPs to all GCP VMs created.
# This can be set to true if terraform is being run outside of the VPC and you
# do not have SSH access to the VMs. In this case, GCP VMs will come up with
# public IPs, allowing SSH commands to be issued for cluster creation.
# In most cases, when running terraform from within the VPC which has SSH
# access to the VMs being deployed, this can be set as default false.
attach_public_ip = false

# (Optional) Email of a pre-created service account to attach to the VMs.
# Leave empty to create VMs without a service account.
# Example: "my-service-account@my-project.iam.gserviceaccount.com"
service_account_email = ""

# (Optional) The full resource ID of the customer-managed KMS key to use for disk encryption.
# Example: "projects/<project>/locations/<location>/keyRings/<keyring>/cryptoKeys/<key>"
customer_managed_kms_key = ""

###############################################################################
# Config Variables
###############################################################################

# Configuration ID to use for VM deployment.
# Choose a config id from the configs.json file in this directory.
# Example: "1", "2", "3", etc.
config_id = "<replace-with-config-id>"

# The machine family type for the instance.
# Example: "n2d-standard", "n2-standard".
# This is combined with the CPU count from the config to form the machine type.
machine_family_type = "replace-with-machine-type"

# (Optional) Specify a custom configuration to override the config_id.
# Uncomment and fill the block below to use a custom VM configuration.
# custom_config = {
#   "InstanceCPUCount" : "16",
#   "SSDTierNumDisks" : 2,
#   "SSDTierDiskSizeinGB" : 1023,
#   "SSDTierDiskType" : "pd-balanced",
#   "HDDTierNumDisks" : 2,
#   "HDDTierDiskSizeinGB" : 1024,
#   "HDDTierDiskType" : "pd-balanced"
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
# Defaults to 192.168.0.0 if left empty.
apps_subnet = "replace-with-apps-subnet"

# A private IPv4 subnet mask for the internal overlay network of the
# kubernetes cluster to run the apps infrastructure.
# Defaults to 255.255.0.0 if left empty.
apps_subnet_mask = "replace-with-apps-subnet-mask"

# Whether to issue the cluster creation command via SSH after VM deployment.
# Usually set to true (default) as Terraform is expected to run in the same VPC
# as the VMs, or public IPs are attached for SSH access.
# Set to false if you want to create the cluster manually.
issue_cluster_create_cmd = true
