variable "function_name" {
    description = "Name given to the lamda"
    type = string
}

variable "source_bucket" {
    description = "Bucket holding binary"
    type = string
}

variable "source_file" {
    description = "Zip of the binary"
    type = string
}

variable "environment" {
    description = "name of deployment (test, api for production)"
    type = string
    default     = "api"
}