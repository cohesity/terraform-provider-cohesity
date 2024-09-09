resource "cohesity_source_vmware" "source1" {
  endpoint    = "vc.xyz.com"
  username    = "administrator"
  password    = "password"
  vmware_type = "VCenter"
}
