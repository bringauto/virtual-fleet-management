# Virtual Fleet

This project serves as testing base for BA daemon. Virtual fleet
implements [industrial portal protocol](https://docs.google.com/document/d/1sjIE4_c9NrQCpUvlgOwejVMWf6U-QSh_9qobpMqOIRU/edit)
using MQTT and it replays mission scenarios to connected cars, also logging states of car.

## Behavior

The application reads all scenarios from the directory specified in the configuration.
Each scenario is a JSON file that describes a sequence of missions for a car.
The application creates a simulation for each scenario.

The application creates an HTTP client. It checks if the fleet management HTTP API is available. 
If it's not, it tries again after some time. The sleep time and number of retries are constants in the code. The sleep time is increasing with each retry.

Then it creates routes and stops from the scenario files. If they already exist and are the same, this step is skipped.
If they exist with different data (e.g. different coordinates), the application logs an error and stops.

The application then starts monitoring the cars. 
It periodically retrieves the list of cars from the fleet management HTTP API and checks if each car is communicating.
If a car is communicating, there is available scenario for it and it's not already active, the application starts the simulation for that car in a new goroutine.

Each simulation runs the missions in its scenario in sequence.
If the loop configuration value is true, the simulation loops the missions indefinitely.
Otherwise, it runs the missions once and then stops once they are all done.

## Requirements

- [golang](https://golang.org/)
- [protobuff compiler v3.17.3](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3)

## Arguments

- --broker-ip=\<ipv4> - IP address of the MQTT broker
- --broker-port=\<port> - port of the MQTT broker

## Tests

## Scenarios

Files are distributed into folders depending the company and car name they will be used for.

For example scenario for company `BringAuto` with car name `CAR1` will be stored in `/bringauto/car1/scenario.json`.
Each car folder can contain multiple scenario files, but right now one scenario per car is supported, first correct file
will be run and other files will be ignored

All scenarios from the company directory will be played in parallel. The application gets the company folder from a
config.

Example scenarios are stored in [virtual-fleet scenarios folder](resources/scenarios/) in json format.

### Directory structure

```
scenarios/
├── companyA/    <-- company folder, passed to the app in config
│   ├── car1/
│   │   └── scenario.json
│   └── car2/
│   │   └── scenario.json
└── companyB/
    ├── car1/
    │   └── scenario.json
    └── car2/
        └── scenario.json
```

### JSON format

JSON files contain the information about the map file that is used for missions (name or path to it), starting station
and list of missions.

Starting station defines in which station the simulation will start.
Order with starting_station is sent to the car as an initial mission. Once this initial mission is completed, first
mission is started.

Each mission contains a delay, the mission's name, and a list of stops.
Delay sets the time, after which the mission is started. This time is counted since the previous mission is started.
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

This scenario will create the order ["London National Theatre"], then play the
mission ["London National Theatre", "Cross Station" ] from map London.osm from timestamp 0 to 150 (calculated from
reaching the starting station of given scenario)
and after that time interval it will switch to second mission  [ "Oasis Academy", "London Waterloo" ]

## Build and run project locally

Run the build script from the project folder:

```
bash build.sh
```

Run the app:

```
./virtual-fleet-management -config <config.json>
```

### Config values

* **host** - address of the server where the virtual fleet will be running
* **api-key** - API key for the Fleet Management HTTP API
* **scenario-dir** - path to the directory with scenarios. `<scenario_dir>/<car_name>/scenario.json`
    * common approach is to name the scenario directory as company
    * inside the `scenario-dir`, there can be directories for multiple cars, all of them can be played at same time.
* **log-path** - path to the log directory
* **loop** - optional argument, if set to `true` the scenarios will be looped infinitely. If `false` each scenario will
  be played only once.

> Arguments `api-key` and `scenario-dir` can be overridden by the command line arguments.

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

