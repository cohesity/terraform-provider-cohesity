###############################################################################
# Authentication Overview
###############################################################################

# You can use one of the following authentication methods to allow Terraform to
# create resources in your AWS account. Terraform uses the AWS SDK credential
# chain in the following priority order to determine which method to use:
#  1) Environment variables
#  2) AWS Credentials & Config Files
#  3) EC2 Instance Roles
#
# Methods:
# 1) Environment variables (IAM User access key and secret or temporary credentials)
#    export AWS_ACCESS_KEY_ID=your_access_key
#    export AWS_SECRET_ACCESS_KEY=your_secret_key
#    export AWS_SESSION_TOKEN=your_session_token   # for temp creds
#
# 2) AWS Credentials & Config Files
# Terraform does not call the AWS CLI. Installing it is optional and useful to
# run 'aws configure' to create/update the shared files.
#    Files: ~/.aws/credentials and ~/.aws/config
#    Default: run 'aws configure' or create [default] entries manually.
#    Named: run 'aws configure --profile myprofile' or create entries.
#    If using a named profile ensure to set the environment variable:
#      - export AWS_PROFILE=myprofile
#    See AWS docs https://docs.aws.amazon.com/cli/v1/userguide/cli-configure-files.html for details.
#
# 3) EC2 Instance Roles
#    When running terraform on an AWS EC2 instance with an IAM role attached,
#    if above two methods are not used, terraform will discover the credentials
#    automatically by leveraging EC2 Instance Metadata Service.
#
# Note:
#  - Whichever method is chosen, the identity must have the permissions needed
#    to create/modify/delete the AWS resources this terraform module creates,
#     for example EC2 instance, EBS volumes, etc.
#  - EC2 Instance Roles method auth works only on EC2 instance.

###############################################################################
# Authentication Options
###############################################################################

# (Optional) IAM Role ARN to assume for deployment. Using any of the above
# authentication methods, this parameter tells terraform to assume a role to
# perform actions on AWS. Note that the base identity only needs to have
# sts:AssumeRole permission for the target role which you want terraform to
# assume.
# Example: "arn:aws:iam::123456789012:role/my-terraform-role"
iam_role_arn = ""

###############################################################################
# Deployment Variables
###############################################################################

# AWS region for deployment.
# Example: "us-west-1", "us-east-2"
# To list available regions:
#   aws ec2 describe-regions --query "Regions[*].RegionName"
region = "replace-with-region"

# Number of EC2 instances to create in the cluster.
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

# ID of the subnet where EC2 instances will be deployed.
# Example: "subnet-abc1234def5678"
# To list subnets:
#   aws ec2 describe-subnets --query "Subnets[*].SubnetId"
subnet_id = "subnet-xxxxxxxx"

# List of security group IDs to attach to the EC2 instances.
# Example: ["sg-0a1b2c3d4e5f6g7h"]
# Use the security group ID output from the security_group_create terraform
# module or find with:
#   aws ec2 describe-security-groups --query "SecurityGroups[*].GroupId"
security_group_ids = ["sg-xxxxxxxxxxxxxxxxx"]

# AMI ID for the EC2 instances (should be the Cohesity-provided image).
# Example: "ami-abc1234def5678"
# To list AMIs:
#   aws ec2 describe-images --owners self amazon --query "Images[*].ImageId"
image_id = "ami-xxxxxxxxxxxxxxxxx"

# Tags to apply to all AWS resources.
# Format: ["key1:value1", "key2:value2"]
# Example: ["environment:dev", "owner:your-name"]
# Note: Tags help identify and organize resources in AWS.
tags = [
  "key1:value1",
  "key2:value2"
]

# Set to true to attach public IPs to all EC2 VMs created.
# This can be set to true if terraform is being run outside of the VPC and you
# do not have SSH access to the VMs. In this case, EC2 VMs will come up with
# public IPs, allowing SSH commands to be issued for cluster creation.
# In most cases, when running terraform from within the VPC which has SSH
# access to the VMs being deployed, this can be set as default false.
attach_public_ip = false

# Enforce AWS Instance Metadata Service Version 2 (IMDSv2).
# true: Enforces IMDSv2 (recommended for security)
# false: Allows IMDSv1 (legacy compatibility)
enforce_imdsv2 = false

# Enable or disable encryption for EBS volumes attached to the EC2 instances.
# - Set to true to encrypt all EBS volumes.
# - Set to false if you do not want encryption.
# - If true, Terraform will use the KMS key specified in kms_key_id (if provided),
#   otherwise AWS will use the default EBS encryption key for your account and region.
encrypt_ebs_volumes = false

# (Optional) The ARN of the AWS KMS key to use for encrypting EBS volumes.
# - Example: "arn:aws:kms:us-west-2:123456789012:key/abcd-1234-..."
# - Leave blank ("") to use the AWS default EBS encryption key.
# Prerequisites:
#   - The KMS key must exist in the same AWS region as your deployment.
#   - The IAM user or role running Terraform must have permission to use this KMS key
#   - To find available KMS keys, use:
#       aws kms list-keys
#   - To get the ARN for a key, use:
#       aws kms describe-key --key-id <key-id>
kms_key_id = ""

# (Optional) Name of an IAM instance profile to attach to each EC2 instance.
#
# This instance profile is typically used to grant temporary credentials to the EC2 VMs.
# In most deployments, the role associated with this instance profile should include permissions
# to assume other IAM roles which have the required access to AWS services such as EC2, EBS, S3, etc.
#
# Use cases this should be used:
# - Registering S3 external target with IAM role for backup and archival purposes.
# - AWS Source registration with IAM role for accessing AWS APIs from within
#   the VMs for operations such as snapshotting, backup etc.
#
# Prerequisites:
# - The instance profile must be created in your AWS account
# - The role associated with the profile should have at minimum permissions to
#   call sts:AssumeRole on the IAM roles to assume.
#
# Example:
#   instance_profile_name = "CohesityInstanceProfile"
instance_profile_name = ""

###############################################################################
# Config Variables
#
# These control the VM configuration. You can use a predefined config from a
# JSON file or override with a custom config.
###############################################################################

# ID of the configuration to choose from the JSON file.
# Example: "1", "2", "3", etc. (see configs.json for available IDs)
config_id = "2"

# Path to the JSON configuration file. Keep this as is unless you have a
# custom location for your configs.
# Example: "./configs.json"
config_file = "./configs.json"

# Custom configuration of the VM which overrides the config provided by the
# config_id. If you want to use a custom setup, uncomment below and the config will be overridden with the custom config.
#
# Example:
# custom_config = {
#   InstanceType                = "m6i.2xlarge"
#   SSDTierNumDisks             = 2
#   SSDTierDiskSizeinGB         = 511
#   SSDTierDiskType             = "gp3"
#   SSDTierDiskIops             = 3000
#   SSDTierDiskThroughputinMBps = 200
#   HDDTierNumDisks             = 2
#   HDDTierDiskSizeinGB         = 512
#   HDDTierDiskType             = "gp3"
#   HDDTierDiskIops             = 3000
#   HDDTierDiskThroughputinMBps = 200
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
