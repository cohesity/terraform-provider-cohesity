################################################################################
# Security Group Creation
################################################################################

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
  }
}

provider "aws" {
  region = var.region

  # Conditionally assume role if iam_role_arn is not empty
  dynamic "assume_role" {
    for_each = var.iam_role_arn != "" ? [1] : []
    content {
      role_arn = var.iam_role_arn
    }
  }
}

locals {
  parsed_tags = { for tag in var.tags : split(":", tag)[0] => split(":", tag)[1] }
  rules       = jsondecode(file(var.rules_file)).rules
}

resource "aws_security_group" "cohesity_sg" {
  name        = "${var.resource_name_prefix}-sg"
  description = "Security Group created for Cohesity Cluster"
  vpc_id      = var.vpc_id

  dynamic "ingress" {
    for_each = [for rule in local.rules : rule if !lookup(rule, "egress", false)]
    content {
      from_port   = ingress.value.from_port
      to_port     = ingress.value.to_port
      protocol    = ingress.value.protocol
      description = lookup(ingress.value, "description", null)
      cidr_blocks = ingress.value.cidr_blocks
    }
  }

  dynamic "egress" {
    for_each = [for rule in local.rules : rule if lookup(rule, "egress", false)]
    content {
      from_port   = egress.value.from_port
      to_port     = egress.value.to_port
      protocol    = egress.value.protocol
      description = lookup(egress.value, "description", null)
      cidr_blocks = egress.value.cidr_blocks
    }
  }

  tags = merge(local.parsed_tags, {
    Name = "${var.resource_name_prefix}-sg"
  })
}
