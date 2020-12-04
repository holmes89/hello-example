variable "aws_access_key" {
  type    = string
  default = "${var.AWS_ACCESS_KEY_ID}"
}

variable "aws_secret_key" {
  type    = string
  default = "${var.AWS_SECRET_ACCESS_KEY}"
}

variable "git_sha" {
  type    = string
  default = "none"
}

source "amazon-ebs" "hello" {
  access_key    = "{{user `aws_access_key`}}"
  ami_name      = "hello-{{git_sha}}"
  instance_type = "t3.nano"
  region        = "us-east-2"
  secret_key    = "{{user `aws_secret_key`}}"
  source_ami    = "ami-0a91cd140a1fc148a"
  ssh_username  = "ubuntu"
}

build {
  sources = ["source.amazon-ebs.hello"]

  provisioner "file" {
    destination = "/home/ubuntu/"
    source      = "hello"
  }
}
