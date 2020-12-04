provider "aws" {
  region = "us-east-2"
}

resource "aws_security_group" "hello_traffic" {
  name        = "hello_traffic"
  description = "Allow all inbound traffic"

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

}

resource "aws_instance" "hello" {
  ami           = "ami-07eaad496802df3b9"
  instance_type = "t3.micro"
  security_groups = ["hello_traffic", "default"]
  key_name = "debug"
  user_data = <<-EOF
              #!/bin/bash
              sudo chmod +x /home/ubuntu/hello
              sudo /home/ubuntu/hello
                EOF
}

output "hello_dns" {
  value = aws_instance.hello.public_dns
}