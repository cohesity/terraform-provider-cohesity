###############################################################################
# Authentication Variables
###############################################################################

# Supported AWS Authentication Methods:
#
# 1. AWS Access Key and Secret Key
#    - Set environment variables:
#        export AWS_ACCESS_KEY_ID=your_access_key
#        export AWS_SECRET_ACCESS_KEY=your_secret_key
#    - Or configure in ~/.aws/credentials and ~/.aws/config
#    - The IAM user identity must have permissions to create/modify the
#      required AWS resources.
#
# 2. IAM Role via EC2 Instance Metadata (Instance Profile)
#    - Use when running Terraform from an EC2 instance with an attached IAM
#      Role.
#    - No manual configuration needed. Terraform will use the instance
#      profile automatically by querying the instance metadata endpoint
#      (http://169.254.169.254).
#    - The attached IAM Role must have permissions to create/modify the
#      required AWS resources.
#    - Note: This method only works when running inside an EC2 instance.
#
# 3. AWS CLI Named Profile (for managing multiple AWS accounts/roles)
#    - Set up AWS CLI by running: aws configure --profile myprofile
#    - This creates entries in ~/.aws/credentials and ~/.aws/config.
#    - Use it in Terraform by setting the AWS_PROFILE environment variable,
#      or by setting the 'profile' variable (see below).
#    - Example:
#        export AWS_PROFILE=myprofile
#    - Or in provider block:
#        provider "aws" {
#          profile = "myprofile"
#          region  = "us-west-2"
#        }
#    - The profile must have permissions to create/modify the required AWS
#      resources.
#
# 4. AssumeRole via STS (Cross-account or privilege separation)
#    - Set iam_role_arn (see below) to the role you want Terraform to assume.
#    - You must have a base identity from one of the above methods with
#      sts:AssumeRole permission for the target role.
#    - Example:
#        iam_role_arn = "arn:aws:iam::123456789012:role/DeploymentRole"
#    - The assumed role must have permissions to create/modify the required
#      AWS resources.
#
# Permissions Required:
#   - The role or user used by Terraform must have permissions to create,
#     modify, and delete AWS resources required for deployment (such as EC2,
#     VPC, Security Groups, EBS, etc.).
#   - For AssumeRole, the base identity must have sts:AssumeRole for the
#     target role specified in iam_role_arn.
#   - The full set of required permissions will be provided separately in
#     documentation.
#
# Note: If running outside EC2, instance profile authentication will not work.
# Use access keys, named profile, or AssumeRole.

# (Optional) IAM Role ARN to assume for deployment (for authentication method 4).
# Example: "arn:aws:iam::123456789012:role/my-terraform-role"
# Only set this if you are using the AssumeRole authentication method above.
iam_role_arn = ""

# (Optional) AWS CLI named profile to use for authentication (method 3).
# Example: "myprofile" (must exist in ~/.aws/credentials or ~/.aws/config)
# Leave blank to use the default profile or other authentication methods.
profile = ""

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
