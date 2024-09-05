credential_file = "GCPKey.json"
project         = "cohesity-engineering-internal"
region          = "us-west1"
zone            = "us-west1-a"
cluster_name    = "cohesity-cluster"
num_instances   = 3
machine_type    = "n1-standard-32"
network         = "cohesity-google-subnet"
subnetwork      = "cohesity-google-subnet-us-west1"
labels = {
  "key" = "value"
}
disk_image              = "cohesity-gcp-7-2-release-20240711-0083b4c9-cluster-image"
allot_external_ip       = true
dns-server-ips          = "10.2.38.16"
ntp-servers             = "time.nist.gov"
domain-names            = "eng.cohesity.com"
add_gcp_external_target = true
gcp_external_target = {
  project_id                   = ""
  client_email_id              = ""
  client_private_key_file_path = ""
  bucket_name                  = ""
  tier_type                    = "Google-standard-storage"
}
cohesity_username = "admin"
cohesity_password = "admin"
support_password  = "Cohe$1tyCohe$1ty"
