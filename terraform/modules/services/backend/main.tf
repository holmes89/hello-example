
resource "aws_security_group" "instance" {
  name        = "${var.server_name}-instance"
  description = "Allow all inbound traffic"

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

}

resource "aws_instance" "hello" {
  ami           = var.ami
  instance_type = var.instance_type
  security_groups = ["${var.server_name}-instance", "default"]
  key_name = "debug"
  user_data = <<-EOF
              #!/bin/bash
              sudo chmod +x /home/ubuntu/hello
              sudo /home/ubuntu/hello
                EOF
  
  lifecycle {
    create_before_destroy = true
  }                
}
