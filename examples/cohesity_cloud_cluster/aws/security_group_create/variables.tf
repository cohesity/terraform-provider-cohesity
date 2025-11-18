###############################################################################
# Authentication variables
###############################################################################

variable "iam_role_arn" {
  description = <<EOT
(Optional) The ARN of the IAM Role to assume for AWS authentication.
If specified, Terraform will use this role for all AWS API calls instead of the default credentials.
Leave blank to use the default AWS authentication mechanism (environment variables, profile, etc).
EOT
  type        = string
  default     = ""
}

###############################################################################
# Security Group Variables
###############################################################################

variable "region" {
  type        = string
  description = "AWS Region to deploy to"
}

variable "vpc_id" {
  type        = string
  description = "ID of the VPC where the security group will be created"
}

variable "resource_name_prefix" {
  description = <<EOT
Prefix for the names of all created resources.
This value is mandatory. It must be greater than 0 and less than or equal to 30 characters long.
It must contain only lowercase letters, numbers, and hyphens, must start with a lowercase letter, and cannot end with a hyphen.
EOT

  type = string

  validation {
    condition = (
      length(var.resource_name_prefix) > 0 &&
      length(var.resource_name_prefix) <= 30 &&
      can(regex("^[a-z][a-z0-9-]*[a-z0-9]$", var.resource_name_prefix))
    )
    error_message = "resource_name_prefix is mandatory, must be 1-30 chars, start with a lowercase letter, contain only lowercase letters, numbers, and hyphens, and not end with a hyphen."
  }
}

variable "tags" {
  type        = list(string)
  description = "List of tags in the format ['key:value', 'key:value']"
  default     = []
}

variable "rules_file" {
  type        = string
  description = "Path to the JSON rules file for security group ingress rules"
  default     = "./rules.json"
}
