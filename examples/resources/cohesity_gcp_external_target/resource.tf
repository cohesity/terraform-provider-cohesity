resource "cohesity_gcp_external_target" "gcp_external_target" {
  client_private_key_file_path = "path/to/the/private/key/file"
  bucket_name                  = "bucket"
  project_id                   = "example-project-id"
  client_email                 = "client-email-id"
  tier_type                    = "Google-standard-storage"
}
