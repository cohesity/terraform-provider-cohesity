###############################################################################
# Authentication Overview for Terraform on GCP
###############################################################################

# Terraform (Google Provider) follows this order to find GCP credentials:
#
# 1. Provider-Specific Credentials (Highest Priority)
#    - If the provider's `credentials` argument (e.g., via `credentials_file_path`) is set,
#      it is ALWAYS used and overrides all other methods below.
#    Steps:
#      a. In Google Cloud Console, visit IAM & Admin > Service Accounts.
#      b. Create/select a service account (with the needed roles, like Compute Admin).
#         See the docs for detailed permissions.
#      c. Click "Manage keys" > "Add Key" > "Create new key" (choose JSON).
#      d. Download the JSON file.
#      e. Set the `credentials_file_path` variable below to the absolute path of this file.
#
# 2. GOOGLE_CREDENTIALS Environment Variable
#    - Set this variable with the contents of your Service Account JSON file (not a path).
#    Steps:
#      a. Export the file content, not path:
#         export GOOGLE_CREDENTIALS="$(cat /path/to/key.json)"
#
# 3. GOOGLE_APPLICATION_CREDENTIALS Environment Variable
#    - Set this variable with the absolute path to your Service Account JSON file.
#    Steps:
#      a. Download your service account key file as above.
#      b. export GOOGLE_APPLICATION_CREDENTIALS="/absolute/path/to/key.json"
#
# 4. Application Default Credentials from gcloud CLI
#    - If no above env vars, Terraform checks for credentials created by the gcloud CLI.
#    Steps:
#      a. Install the gcloud CLI (https://cloud.google.com/sdk/docs/install)
#      b. Run:
#         gcloud auth application-default login
#      c. This will create a file at: ~/.config/gcloud/application_default_credentials.json
#    Note: Terraform does not use gcloud directly, but relies on the credentials file it creates.
#
# 5. GCP Environment (Compute Engine)
#    - If all above not set, and running on a GCP VM, uses the attached service account.
#    Steps:
#      a. Ensure your Compute Engine instance has a service account with the necessary IAM roles.
#      b. Attach this service account when creating the resource or update an existing one:
#         See: https://cloud.google.com/compute/docs/access/service-accounts
#
# Reference: https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference
#
# Best practices:
#  - Use least-privilege roles for Service Accounts.
#  - Prefer Service Accounts over personal logins for shared/deployment environments.

###############################################################################
# Authentication Variables
###############################################################################

# Provider-Specific Credentials (Option 1 - Highest Priority, see above):
# ----------------------------------------------------------------------
# Set 'credentials_file_path' below to the absolute path of your GCP Service
# Account credentials JSON file. This will be used if your provider
# configuration is set up to use this variable, and will override all other
# authentication mechanisms listed above.
#
# Example: credentials_file_path = "/absolute/path/to/your/credentials.json"
credentials_file_path = ""

# (Optional) Email address of the GCP service account to impersonate for all operations.
# If set, Terraform will act as this service account instead of the one from the credentials file or ADC.
# The identity running Terraform must have Service Account Token Creator (roles/iam.serviceAccountTokenCreator)
# on the impersonated service account.
# Example: impersonate_service_account = "my-impersonated-sa@my-gcp-project.iam.gserviceaccount.com"
impersonate_service_account = ""

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
# Note that the identity used to run this terraform should have Service Account
# User role (roles/iam.serviceAccountUser) on the service account to be attached.
# Example: "my-attached-service-account@my-project.iam.gserviceaccount.com"
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

# The GCP machine family type for the VM instances.
# Example values: "n2d-standard", "n2-standard", "n4-standard", "n4d-standard"
# This value is combined with the CPU count from the config to form the final machine type (e.g., "n4-standard-16").
# Important:
#            N2* instance types (n2, n2d) support Persistent Disk configurations but do not support Hyperdisk. Use config IDs: 1, 2, 3, 4 (see configs.json for details).
#            N4* and C4* instance types (n4,n4d,c4,c4d) require Hyperdisks. Use config IDs: 6 (see configs.json for details).
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
