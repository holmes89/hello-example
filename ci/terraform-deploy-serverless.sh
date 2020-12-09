#!/bin/bash
set -e

echo "Deploying Lambda"
cd hello-code/terraform/prod/services/serverless

export VERSION=$(git rev-parse --short HEAD)
export FILENAME="releases/server-$VERSION.zip"

echo "File $FILENAME"

terraform init
terraform apply -auto-approve -var "file_name=$FILENAME"