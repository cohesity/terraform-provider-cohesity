resource "cohesity_configure_snmp" "snmp" {
  server = "0.0.0.0"
  snmp_version = "V2C"
  agent_port = 161
  trap_port = 162

  trap_user {
    user_name     = "trap_user"
    auth_password = "auth_password"
    auth_protocol = "auth_protocol"
    priv_password = "priv_password"
    priv_protocol = "priv_protocol"
  }

  read_user {
    user_name     = "read_user"
    auth_password = "auth_password"
    auth_protocol = "auth_protocol"
    priv_password = "priv_password"
    priv_protocol = "priv_protocol"
  }
}