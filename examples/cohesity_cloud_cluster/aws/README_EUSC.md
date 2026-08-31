# AWS European Sovereign Cloud (EUSC) — NGCE cluster provisioning

Use these examples to deploy a Cohesity NextGen Cloud Edition (NGCE) cluster in
**AWS European Sovereign Cloud** (`eusc-de-east-1`, partition `aws-eusc`,
dnsSuffix `amazonaws.eu`).

Commercial AWS examples continue to work unchanged: leave `aws_custom_endpoints`
unset (default `null`).

## Why custom endpoints?

The HashiCorp AWS provider defaults to `*.amazonaws.com`. EUSC APIs live on
`*.amazonaws.eu`. Without `aws_custom_endpoints`, Terraform plan/apply fails
against the commercial endpoints.

## Recommended instance / disk config

| Setting | Value |
|---|---|
| InstanceType | `m6i.2xlarge` (8 vCPU, 32 GiB) |
| SSD tier | `gp3`, 2 × 100 GB, 3000 IOPS, 125 MB/s |
| HDD tier | `gp3`, 1 × 512 GB, 3000 IOPS, 125 MB/s |

Avoid 4-vCPU shapes (e.g. `r6i.xlarge`) for CAD archival labs: NGCE autoscaler
can derive `bridge_blob_store_downtier_cf_to_cloud_num_deque_per_batch=1` and
crash `bridge_exec`.

## Steps

1. Create security group:

```bash
cd security_group_create
cp terraform.tfvars.eusc-example terraform.tfvars
# edit vpc_id / prefix / tags
terraform init && terraform apply
```

2. Create cluster:

```bash
cd ../cluster_create
cp terraform.tfvars.eusc-example terraform.tfvars
# edit subnet_id, security_group_ids, image_id, cluster_name
terraform init && terraform apply
```

3. Ensure CLI / env use the same partition (MFA session recommended), e.g.
   STS `https://sts.eusc-de-east-1.amazonaws.eu`.
