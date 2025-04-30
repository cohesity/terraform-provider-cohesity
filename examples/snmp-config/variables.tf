variable "cluster_vip" {
  description = "Cohesity IRIS VIP"
  type        = string
}

variable "cohesity_username" {
  description = "Cohesity IRIS username"
  type        = string
}

variable "cohesity_password" {
  description = "Cohesity IRIS user password"
  type        = string
}

variable "support_password" {
  description = "Cohesity support user password"
  type        = string
}

variable "server" {
  description = "Server is the IP address of Network Management System."
  type        = string
}

variable "snmp_version" {
  description = "SnmpVersion is the SNMP version to talk with SNMP agent. Options: V2C, V3"
  type        = string
  default     = "V2C"
}

variable "agent_port" {
  description = "AgentPort is the TCP port SNMP agent is using."
  type        = number
  default     = 161
}

variable "trap_port" {
  description = "TrapPort is the TCP port SNMP agent is using."
  type        = number
  default     = 162
}

variable "trap_user" {
  description = "SNMP Trap User Info for this cluster"
  type = object({
    user_name                   = string
    auth_protocol               = string
    auth_password               = string
    priv_protocol               = string
    priv_password               = string
  })
  default = null
}

variable "write_user" {
  description = "SNMP Write User Info for this cluster"
  type = object({
    user_name                   = string
    auth_protocol               = string
    auth_password               = string
    priv_protocol               = string
    priv_password               = string
  })
  default = null
}

variable "read_user" {
  description = "SNMP Read User Info for this cluster"
  type = object({
    user_name                   = string
    auth_protocol               = string
    auth_password               = string
    priv_protocol               = string
    priv_password               = string
  })
  default = null
}