#!/bin/bash

set -e

test -z "$1" && go run generateTestData/generate.go -n=100000 || go run generateTestData/generate.go -n=$1
printf "\n----------------------------------------------------------------\n"
go test -v ./...
printf "\n----------------------------------------------------------------\n"
go run showData.go
