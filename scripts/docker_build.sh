#!/bin/bash
cd /virtual-fleet/tmp/internal/pkg/ba_proto
go mod init ba_proto
go mod tidy
cd ../virtual_fleet
go mod init virtual_fleet
go mod tidy
cd ../proto_helper
go mod init proto_helper
go mod tidy
cd ../../..
go mod init internal/app/main
go mod tidy
go mod edit -replace ba_proto=./internal/pkg/ba_proto
go mod edit -replace virtual_fleet=./internal/pkg/virtual_fleet
go mod edit -replace proto_helper=./internal/pkg/proto_helper
go get ba_proto
go get virtual_fleet
go get proto_helper
cd ./internal/app
go build -o virtual-fleet-app
mv ./virtual-fleet-app ../../virtual-fleet-app
mv /virtual-fleet/tmp/scenarios /virtual-fleet/