provider "aws" {
  region = "us-east-2"
}

variable "hello_ami" {
  type = string
}

terraform {  
  backend "s3" {    
    bucket         = "hello-example-state"
    key            = "global/s3/terraform.tfstate"
    region         = "us-east-2"   
    dynamodb_table = "hello-example-locks" 
    encrypt        = true  
  }
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
  ami           = var.hello_ami
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