package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	ip "virtual_fleet"
)

func main() {

	brokerIp := flag.String("broker-ip", "127.0.0.1", "IPv4 address of MQTT broker")
	brokerPort := flag.String("broker-port", "1883", "Port of MQTT broker")
	logPath := flag.String("log-path", "./", "Path for log file")
	scenariosPath := flag.String("scenario-dir", "./scenarios", "Path of scenarios folder")
	loop := flag.Bool("loop", false, "Set true if scenarios should be run in loops")


	flag.Parse()
	setUpLogger(*logPath)
	var server = *brokerIp + ":" + *brokerPort
	setSignalHandler()
	ip.Client.Start(server, "", "", *scenariosPath, *loop)
}

func setUpLogger(path string){
		file, err := os.OpenFile(path + "/virtual-fleet.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
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
		ip.Client.Disconnect()
		fmt.Printf("[INFO] signal received, exiting\n");
		os.Exit(0)
	}()
}

