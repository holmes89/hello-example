#!/bin/bash
set -e

export GOPATH=$PWD/gopath
export PATH=$PWD/gopath/bin:$PATH

cd hello-code

echo "Fetching Deps..."
go get -t ./...

echo "Checking format..."
test -z $(go fmt ./...)

echo "Go Vet"
go vet ./...

echo "Installing Linters"
go get -u golang.org/x/lint/golint
go install golang.org/x/lint/golint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0
curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.5.0

echo "Linting"
golint ./... -set_exit_status
golangci-lint run
gosec ./...