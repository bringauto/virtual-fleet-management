package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"virtual_fleet_management/pkg/http_client"
	"virtual_fleet_management/pkg/scenario"
	"virtual_fleet_management/pkg/simulation"
)

func main() {

	hostIp := flag.String("host", "http://172.18.0.1:8081", "IPv4 address")
	apiKey := flag.String("api-key", "123456", "API key")
	logPath := flag.String("log-path", "./", "Path for log file")
	scenariosPath := flag.String("scenario-dir", "./scenarios/virtual_vehicle", "Path of scenarios folder")
	loop := flag.Bool("loop", false, "Set true if scenarios should be run in loops")

	flag.Parse()
	setUpLogger(*logPath)
	cars := scenario.GetCarIdList(*scenariosPath)

	var allScenarios []scenario.Scenario
	for _, car := range cars {
		allScenarios = append(allScenarios, scenario.GetScenario(car, *scenariosPath))
	}

	var simulations map[string]simulation.Simulation
	for _, currScenario := range allScenarios {
		simulations[currScenario.CarId] = simulation.New(currScenario, *loop)
	}

	client := http_client.CreateClient(*hostIp, *apiKey)
	client.AddStop(allScenarios[0].Routes[0].Stations[1])
	//var url = *brokerIp + ":" + *brokerPort

	//setSignalHandler()
	//ip.Client.Start(server, "", "", *scenariosPath, *loop, 0)
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
		//ip.Client.Disconnect()
		fmt.Printf("[INFO] signal received, exiting\n")
		os.Exit(0)
	}()
}
