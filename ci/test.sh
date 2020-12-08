#!/bin/bash
set -e

export GOPATH=$PWD/gopath
export PATH=$PWD/gopath/bin:$PATH

cd hello-code

echo "Fetching Deps..."
go get -t ./...

echo "Testing..."
go test ./...