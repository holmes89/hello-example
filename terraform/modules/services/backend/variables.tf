variable "ami" {
    description = "The image to load onto the server"
    type = string
}

variable "server_name" {
    description = "Name given to the server"
    type = string
}

variable "instance_type" {
  description = "The type of EC2 Instances to run (e.g. t2.micro)"
  type        = string
}