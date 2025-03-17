################################################################################
# Outputs
################################################################################

output "instance_info" {
  value = {
    for i in range(length(aws_instance.vm)) :
    aws_instance.vm[i].id => {
      private_ip = aws_instance.vm[i].private_ip
      public_ip  = var.attach_public_ip ? aws_instance.vm[i].public_ip : null
    }
  }
}
