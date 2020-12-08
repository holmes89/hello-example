#!/bin/bash
set -e

export GOPATH=$PWD/gopath
export PATH=$PWD/gopath/bin:$PATH

cd hello-code

go get -t ./...
go test ./...