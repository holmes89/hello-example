#!/bin/bash
set -e

echo "Building $BUILD_TYPE Image"
cp gh-release/hello_linux_$BUILD_TYPE hello-code/.
cd hello-code
packer build -force -machine-readable -var "git_sha=$(git rev-parse --short HEAD)" -var "aws_access_key=$AWS_ACCESS_KEY_ID" -var "aws_secret_key=$AWS_SECRET_ACCESS_KEY" -var "aws_secret_key=$BUILD_TYPE" hello-image-$BUILD_TYPE.pkr.hcl
cp manifest.json ../packer-manifest-$BUILD_TYPE