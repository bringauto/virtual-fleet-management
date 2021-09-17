package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"proto_helper"
	"syscall"
	ip "virtual_industrial_portal"
)

func main() {

	brokerIp := flag.String("broker-ip", "127.0.0.1", "IPv4 address of MQTT broker")
	brokerPort := flag.String("broker-port", "1883", "Port of MQTT broker")
	logPath := flag.String("log-path", "./", "Path for log file")
	scenariosPath := flag.String("scenarios", "./scenarios", "Path of scenarios folder")


	flag.Parse()
	setUpLogger(*logPath)
	var server = *brokerIp + ":" + *brokerPort
	proto_helper.CreateMessageBinaries()
	setSignalHandler()
	ip.Client.Start(server, "", "", *scenariosPath)
}

func setUpLogger(path string){
		file, err := os.OpenFile(path + "/virtual-industrial-portal.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
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

