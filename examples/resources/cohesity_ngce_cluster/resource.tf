resource "cohesity_ngce_cluster" "cluster" {
  name                     = "cluster_name"
  node_ips                 = ["0.0.0.0", "0.0.0.1"]
  hostname                 = "hostname"
  subnet_gateway           = "subnetwork_gateway"
  subnet_mask              = "subnet_mask"
  dns_server_ips           = "dns-server-ips"
  ntp_servers              = "ntp-servers"
  domain_names             = "LOCAL"
  metadata_fault_tolerance = 1
  add_apps                 = false
}