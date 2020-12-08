variable "server_name" {
    description = "Name given to the server"
    type = string
    default     = "hello-amd64"
}

variable "ami" {
    description = "The image to load onto the server"
    type = string
}
