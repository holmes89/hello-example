#!/bin/bash
set -e

echo "Deploying $BUILD_TYPE Image"
apk add --no-cache jq
export AMI=$(jq -r '.builds[0].artifact_id|split(":")[1]' ./packer-manifest-$BUILD_TYPE/manifest.json)
cd hello-code/terraform/prod/services/backend/$BUILD_TYPE
terraform init
terraform apply -auto-approve -var "ami=$AMI"