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

variable "smtp_server" {
  description = "Server is the IP address of Network Management System."
  type        = string
}

variable "port" {
  description = "Specifies the SMTP port. Usually 465 or 587. For authenticated connection, it is generally 587."
  type        = number
}

variable "sender_email_address" {
  description = "This is used for setting Sender field in SMTP header. This has to be in valid email format, and could be different from username, and it's a required field if the username is not of valid email address."
  type        = string
}

variable "username" {
  description = "Specifies the username which will be used to connect to the SMTP server. If username is not specified, then it would imply that SMTP server is set up for unauthenticated access."
  type        = string
}

variable "password" {
  description = "Specifies the password of the SMTP user. This is required if username is specified in the request."
  type        = string
}

variable "use_ssl" {
  description = "This is set to true when the SMTP server uses SSL/TLS without supporting STARTTLS. Typically, this is used for port 465."
  type        = bool
  default     = false
}