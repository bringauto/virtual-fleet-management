package virtual_industrial_portal


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
	return scenario
}

func (scenario Scenario)getCurentMission()[]string{
	if(len(scenario.scenarioStructs) > 0){
		return scenario.scenarioStructs[0].Missions[0].Stops
	}
}