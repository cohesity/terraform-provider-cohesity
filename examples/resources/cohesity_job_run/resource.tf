resource "cohesity_job_run" "terraform-job1" {
  depends_on = [time_sleep.wait_1_minutes2]
  name       = "terraform-job1"
  state      = "start"
  timestamp  = "00:00"
}
