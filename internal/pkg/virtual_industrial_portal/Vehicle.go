package virtual_industrial_portal

import (
	pb "ba_proto"
	"fmt"
	"log"
	"proto_helper"
	"time"
)

type Vehicle struct {
	daemonTopic, industrialPortalTopic, sessionId string
	connectionState                               int
	vehicleState                                  pb.CarStatus_State
	//stops, actualMission                          []string
	timeoutTimer, responseTimer                   vehicleTimer
	//missionChanged								  bool
	scenario									  *Scenario
}

type vehicleTimer struct {
	timer       *time.Timer
	cancelTimer chan struct{}
	durationSec int
}

func NewVehicle(topic string, scenario *Scenario) *Vehicle {
	vehicle := new(Vehicle)
	vehicle.daemonTopic = topic + "/daemon"
	vehicle.industrialPortalTopic = topic + "/industrial_portal"
	vehicle.sessionId = ""
	vehicle.connectionState = CONNECTION_DISCONNECTED
	vehicle.vehicleState = pb.CarStatus_ERROR
	//vehicle.stops = stopList
	//vehicle.actualMission = stopList
	vehicle.timeoutTimer = vehicleTimer{timer: nil, cancelTimer: make(chan struct{}), durationSec: 30}
	vehicle.responseTimer = vehicleTimer{timer: nil, cancelTimer: make(chan struct{}), durationSec: 10}
	//vehicle.missionChanged = true
	vehicle.scenario = scenario

/* 	for _, scenario := range vehicle.scenario.scenarioStructs{
		log.Printf("[INFO] Vehicle creation: %v %v\n",topic, scenario)
	} */
	return vehicle
}

const (
	CONNECTION_DISCONNECTED = iota
	CONNECTION_CONNECTING
	CONNECTION_CONNECTED
)

func (vehicle *Vehicle) parseMessage(binaryMessage []byte) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[ERROR] ", r)
		}
	}()

	daemonMessage, err := proto_helper.GetDaemonMessageFromBinary(binaryMessage)

	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	if daemonMessage.ProtoReflect().GetUnknown() != nil {
		panic(fmt.Sprintf("Unknown message, unrecognized bytes: %v", daemonMessage.ProtoReflect().GetUnknown()))
	}

	var message = daemonMessage.GetType()
	switch message.(type) {
	case *pb.MessageDaemon_Connect:
		vehicle.parseConnect(daemonMessage.GetConnect())
	case *pb.MessageDaemon_Status:
		vehicle.parseStatus(daemonMessage.GetStatus())
	case *pb.MessageDaemon_CommandResponse:
		vehicle.parseCommandResponse(daemonMessage.GetCommandResponse())
	default:
		panic(fmt.Sprintf("Unknown message"))
	}
}

func (vehicle *Vehicle) parseConnect(message *pb.Connect) {
	if message.Company == "" || message.Name == "" || message.SessionId == "" {
		panic(fmt.Sprintf("Message connect has wrong format, ignoring: %v", message))
	}

	if vehicle.connectionState != CONNECTION_DISCONNECTED {
		vehicle.sendConnectResponse(message.SessionId, pb.ConnectResponse_ALREADY_LOGGED)
		panic(fmt.Sprintf("Vehicle %v is trying to connect to working session (received sessionID: %v, active sessionID: %v)", vehicle.daemonTopic, message.SessionId, vehicle.sessionId))
	}

	vehicle.resetTimeoutTimer()
	vehicle.sessionId = message.SessionId
	vehicle.changeState(CONNECTION_CONNECTING)
	vehicle.sendConnectResponse(vehicle.sessionId, pb.ConnectResponse_OK)
	log.Printf("[INFO] vehicle %v connected with session id %v\n", vehicle.daemonTopic, message.SessionId)
}

func (vehicle *Vehicle) parseStatus(message *pb.Status) {
	log.Printf("[INFO] Received status from %v, sessionId %v (State: %v, Stop: %v, Lon: %v, Lat: %vm, Alt: %v, Speed: %vm/s, Fuel %v%%)\n",
		vehicle.daemonTopic,
		message.SessionId,
		carStateEnumToString(message.CarStatus.State),
		message.CarStatus.Stop.To,
		message.CarStatus.Telemetry.Position.Longitude,
		message.CarStatus.Telemetry.Position.Latitude,
		message.CarStatus.Telemetry.Position.Altitude,
		message.CarStatus.Telemetry.Speed,
		message.CarStatus.Telemetry.Fuel)

	vehicle.connectionValidityCheck(message.SessionId, "status")
	vehicle.resetTimeoutTimer()
	vehicle.sendStatusResponse()

	switch message.Server.Type {
	case pb.Status_ServerError_SERVER_ERROR:
		var doneStops []string
		for _, stop := range message.Server.GetStops() {
			doneStops = append(doneStops, stop.To)

		}
		log.Printf("[WARNING] received server error with stops: %v, mission is: %v in connection in %v, sessionID: %v\n", doneStops, vehicle.scenario.getStopList(), vehicle.daemonTopic, vehicle.sessionId)
		for _, stop := range doneStops {
			vehicle.scenario.markStopAsDone(stop)
		}
	case pb.Status_ServerError_OK:
		if message.CarStatus.State == pb.CarStatus_IN_STOP && message.CarStatus.State != vehicle.vehicleState { //check for stop
			vehicle.scenario.markStopAsDone(message.CarStatus.Stop.To)
		}
	}

	if vehicle.connectionState == CONNECTION_CONNECTING { //first status message
		vehicle.changeState(CONNECTION_CONNECTED)
	}
	vehicle.vehicleState = message.CarStatus.State

	if(vehicle.scenario.missionChanged){
		vehicle.sendCommand()
	}
}


func (vehicle *Vehicle) parseCommandResponse(message *pb.CommandResponse) {
	vehicle.connectionValidityCheck(message.SessionId, "command response")
	vehicle.resetTimeoutTimer()

	if vehicle.responseTimer.timer != nil {
		vehicle.responseTimer.cancelTimer <- struct{}{}
	} else {
		panic(fmt.Sprintf("Received unexpected command response from (%v)", vehicle.daemonTopic))
	}
}

func (vehicle *Vehicle) connectionValidityCheck(receivedSessionId, messageType string) {
	if vehicle.connectionState == CONNECTION_DISCONNECTED {
		panic(fmt.Sprintf("Received message (%v) from disconnected car (%v), received sessionID: %v", messageType, vehicle.daemonTopic, receivedSessionId))
	}
	if vehicle.sessionId != receivedSessionId {
		panic(fmt.Sprintf("Received message (%v) from %s with wrong session id (should be: %v is: %v)", messageType, vehicle.daemonTopic, vehicle.sessionId, receivedSessionId))
	}
}

func (vehicle *Vehicle) changeState(state int) {
	vehicle.connectionState = state
	log.Printf("[INFO] Vehicle %v state changed to %v\n", vehicle.daemonTopic, connectionEnumToString(state))
}

func (vehicle *Vehicle) resetTimeoutTimer() {
	if vehicle.timeoutTimer.timer == nil {
		vehicle.timeoutTimer.timer = time.NewTimer(time.Duration(vehicle.timeoutTimer.durationSec) * time.Second)

	} else {
		vehicle.timeoutTimer.timer.Reset(time.Duration(vehicle.timeoutTimer.durationSec) * time.Second)
	}
	go func() {
		select {
		case <-vehicle.timeoutTimer.timer.C:
			log.Printf("[WARNING] Vehicle timeout! reseting vehicle %v\n", vehicle.industrialPortalTopic)
			if vehicle.connectionState != CONNECTION_DISCONNECTED {
				if vehicle.responseTimer.timer != nil {
					vehicle.responseTimer.cancelTimer <- struct{}{}
				}
				vehicle.resetVehicle()
			}
		case <-vehicle.timeoutTimer.cancelTimer:
			vehicle.timeoutTimer.timer = nil
		}

	}()
}

func (vehicle *Vehicle) startResponseTimer() {
	if vehicle.responseTimer.timer == nil {
		vehicle.responseTimer.timer = time.NewTimer(time.Duration(vehicle.responseTimer.durationSec) * time.Second)

	} else {
		vehicle.responseTimer.timer.Reset(time.Duration(vehicle.responseTimer.durationSec) * time.Second)
	}
	go func() {
		select {
		case <-vehicle.responseTimer.timer.C:
			log.Printf("[WARNING] Vehicle %s failed to send command response\n", vehicle.daemonTopic)
			if vehicle.connectionState != CONNECTION_DISCONNECTED {
				if vehicle.timeoutTimer.timer != nil {
					vehicle.timeoutTimer.cancelTimer <- struct{}{}
				}
				vehicle.resetVehicle()
			}
		case <-vehicle.responseTimer.cancelTimer:
			vehicle.responseTimer.timer = nil
		}
	}()

}

func (vehicle *Vehicle) resetVehicle() {
	vehicle.changeState(CONNECTION_DISCONNECTED)
	vehicle.sessionId = ""
	vehicle.vehicleState = pb.CarStatus_ERROR
	vehicle.timeoutTimer.timer = nil
	vehicle.responseTimer.timer = nil
}

func (vehicle *Vehicle) sendConnectResponse(sessionId string, responseType pb.ConnectResponse_Type) {
	var connectResponse = proto_helper.GetIndustrialPortalConnectResponse(responseType, sessionId)
	Client.publish(vehicle.industrialPortalTopic, connectResponse)
}

func (vehicle *Vehicle) sendCommand() {
	log.Printf("[INFO] Sending command to %v, sessionID: %v, command: START, mission:%v\n", vehicle.industrialPortalTopic, vehicle.sessionId, vehicle.scenario.getStopList())
	vehicle.startResponseTimer()
	var command = proto_helper.GetIndustrialPortalCommand(pb.CarCommand_START, vehicle.scenario.getStopList(), vehicle.sessionId)
	Client.publish(vehicle.industrialPortalTopic, command)
	vehicle.scenario.markMissionAccepted()
}

func (vehicle *Vehicle) sendStatusResponse() {
	var statusResponse = proto_helper.GetIndustrialPortalStatusResponse(vehicle.sessionId)
	Client.publish(vehicle.industrialPortalTopic, statusResponse)
}

func connectionEnumToString(state int) string {
	switch state {
	case CONNECTION_DISCONNECTED:
		return "Disconnected"
	case CONNECTION_CONNECTING:
		return "Connecting"
	case CONNECTION_CONNECTED:
		return "Connected"
	}
	return "Unknown state"
}

func carStateEnumToString(state pb.CarStatus_State) string {
	switch state {
	case pb.CarStatus_IDLE:
		return "IDLE"
	case pb.CarStatus_DRIVE:
		return "DRIVE"
	case pb.CarStatus_IN_STOP:
		return "IN_STOP"
	case pb.CarStatus_OBSTACLE:
		return "OBSTACLE"
	case pb.CarStatus_ERROR:
		return "ERROR"
	}
	return "Unknown state"
}
