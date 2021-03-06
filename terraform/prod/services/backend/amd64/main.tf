provider "aws" {
  region = "us-east-2"
}

terraform {  
  backend "s3" {    
    bucket         = "hello-example-state"
    key            = "prod/services/backend/amd64/terraform.tfstate"
    region         = "us-east-2"   
    dynamodb_table = "hello-example-locks" 
    encrypt        = true  
  }
}

module "backend" {
    source = "../../../../modules/services/backend"

    server_name = var.server_name
    instance_type = "t3a.nano"
    ami = var.ami
}