provider "aws" {
  region = "us-west-1"
  access_key = "access_key"
  secret_key = "secret_key"
}
provider "cohesity" {
        cluster_vip = "10.10.10.10"
        cluster_username = "admin"
        cluster_password = "admin"
        cluster_domain = "LOCAL"
}
resource "aws_instance" "CohesityNode" {
  ami           = "ami-xx"
  instance_type = "m4.4xlarge"
  availability_zone = "us-west-1c"
  subnet_id = "subnet-xx"
  private_ip = "10.101.33.177"
  vpc_security_group_ids = ["sg-xx"]
  ebs_block_device {
    device_name="/dev/xvdb"
    volume_type="gp2"
    volume_size=512
    }
  ebs_block_device {
   device_name= "/dev/xvdc"
   volume_type= "st1"
   volume_size= 1000
   }
  tags = {
    Name = "Terraform-demo-Node1"
    environment = "QA"
    user = "anvesh"
    squad = "API_Integrations"
    location = "sanjose"
    weekend = "snooze"
    expiry = "06/26/2020"
    email = "email@cohesity.com"
    non-business-hrs = "snooze"
  }
}
resource "time_sleep" "wait_10_minutes" {
  depends_on = [aws_instance.CohesityNode]
  create_duration = "600s"
}
resource "cohesity_cloud_edition_cluster" "TestCluster"{
            cluster_name = "TerraformCloudEdition"
            dns_servers = ["10.10.10.1"]
            ntp_servers = ["time.google.com"]
            domain_names = ["eng.cohesity.com"]
            cluster_subnet_mask = "255.255.240.0"
            cluster_gateway = "10.10.10.1"
            enable_encryption = true
            license_key = "0000-000-000"
            enable_fips_mode = true
            encryption_keys_rotation_period = 1
            metadata_fault_tolerance = 0
            node_ips = ["10.10.10.10"]
            depends_on = [time_sleep.wait_10_minutes]
}
