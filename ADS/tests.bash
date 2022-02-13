#!/bin/bash

set -e

go run generateTestData/generate.go -n=100
printf "\n----------------------------------------------------------------\n"
go test -v ./...
printf "\n----------------------------------------------------------------\n"
go run showData.go
