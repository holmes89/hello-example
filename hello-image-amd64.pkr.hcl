variable "aws_access_key" {
  type    = string
}

variable "aws_secret_key" {
  type    = string
}

variable "git_sha" {
  type    = string
  default = "UNKNOWN"
}

source "amazon-ebs" "hello" {
  access_key    = "${var.aws_access_key}"
  ami_name      = "hello-${var.git_sha}-amd64"
  instance_type = "t3.nano"
  region        = "us-east-2"
  secret_key    = "${var.aws_secret_key}"
  source_ami    = "ami-0a91cd140a1fc148a"
  ssh_username  = "ubuntu"
}

build {
  sources = ["source.amazon-ebs.hello"]

  provisioner "file" {
    destination = "/home/ubuntu/hello"
    source      = "hello_linux_amd64"
  }

  post-processor "manifest" {
    output = "manifest.json"
    strip_path = true
    custom_data = {
      sha = "${var.git_sha}"
    }
  }
}

