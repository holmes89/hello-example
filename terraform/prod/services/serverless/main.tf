provider "aws" {
  region = "us-east-2"
}

terraform {  
  backend "s3" {    
    bucket         = "hello-example-state"
    key            = "prod/services/serverless/terraform.tfstate"
    region         = "us-east-2"   
    dynamodb_table = "hello-example-locks" 
    encrypt        = true  
  }
}

module "serverless" {
    source = "../../../modules/services/serverless"
    function_name = "hello-serverless"
    source_bucket = "hello-example-lambda"
    source_file = "${var.file_name}"
}