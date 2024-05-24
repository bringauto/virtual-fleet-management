package scenario

import (
	"fmt"
	"log"
)

type ScenarioStruct struct {
	Map      string          `json:"map"`
	Missions []MissionStruct `json:"missions"`
	Routes   []RouteStruct   `json:"routes"`
}

type MissionStruct struct {
	DelaySeconds int32  `json:"delay_seconds"`
	Name         string `json:"name"`
	Stops        []struct {
		Name string `json:"name"`
	} `json:"stops"`
	Route string `json:"route"`
}

type Position struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude  float32 `json:"altitude"`
}

type StationStruct struct {
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

type RouteStruct struct {
	Name     string          `json:"name"`
	Stations []StationStruct `json:"stations"`
}

type Scenario struct {
	CarId    string
	Missions []MissionStruct
	Routes   []RouteStruct
}

func NewScenario(scenarioStruct ScenarioStruct, carId string) Scenario {
	scenario := Scenario{
		CarId:    carId,
		Missions: scenarioStruct.Missions,
		Routes:   scenarioStruct.Routes,
	}
	if !scenario.isValid() {
		panic(fmt.Sprintf("[ERROR] Scenario for car %v is not valid", carId))
	}
	return scenario
}

func (scenario *Scenario) isValid() bool {
	for _, mission := range scenario.Missions {
		if !scenario.areStopsOnRoute(mission) {
			log.Printf("[ERROR] Scenario %v: Stops in mission %v are not on the route %v\n.", scenario.CarId, mission.Name, mission.Route)
			return false
		}
	}
	return true
}

// / Check if all stops in the mission are present on the route
func (scenario *Scenario) areStopsOnRoute(mission MissionStruct) bool {
	stations := scenario.getStations(mission.Route)
	for _, stop := range mission.Stops {
		if !containsStation(stations, stop.Name) {
			return false
		}
	}
	return true
}

func containsStation(stations []StationStruct, stationName string) bool {
	for _, station := range stations {
		if station.Name == stationName {
			return true
		}
	}
	return false
}

func (scenario *Scenario) getStations(routeName string) (stationList []StationStruct) {
	for _, route := range scenario.Routes {
		if route.Name == routeName {
			return route.Stations
		}
	}
	log.Printf("[ERROR] Selected route does not exist: %v\n", routeName)
	return nil
}

func (scenario *Scenario) GetTotalDelay() (totalDelay int32) {
	for _, mission := range scenario.Missions {
		totalDelay += mission.DelaySeconds
	}
	return totalDelay
}
