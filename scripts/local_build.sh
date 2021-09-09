#!/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
cd $SCRIPT_DIR/..
git clean -xdf
cd ./internal/pkg/ba_proto
go mod init ba_proto
go mod tidy
cd ../virtual_industrial_portal
go mod init virtual_industrial_portal
go mod tidy
cd ../proto_helper
go mod init proto_helper
go mod tidy
cd ../../..
go mod init internal/app/main
go mod tidy
go mod edit -replace ba_proto=./internal/pkg/ba_proto
go mod edit -replace virtual_industrial_portal=./internal/pkg/virtual_industrial_portal
go mod edit -replace proto_helper=./internal/pkg/proto_helper
go get ba_proto
go get virtual_industrial_portal
go get proto_helper
cd ./internal/app
go build -o virtual-industrial-portal-app
mv ./virtual-industrial-portal-app ../../virtual-industrial-portal-app