# Virtual Fleet
This project serves as testing base for BA daemon. Virtual fleet implements [industrial portal protocol](https://docs.google.com/document/d/1sjIE4_c9NrQCpUvlgOwejVMWf6U-QSh_9qobpMqOIRU/edit) using MQTT and it replays mission scenarios to connected cars, also logging states of car.

## Requirements
- [golang](https://golang.org/)
- [protobuff compiler v3.17.3](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3)

## Arguments
- --broker-ip=\<ipv4> - IP address of the MQTT broker
- --broker-port=\<port> - port of the MQTT broker

## Tests

## Scenarios
Scenarios are stored in [virtual-fleet scenarios folder](resources/scenarios/) in json format. Files are distributed  into folders depending on topics they will be used on.
For example scenarios for topic bringauto/kralovopolska/car1 will be stored in virtual-fleet/scenarios/bringauto/kralovopolska/car1/. Each car folder can contain multiple scenario files, but right now one scenario per car is supported, first correct file will be run and other files will be ignored

### JSON format
Json files contain information about map file that is used for missions (name or path to it), starting station, and list of missions.

Starting station defines in which station the simulation will start. 
Order with starting_station is sent to the car as an initial mission. Once this initial mission is completed, first mission is started.

Each mission contains delay, name and list of stops. Delay informs the virtual fleet after how long period after establishing first connection of given car will be the mission started.

Example:

```
{
    "map": "London.osm",
    "starting_station": "London National Theatre",
    "missions": [
        {
            "delay_seconds": 0,
            "name": "mission1",
            "stops": [
                {
                    "name": "London National Theatre"
                },
                {
                    "name": "Cross Station"
                }
            ],
            "route": "Short"
        },
        {
            "delay_seconds": 150,
            "name": "mission2",
            "stops": [
                {
                    "name": "Oasis Academy"
                },
                {
                    "name": "London Waterloo"
                }
            ],
            "route": "Long"
        }
    ],
    "routes": [
        {
            "name": "Short",
            "stations": [
                {
                    "name": "London National Theatre",
                    "position": {
                        "latitude": 51.50719991926,
                        "longitude": -0.11572123647,
                        "altitude": 0
                    }
                },
                {
                    "name": "Cross Station",
                    "position": {
                        "latitude": 51.50847034843,
                        "longitude": -0.12557298213,
                        "altitude": 0
                    }
                },
            ]
        },
        {
            "name": "Long",
            "stations": [
                {
                {
                    "name": "London Waterloo",
                    "position": {
                        "latitude": 51.50423320901,
                        "longitude": -0.11283786178,
                        "altitude": 0
                    }
                },
                {
                    "name": "London National Theatre",
                    "position": {
                        "latitude": 51.50719991926,
                        "longitude": -0.11572123647,
                        "altitude": 0
                    }
                },
                {
                    "name": "Oasis Academy",
                    "position": {
                        "latitude": 51.50031457906,
                        "longitude": -0.11135928467,
                        "altitude": 0
                    }
                }
            ]
        }
    ]
}
```
This scenario will place order ["London National Theatre"], then play mission ["London National Theatre", "Cross Station" ] from map London.osm from timestamp 0 to 150 (calculated from first connection with given car)
and after that time interval it will switch to second mission  [ "Oasis Academy", "London Waterloo" ]



## Build and run project localy
Run build script from project folder:
```
bash build.sh
```
Run the app:
```
./virtual-fleet-management -config <config.json>
```

### Config values

* **host** - address of the server where the virtual fleet will be running
* **api-key** - api key for the Fleet Management HTTP API
* **scenario-dir** - path to the directory with scenarios. `<scenario_dir>/<car_name>/scenario.json`
  * common approach is to name the scenario directory as company
* **log-path** - path to the log directory
* **loop** - optional argument, if set the scenarios will be looped infinitely

## Build and run docker image
Build the image using:
```
docker build --tag virtual-fleet .
```
You can list docker images using:
```
docker images
```
and find image id of your docker container. Run the image using:
```
docker run -ti --rm virtual-fleet /virtual-fleet/virtual-fleet-app --broker-ip=<MQTT broker ipv4> --broker-port=<MQTT broker port> scenario-dir=<path to scenario dir>
```