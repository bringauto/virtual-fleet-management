package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"virtual_fleet_management/pkg/http_client"
	"virtual_fleet_management/pkg/scenario"
	"virtual_fleet_management/pkg/simulation"
)

func main() {

	hostIp := flag.String("host", "http://127.0.0.1:8081", "IPv4 address")
	apiKey := flag.String("api-key", "123456", "API key")
	logPath := flag.String("log-path", "./", "Path for log file")
	scenariosPath := flag.String("scenario-dir", "./scenarios/virtual_vehicle", "Path of scenarios folder")
	loop := flag.Bool("loop", false, "Set true if scenarios should be run in loops")

	flag.Parse()
	setUpLogger(*logPath)
	setSignalHandler()

	cars := scenario.GetCarIdList(*scenariosPath)

	var allScenarios []scenario.Scenario
	for _, car := range cars {
		allScenarios = append(allScenarios, scenario.GetScenario(car, *scenariosPath))
	}

	client := http_client.CreateClient(*hostIp, *apiKey)

	var simulations = make(map[string]*simulation.Simulation)
	for _, currScenario := range allScenarios {
		simulations[currScenario.CarId] = simulation.New(currScenario, *loop, client)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	startNewCars(client, simulations)
	wg.Wait() // Wait for all goroutines to finish
}

func startNewCars(client *http_client.Client, simulations map[string]*simulation.Simulation) {
	existingCars := make(map[int32]bool)

	go func() { // Start a new goroutine
		for {
			cars := client.GetCars()
			for _, car := range cars {
				if _, exists := existingCars[*car.Id]; !exists {
					// If the car is not in the existingCars map, it's a new car.
					existingCars[*car.Id] = true

					// Start the simulation for the new car.
					if carSimulation, ok := simulations[car.Name]; ok {
						carSimulation.SetCarId(car.Id)
						go carSimulation.Start() // Start the simulation in a new goroutine
					} else {
						log.Printf("[INFO] New car connected: %v, this car doesn't have any available scenario", car.Name)
					}
				}
			}

			// Sleep for a while before the next iteration to avoid busy looping.
			time.Sleep(10 * time.Second) // TODO timeout
		}
	}()
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
		os.Exit(0)
	}()
}
