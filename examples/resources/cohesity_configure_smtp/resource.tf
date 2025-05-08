resource "cohesity_configure_smtp" "smtp" {
  smtp_server = "0.0.0.0"
  port = 123
  sender_email_address = "example@example.com"
  username = "username"
  password = "password"
  use_ssl = false
}