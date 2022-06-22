# Virtual industrial portal
This project serves as testing base for BA daemon. Virtual industrial portal implements [industrial portal protocol](https://docs.google.com/document/d/1sjIE4_c9NrQCpUvlgOwejVMWf6U-QSh_9qobpMqOIRU/edit) using MQTT and it replays mission scenarios to connected cars, also logging states of car.

## Requirements
- [golang](https://golang.org/)
- [protobuff compiler v3.17.3](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3) 

## Arguments
- --broker-ip=<ipv4> -ip address of broker
- --broker-port=<port> - port of broker

## Tests

## Scenarios 
Scenarios are stored in [virtual-industrial-portal](https://gitlab.bringauto.com/bring-auto/host-platform/virtual-industrial-portal) in folder scenarios in json format. Files are distributed  into folders depending on topics they will be used on.
For example scenarios for topic bringauto/kralovopolska/car1 will be stored in virtual-industrial-portal/scenarios/bringauto/kralovopolska/car1/. Each car folder can contain multiple scenario files, but right now one scenario per car is supported, first correct file will be run and other files will be ignored

### JSON format
Json files contain information about map file that is used for missions (name or path to it) and list of missions. Each mission contain timestamp, name and list of stops. Timestamp informs the virtual industrial portal after how long period after establishing first connection of given car will be the mission started. Example:

```
{
	"map": "borsodchem.osm",
	"missions": [
		{
		 "timestamp": "0",
                 "name":"mission1",
		 "stops": [
			{"name": "Lab A-blok"},
			{"name": "Plnička"}
			]
		},
		{
	 	 "timestamp": "30",
                 "name":"mission2",
		 "stops": [
			{"name": "Vodík"},
			{"name": "Lab A-blok"}
			]
		}
	]
}
```
This scenario will play mission ["Lab A-blok", "Plnička" ] from map borsodchem.osm from timestamp 0 to 30 (calculated from fir connection with given car) and after that time interval it will switch to second mission  [ "Vodík", "Lab A-blok" ]



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
docker run -ti --rm virtual-industrial-portal /virtual-industrial-portal/virtual-industrial-portal --broker-ip=<MQTT broker ipv4> --broker-port=<MQTT broker port> scenario-dir=<path to scenario dir>
```