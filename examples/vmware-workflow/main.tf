provider "cohesity" {
  cluster_vip      = "10.10.10.10"
  cluster_username = "admin"
  cluster_password = "admin"
  cluster_domain   = "LOCAL"
}

resource "cohesity_source_vmware" "source1" {
  endpoint                  = "vc.cohesity.com"
  username                  = "administrator"
  password                  = "password"
  vmware_type               = "VCenter"
  cap_streams_per_datastore = true
  number_of_streams         = 5
  enable_latency_throttling = true
  new_task_latency          = 110
  active_task_latency       = 120

}

resource "time_sleep" "wait_1_minutes" {
  depends_on      = [cohesity_source_vmware.source1]
  create_duration = "60s"
}

resource "cohesity_job_vmware" "job1" {
  depends_on        = [time_sleep.wait_1_minutes]
  name              = "terraform-job1"
  protection_source = "vc.cohesity.com"
  policy            = "Bronze"
  storage_domain    = "DefaultStorageDomain"
  delete_snapshots  = true
  full_sla          = 200
  incremental_sla   = 140
  qos_type          = "BackupSSD"
  priority          = "Low"
  include           = ["chandra-pwsh01"]
}

resource "time_sleep" "wait_1_minutes2" {
  depends_on      = [cohesity_source_vmware.source1]
  create_duration = "60s"
}

resource "cohesity_job_run" "terraform-job1" {
  depends_on = [time_sleep.wait_1_minutes2]
  name       = "terraform-job1"
  state      = "start"
  timestamp  = "00:00"
}
