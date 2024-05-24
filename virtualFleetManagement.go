package main

import (
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

func main() {
	hostIp, apiKey, logPath, scenariosPath, loop := parseFlags()

	setUpLogger(logPath)
	setSignalHandler()

	allScenarios := scenario.GetAllScenariosFromDir(scenariosPath)
	client := http.CreateClient(hostIp, apiKey)

	simulations := createSimulations(allScenarios, loop, client)

	monitorAndStartNewCars(client, simulations)
}

func parseFlags() (string, string, string, string, bool) {
	hostIp := flag.String("host", "http://127.0.0.1:8081", "IPv4 address")
	apiKey := flag.String("api-key", "123456", "API key")
	logPath := flag.String("log-path", "./", "Path for log file")
	scenariosPath := flag.String("scenario-dir", "./scenarios/virtual_vehicle", "Path of scenarios folder")
	loop := flag.Bool("loop", false, "Set true if scenarios should be run in loops")

	flag.Parse()

	return *hostIp, *apiKey, *logPath, *scenariosPath, *loop
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
