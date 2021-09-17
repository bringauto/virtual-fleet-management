package virtual_industrial_portal

import (
	"log"
)

type ScenarioStruct struct {
	Map      string `json:"map"`
	Missions []struct {
		Timestamp string `json:"timestamp"`
		Stops     []struct {
			Name string `json:"name"`
		} `json:"stops"`
	} `json:"missions"`
}

type Scenario struct{
	scenarioStructs []ScenarioStruct
}


func NewScenario(scenarioStructs []ScenarioStruct)*Scenario{
	scenario := new(Scenario)
	scenario.scenarioStructs = scenarioStructs
	log.Printf("[INFO] creating new scenario %v\n", scenarioStructs)
	return scenario
}