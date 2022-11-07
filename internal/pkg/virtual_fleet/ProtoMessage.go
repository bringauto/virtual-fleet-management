package virtual_industrial_portal

import (
	pb "ba_proto"
	"errors"
	"fmt"
	"io/ioutil"

	//proto "google.golang.org/protobuf"
	proto "github.com/golang/protobuf/proto"
)

func GetIndustrialPortalConnectResponse(responseType pb.ConnectResponse_Type, sessionId string) []byte {
	var connectReponse = pb.ConnectResponse{Type: responseType, SessionId: sessionId}
	var ipConnectResponse = pb.MessageIndustrialPortal_ConnectReponse{ConnectReponse: &connectReponse}
	var message = pb.MessageIndustrialPortal{Type: &ipConnectResponse}

	return getBinaryFromMessage(&message)
}

func GetIndustrialPortalStatusResponse(sessionId string) []byte {
	var statusResponse = pb.StatusResponse{Type: pb.StatusResponse_OK, SessionId: sessionId}
	var ipStatusResponse = pb.MessageIndustrialPortal_StatusResponse{StatusResponse: &statusResponse}
	var message = pb.MessageIndustrialPortal{Type: &ipStatusResponse}
	return getBinaryFromMessage(&message)
}

func GetIndustrialPortalCommand(action pb.CarCommand_Action, stops []string, route string, sessionId string, stations []StationStruct) []byte {
	stopList := []*pb.Stop{}
	for _, element := range stops {
		stopList = append(stopList, &pb.Stop{To: element})
	}

	stationList := []*pb.Station{}
	if stations != nil {
		for _, station := range stations {
			stationList = append(stationList, &pb.Station{Name: station.Name, Position: &pb.Station_Position{Latitude: station.Position.Latitude, Longitude: station.Position.Longitude, Altitude: station.Position.Altitude}})
		}
	}

	var carCommand = pb.CarCommand{Stops: stopList, Action: action, Route: route, RouteStations: stationList}

	var command = pb.Command{CarCommand: &carCommand, SessionId: sessionId}
	var ipCommand = pb.MessageIndustrialPortal_Command{Command: &command}
	var message = pb.MessageIndustrialPortal{Type: &ipCommand}
	return getBinaryFromMessage(&message)
}

func GetDaemonConnect(company, name, sessionId string) []byte {
	var connect = pb.Connect{Company: company, Name: name, SessionId: sessionId}
	var daeConnect = pb.MessageDaemon_Connect{Connect: &connect}
	var message = pb.MessageDaemon{Type: &daeConnect}
	return getBinaryFromMessage(&message)
}

func GetDaemonCommandResponse(sessionId string) []byte {
	var commandResponse = pb.CommandResponse{Type: pb.CommandResponse_OK, SessionId: sessionId}
	var daeCommandRespond = pb.MessageDaemon_CommandResponse{CommandResponse: &commandResponse}
	var message = pb.MessageDaemon{Type: &daeCommandRespond}
	return getBinaryFromMessage(&message)
}

func GetDaemonStatus(latitude, longtitude, altitude, speed, fuel float64, sessionId, stopTo string, serverErrorType pb.Status_ServerError_Type, finishedStops []string) []byte {

	var position = pb.CarStatus_Position{Latitude: latitude, Longitude: altitude, Altitude: altitude}
	var telemetry = pb.CarStatus_Telemetry{Speed: speed, Fuel: fuel, Position: &position}
	var stop = pb.Stop{To: stopTo}
	var carStatus = pb.CarStatus{Telemetry: &telemetry, State: pb.CarStatus_DRIVE, Stop: &stop}

	stopList := []*pb.Stop{}
	for _, element := range finishedStops {
		stopList = append(stopList, &pb.Stop{To: element})
	}
	if serverErrorType == pb.Status_ServerError_OK && len(finishedStops) > 0 {
		panic(fmt.Sprint("Status message contains stops, but server error is OK"))
	}

	var serverError = pb.Status_ServerError{Type: serverErrorType, Stops: stopList}

	var status = pb.Status{CarStatus: &carStatus, Server: &serverError, SessionId: sessionId}
	var daeStatus = pb.MessageDaemon_Status{Status: &status}
	var message = pb.MessageDaemon{Type: &daeStatus}
	return getBinaryFromMessage(&message)
}

func getBinaryFromMessage(message proto.Message) []byte {
	out, err := proto.Marshal(message)

	if err != nil {
		panic(fmt.Sprintf("Error while creating status message %v \n", err))
	}
	return out
}

func GetDaemonMessageFromBinary(binaryMessage []byte) (mes *pb.MessageDaemon, err error) {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if recover() != nil {
			err = errors.New(fmt.Sprintf("[ERROR] Unable to parse daemon message from binary %v\n", binaryMessage))
		}
	}()

	message := &pb.MessageDaemon{}

	err_mar := proto.Unmarshal(binaryMessage, message)

	if err_mar != nil {
		panic(fmt.Sprintf("Unable to convert binary to daemon message %v", err))
	}
	return message, nil
}

func GetIndustrialMessageFromBinary(binaryMessage []byte) *pb.MessageIndustrialPortal {
	message := &pb.MessageIndustrialPortal{}

	err := proto.Unmarshal(binaryMessage, message)

	if err != nil {
		panic(fmt.Sprintf("Unable to convert binary to daemon message %v", err))
	}
	return message
}

func PrintDataToFile(filepath, filename string, message []byte) {
	err := ioutil.WriteFile(filepath+"/"+filename, message, 0644)
	if err != nil {
		panic(fmt.Sprintf("Cannot write proto buffer into file %s\n", "./"+filename))
	}
}
