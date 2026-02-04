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
#     for example Security Group.
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
# Example tfvars for security_group_create
###############################################################################

# AWS region for deployment.
# Example: "us-west-2", "us-east-1"
# To list available regions:
#   aws ec2 describe-regions --query "Regions[*].RegionName"
region = "replace-with-region"

# VPC ID where the security group will be created.
# Example: "vpc-xxxxxxxxxxxxxxxxx"
# To list VPCs:
#   aws ec2 describe-vpcs --query "Vpcs[*].VpcId"
vpc_id = "vpc-xxxxxxxxxxxxxxxxx"

# Prefix for naming the security group.
# Example: "cohesity"
# Restrictions:
#   - Only letters, numbers, hyphens
#   - Max 30 chars
#   - Cannot start/end with hyphen
#   - No consecutive hyphens
resource_name_prefix = "replace-with-prefix"

# Tags to apply to the security group.
# Format: ["key1:value1", "key2:value2"]
# Example: ["environment:development", "owner:your-name"]
# Note: Tags help identify and organize resources in AWS.
tags = [
  "key1:value1",
  "key2:value2"
]

# Path to the rules file (relative to this directory or absolute path).
# Example: "./rules.json"
# This file should contain the security group ingress rules in JSON format.
# If you are unsure, use the provided rules.json as a template.
rules_file = "./rules.json"
