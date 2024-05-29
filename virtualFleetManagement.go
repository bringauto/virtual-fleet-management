package main

import (
	"encoding/json"
	"flag"
	"fmt"
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"virtual-fleet-management/pkg/http"
	"virtual-fleet-management/pkg/scenario"
	"virtual-fleet-management/pkg/simulation"
)

const sleepTime = 10

// TODO move to GitHub
func main() {
	//hostIp, apiKey, logPath, scenariosPath, loop := parseFlags()
	config := parseFlags()

	setUpLogger(config.LogPath)
	setSignalHandler()

	allScenarios := scenario.GetAllScenariosFromDir(config.ScenariosPath)
	client := http.CreateClient(config.HostIp, config.ApiKey)

	simulations := createSimulations(allScenarios, config.Loop, client)

	monitorAndStartNewCars(client, simulations)
}

type Config struct {
	HostIp        string `json:"host"`
	ApiKey        string `json:"api-key"`
	LogPath       string `json:"log-path"`
	ScenariosPath string `json:"scenario-dir"`
	Loop          bool   `json:"loop"`
}

func parseFlags() (config Config) {
	configFile := flag.String("config", "", "Path to JSON configuration file")
	help := flag.Bool("help", false, "Show help")
	apiKey := flag.String("api-key", "", "API key for the fleet management server. Optional, will override config value")
	flag.Parse()

	if *help {
		print(
			"Run: ./virtual-fleet-management -config config.json\n",
			"Options:\n")
		flag.PrintDefaults()
		print(
			"Config values:\n",
			"\thost: IP address of the Fleet Management HTTP API\n",
			"\tapi-key: API key for the Fleet Management HTTP API\n",
			"\tlog-path: Path to the log file\n",
			"\tscenario-dir: Path to the directory containing the scenario files\n",
			"\tloop: Whether to loop the scenarios or not\n")
		os.Exit(0)
	}

	if *configFile == "" {
		log.Fatal("No config file provided. Set path with -config option.")

	}
	file, err := os.Open(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	if apiKey != nil && *apiKey != "" {
		config.ApiKey = *apiKey
	}

	if config.HostIp == "" || config.ApiKey == "" || config.LogPath == "" || config.ScenariosPath == "" {
		log.Fatal("All arguments must be set. Use --help to see the available options.")
	}
	return config
}

func createSimulations(allScenarios []scenario.Scenario, loop bool, client *http.Client) map[string]*simulation.Simulation {
	var simulations = make(map[string]*simulation.Simulation)
	for _, currScenario := range allScenarios {
		simulations[currScenario.CarId] = simulation.New(currScenario, loop, client)
	}
	return simulations
}

// Check if last status of the car is within the accepted delay
func isVehicleCommunicating(car openapi.Car) bool {
	const acceptedDelay = int64(2 * sleepTime * 1000)
	carTime := *car.LastState.Timestamp + acceptedDelay
	return carTime >= time.Now().UnixMilli()
}

func monitorAndStartNewCars(client *http.Client, simulations map[string]*simulation.Simulation) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(simulations))
	log.Print("[INFO] Starting monitoring of cars. Simulation will start, once a car starts communicating.")
	go func() { // Start a new goroutine
		activeCars := make(map[int32]bool)
		for {
			cars := client.GetCars()
			for _, car := range cars {
				if _, active := activeCars[*car.Id]; !active {
					if isVehicleCommunicating(car) {
						// If the car is not in the activeCars map, it's a new car.
						activeCars[*car.Id] = true
						// Start the simulation for the new car.
						if carSimulation, ok := simulations[car.Name]; ok {
							carSimulation.SetCarId(car.Id)
							go carSimulation.Start(&waitGroup) // Start the simulation in a new goroutine
						} else {
							log.Printf("[INFO] New car connected: %v, this car doesn't have any available scenario", car.Name)
						}
					}
				}
			}
			time.Sleep(sleepTime * time.Second)
		}
	}()

	waitGroup.Wait() // Wait for all scenarios to finish
}

func setUpLogger(path string) {
	file, err := os.OpenFile(path+"/virtual-fleet.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Lmicroseconds)

}

func setSignalHandler() {
	ic := make(chan os.Signal, 1)
	signal.Notify(ic, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ic
		fmt.Printf("[INFO] signal received, exiting\n")
		// TODO do we want to cancel orders?
		os.Exit(0)
	}()
}
