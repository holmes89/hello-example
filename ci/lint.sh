#!/bin/bash
set -e

export GOPATH=$PWD/gopath
export PATH=$PWD/gopath/bin:$PATH

ls 
cd hello-example

go get -t ./...
go vet ./...
go get -u golang.org/x/lint/golint
go install golang.org/x/lint/golint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0
curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.5.0

golint ./...
golangci-lint run
gosec ./...