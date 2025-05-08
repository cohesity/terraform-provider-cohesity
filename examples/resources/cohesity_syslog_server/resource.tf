resource "cohesity_syslog_server" "syslog" {
  name        = "syslog"
  ip_address = "0.0.0.0"
  port       = 514
  protocol   = "udp"
  ca_certificate = "*****"
  program_names = ["sshd"]
  is_enabled = false
}