package simulation

import (
	"fmt"
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"log"
	"math"
	"reflect"
	"strconv"
	"time"
	"virtual_fleet_management/pkg/http_client"
	"virtual_fleet_management/pkg/scenario"
)

type CancelableTimer struct {
	timer       *time.Timer
	cancelTimer chan struct{}
	durationSec int
}

type Simulation struct {
	loop bool
	//scenario          scenario.Scenario //TODO needed?
	currentMission    *scenario.MissionStruct
	remainingMissions []scenario.MissionStruct
	missionChanged    bool
	startTimestamp    int64
	missionTimer      CancelableTimer
	client            http_client.Client
}

func New(scenario2 scenario.Scenario, loop bool) Simulation {
	simulation := Simulation{
		//scenario:          scenario2,
		missionTimer:      CancelableTimer{timer: nil, cancelTimer: make(chan struct{})},
		loop:              loop,
		remainingMissions: scenario2.Missions,
	}
	simulation.initDatabase(scenario2)
	return simulation
}

const gpsEqualityThreshold = 1e-6

func gpsEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) <= gpsEqualityThreshold
}

func isPositionEqual(position1 scenario.Position, position2 openapi.GNSSPosition) bool {
	return gpsEqual(position1.Latitude, *position2.Latitude) && gpsEqual(position1.Longitude, *position2.Longitude)
}
func findStationId(station scenario.StationStruct, existingStations []openapi.Stop) *int32 {
	for _, existingStation := range existingStations {
		if station.Name == existingStation.Name {
			if isPositionEqual(station.Position, existingStation.Position) {
				return existingStation.Id
			} else {
				log.Printf("[ERROR] Station %v already exists, but with different position", station.Name)
				// TODO update station? Exit program?
				return nil
			}
		}
	}
	return nil
}

func findRouteId(route *openapi.Route, existingRoutes []openapi.Route) *int32 {
	for _, existingRoute := range existingRoutes {
		if route.Name == existingRoute.Name {
			if reflect.DeepEqual(route.StopIds, existingRoute.StopIds) {
				return existingRoute.Id
			} else {
				log.Printf("[ERROR] Route %v already exists, but with different stops", route.Name)
				// TODO update route? Exit program?
				return nil
			}
		}
	}
	return nil

}
func (simulation Simulation) initDatabase(scenario2 scenario.Scenario) {
	existingStations := simulation.client.GetStops()
	for _, route := range scenario2.Routes {
		var stopIds []int32
		for _, station := range route.Stations {
			stationId := findStationId(station, existingStations)
			if stationId == nil {
				stationId = simulation.client.AddStop(station)
			}
			stopIds = append(stopIds, *stationId)

		}
		newRoute := openapi.NewRoute(route.Name)
		newRoute.SetStopIds(stopIds)
		routeId := findRouteId(newRoute, simulation.client.GetRoutes())
		if routeId == nil {
			routeId := simulation.client.AddRoute(newRoute)
			// TODO save routeIds and stopIds into dictionary
		}
	}
}

// TODO remove this
func (simulation Simulation) markStopAsDone(stopToMark string) {
	if len(simulation.currentMission.Stops) < 1 {
		panic(fmt.Sprintf("[%v] Vehicle trying to mark wrong stop (%v) as done, received when no stop is on mission %v", simulation.topic, stopToMark, simulation.currentMission.Name))
	}

	if stopToMark == simulation.currentMission.Stops[0].Name {
		simulation.currentMission.Stops = simulation.currentMission.Stops[1:]
		simulation.missionChanged = true
	} else {
		return
	}

	simulation.updateMissionIfEmpty()
}

func (simulation Simulation) updateMissionIfEmpty() {
	if len(simulation.currentMission.Stops) > 0 {
		return
	}

	if len(simulation.remainingMissions) > 0 {
		log.Printf("[INFO] [%v] Car have finished mission %v, waiting for next scheduled mission", simulation.topic, simulation.currentMission.Name)
		return
	}

	if simulation.loop {
		log.Printf("[INFO] [%v] Car have finished all of its missions, starting simulation again", simulation.topic)
		simulation.resetsimulation()
		return
	}

	log.Printf("[INFO] [%v] All missions have been finished", simulation.topic)
}

func (simulation Simulation) markMissionAccepted() {
	simulation.missionChanged = false
}

func (simulation Simulation) setNextMissionTimer() {
	if len(simulation.remainingMissions) > 0 {
		missionTimeOffset, err := strconv.ParseInt(simulation.remainingMissions[0].Timestamp, 10, 64)
		if err != nil {
			log.Printf("[INFO] [%v] Next mission (%v) timestamp has wrong format(%v), defaulting to 1 minute", simulation.topic, simulation.remainingMissions[0].Name, simulation.remainingMissions[0].Timestamp)
			missionTimeOffset = 60
		}
		startNextMissionTimestamp := missionTimeOffset + simulation.startTimestamp
		calculatedTimerTime := startNextMissionTimestamp - time.Now().Unix()
		if calculatedTimerTime < 1 {
			log.Printf("[WARNING] [%v] Calculated time to next mission (%v) seems wrong, defaulting to one minute", simulation.topic, calculatedTimerTime)
			calculatedTimerTime = 60
		}
		log.Printf("[INFO] [%v] Next mission (%v) timestamp: %v, mission will start in %vs", simulation.topic, simulation.remainingMissions[0].Name, startNextMissionTimestamp, calculatedTimerTime)
		simulation.startMissionTimer(int(calculatedTimerTime))
	}
}

func (simulation Simulation) startMissionTimer(duration int) {
	if simulation.missionTimer.Timer == nil {
		simulation.missionTimer.Timer = time.NewTimer(time.Duration(duration) * time.Second)

	} else {
		simulation.missionTimer.Timer.Reset(time.Duration(duration) * time.Second)
	}

	go func() {
		select {
		case <-simulation.missionTimer.Timer.C:
			if len(simulation.currentMission.Stops) > 0 {
				log.Printf("[WARNING] [%v] Mission (%v) timeout! Starting new mission (%v)\n", simulation.topic, simulation.currentMission.Name, simulation.remainingMissions[0].Name)
			} else {
				log.Printf("[INFO] [%v] Starting next scheduled mission (%v)\n", simulation.topic, simulation.remainingMissions[0].Name)
			}
			simulation.popNextMission()
		case <-simulation.missionTimer.CancelTimer:
			simulation.missionTimer.Timer = nil
		}
	}()
}

func (simulation Simulation) resetsimulation() {
	simulation.remainingMissions = simulation.simulationData.Missions
	simulation.startTimestamp = time.Now().Unix()
	if len(simulation.remainingMissions) > 0 {
		simulation.popNextMission()
	}
}

func (simulation Simulation) popNextMission() {
	if len(simulation.remainingMissions) > 0 {
		simulation.currentMission = simulation.remainingMissions[0]
		simulation.remainingMissions = simulation.remainingMissions[1:]
		simulation.setNextMissionTimer()
		simulation.missionChanged = true
	}
}
