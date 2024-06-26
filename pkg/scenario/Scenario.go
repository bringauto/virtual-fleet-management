package scenario

import (
	"fmt"
	"log"
)

type ScenarioStruct struct {
	Map             string          `json:"map"`
	StartingStation string          `json:"starting_station"`
	Missions        []MissionStruct `json:"missions"`
	Routes          []RouteStruct   `json:"routes"`
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
	CarId           string
	StartingStation string
	StartingRoute   string
	Missions        []MissionStruct
	Routes          []RouteStruct
}

func NewScenario(scenarioStruct ScenarioStruct, carId string) Scenario {
	scenario := Scenario{
		CarId:           carId,
		StartingStation: scenarioStruct.StartingStation,
		Missions:        scenarioStruct.Missions,
		StartingRoute:   getStartingRoute(scenarioStruct.Routes, scenarioStruct.StartingStation),
		Routes:          scenarioStruct.Routes,
	}
	if !scenario.IsValid() {
		panic(fmt.Sprintf("[ERROR] Scenario for car %v is not valid", carId))
	}
	return scenario
}

func (scenario *Scenario) IsValid() bool {
	for _, route := range scenario.Routes {
		if route.Stations == nil {
			log.Printf("[ERROR] Scenario %v: Route %v has no stations\n.", scenario.CarId, route.Name)
			return false
		}
		if !scenario.areStationsUnique(route.Name) {
			return false
		}
	}
	if scenario.StartingRoute == "" {
		log.Printf("[ERROR] Scenario %v: starting_station is not on any route\n.", scenario.CarId)
		return false
	}
	for _, mission := range scenario.Missions {
		if mission.DelaySeconds < 0 {
			log.Printf("[ERROR] Scenario %v: Mission %v has negative delay\n.", scenario.CarId, mission.Name)
			return false
		}
		if mission.Route == "" {
			log.Printf("[ERROR] Scenario %v: Mission %v has no route defined\n.", scenario.CarId, mission.Name)
			return false
		}
		if mission.Stops == nil {
			log.Printf("[ERROR] Scenario %v: Mission %v has no stops\n.", scenario.CarId, mission.Name)
			return false
		}
		if !scenario.areStopsOnRoute(mission) {
			log.Printf("[ERROR] Scenario %v: Stops in mission %v are not on the route %v\n.", scenario.CarId, mission.Name, mission.Route)
			return false
		}
	}
	return true
}

func getStartingRoute(routes []RouteStruct, startingStation string) string {
	for _, route := range routes {
		if containsStation(route.Stations, startingStation) {
			return route.Name
		}
	}
	return ""
}

func (scenario *Scenario) areStationsUnique(routeName string) bool {
	stations := scenario.getStations(routeName)
	stationMap := make(map[string]bool)

	for _, station := range stations {
		if _, exists := stationMap[station.Name]; exists {
			log.Printf("[ERROR] Duplicate station found on route %v: %v\n", routeName, station.Name)
			return false
		}
		stationMap[station.Name] = true
	}
	return true
}

// / Check if all stops in the mission are present on the route
func (scenario *Scenario) areStopsOnRoute(mission MissionStruct) bool {
	stations := scenario.getStations(mission.Route)
	for _, stop := range mission.Stops {
		if !containsStation(stations, stop.Name) {
			log.Printf("[ERROR] Stop %v in mission %v is not on the route %v\n", stop.Name, mission.Name, mission.Route)
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
