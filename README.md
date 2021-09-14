# Virtual industrial portal
This project serves as testing base for BA daemon. Virtual industrial portal implements [industrial portal protocol](https://docs.google.com/document/d/1sjIE4_c9NrQCpUvlgOwejVMWf6U-QSh_9qobpMqOIRU/edit) using MQTT and it replays mission scenarios to connected cars, also logging states of car.

## Requirements
- [golang](https://golang.org/)
- [protobuff compiler v3.17.3](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3) 

## Arguments
- --broker-ip=<ipv4> -ip address of broker
- --broker-port=<port> - port of broker

## Tests


## Build and run project localy
Run build script from project folder:
```
bash ./scripts/local_build.sh
```
Run the app: 
```
./virtual-industrial-portal
```

## Compiling proto
You must have [protobuf-compiler](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3) version 3.17.3 and exported go path:
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
```
In project folder run:
```
protoc -I=./autonomy-host-protocol --go_out=./autonomy-host-protocol autonomy-host-protocol/*.proto
```

## Build and run docker image
Build the image using:
```
docker build --tag virtual-industrial-portal .
```
You can list docker images using:
```
docker images
```
and find image id of your docker container. Run the image using:
```
docker run -ti --rm virtual-industrial-portal /virtual-industrial-portal/virtual-industrial-portal --broker-ip=<MQTT broker ipv4> --broker-port=<MQTT broker port>
```