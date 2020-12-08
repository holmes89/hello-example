#!/bin/bash
set -e

export GOPATH=$PWD/gopath
export PATH=$PWD/gopath/bin:$PATH

ls
cd hello-example

go get -t ./...
go test ./...