package scenario

import "log"

type ScenarioStruct struct {
	Map      string          `json:"map"`
	Missions []MissionStruct `json:"missions"`
	Routes   []RouteStruct   `json:"routes"`
}

type MissionStruct struct {
	Timestamp string `json:"timestamp"`
	Name      string `json:"name"`
	Stops     []struct {
		Name string `json:"name"`
	} `json:"stops"`
	Route string `json:"route"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
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
	carId    string
	missions []MissionStruct
	routes   []RouteStruct
}

func NewScenario(scenarioStruct ScenarioStruct, carId string) *Scenario {
	scenario := new(Scenario)
	scenario.carId = carId
	scenario.missions = scenarioStruct.Missions
	scenario.routes = scenarioStruct.Routes
	if !scenario.isValid() {
		return nil
	}
	return scenario
}

func (scenario *Scenario) isValid() bool {
	for _, mission := range scenario.missions {
		if !scenario.areStopsOnRoute(mission) {
			log.Printf("[ERROR] Scenario %v: Stops in mission %v are not on the route %v\n.", scenario.carId, mission.Name, mission.Route)
			// TODO panic?
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
	for _, route := range scenario.routes {
		if route.Name == routeName {
			return route.Stations
		}
	}
	log.Printf("[ERROR] Selected route does not exist: %v\n", routeName)
	return nil
}

func (scenario *Scenario) getAllStations() (stationList []StationStruct) {
	for _, route := range scenario.routes {
		for _, station := range route.Stations {
			stationList = append(stationList, station)
		}
	}
	return stationList
}

func (scenario *Scenario) getAllMissionStops() (stopList []string) {
	for _, mission := range scenario.missions {
		for _, stop := range mission.Stops {
			stopList = append(stopList, stop.Name)
		}
	}
	return stopList
}
