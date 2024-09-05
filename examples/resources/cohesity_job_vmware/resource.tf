resource "cohesity_job_vmware" "job1" {
  depends_on        = [time_sleep.wait_1_minutes]
  name              = "terraform-job1"
  protection_source = "vc.cohesity.com"
  policy            = "Bronze"
  storage_domain    = "DefaultStorageDomain"
  include           = ["vm-name"]
}