# Virtual industrial portal
This project serves as testing base for BA daemon. Virtual industrial portal implements [industrial portal protocol](https://docs.google.com/document/d/1sjIE4_c9NrQCpUvlgOwejVMWf6U-QSh_9qobpMqOIRU/edit) using MQTT and it replays mission scenarios to connected cars, also logging states of car.

## Requirements
- [golang](https://golang.org/)
- [protobuff compiler v3.17.3](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3) 

## Arguments
- --broker-ip=<ipv4> -ip address of broker
- --broker-port=<port> - port of broker

## Tests
TODO

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

## Docker compose
Docker compose file requires four images to work: bringauto-mosquitto, virtual-vehicle-utility, bringauto-daemon and virtual-vehicle-utility.
Parameters of each program can be edited in docker-compose file. Logs from apps can be found in [volumes folder](./docker-volumes)Compose file contains three profiles:
- all - this profile will start all docker images
- without-industrial-portal - profile will start all images without industrial portal
- without-autonomy - profile will start all images without virtual-vehicle-utility
Run docker:
docker-compose --profile <profile> up

### Docker images
- virtual-industrial-portal - model implementation of [industrial portal](https://docs.google.com/document/d/1vWT44qj_lLXfPZ30w_oSX4ZQvKihFTzvmrHMQE7gHwE/edit#heading=h.g90z0bq4rgs4) that simulates actions of user (creating car missions) and communicate with car using [mqtt](https://docs.google.com/document/d/1PmaJqZK_cUQxfLh7RTU5O7t-CTSn_aWTIL6rZT9gB6M/edit), IPv4 address in docker is 10.5.0.4, connection to broker is not secured and on port 1883
- bringauto-mosquitto - MQTT broker, listens on 10.5.0.2:1883, does not use ssl
- bringauto-daemon - connection between autonomy (virtual-vehicle-utility) and industrial portal, using CarState protocol and Industrial portal protocol, IPv4 is 10.5.0.4, with autonomy communicates on port 1536, with mqtt broker on port 1883
- virtual-vehicle-utility - program simulating autonomous driving of car on specified map, communicates with bringauto daemon on port 1536 and implements [CarState protocol](https://docs.google.com/document/d/1cW5t_ue0wQmp-InI-M2fug6mxXvLrYcfVjTRHR4et_c/edit)
