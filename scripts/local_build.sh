#!/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
cd $SCRIPT_DIR/..
git clean -xdf
cd ./internal/pkg/ba_proto
go mod init ba_proto
go mod tidy
cd ../virtual_fleet
go mod init virtual_fleet
go mod tidy
cd ../../..
go mod init internal/app/main
go mod tidy
go mod edit -replace ba_proto=./internal/pkg/ba_proto
go mod edit -replace virtual_fleet=./internal/pkg/virtual_fleet
go get ba_proto
go get virtual_fleet
cd ./internal/app
go build -o virtual-fleet-app
mv ./virtual-fleet-app ../../virtual-fleet-app