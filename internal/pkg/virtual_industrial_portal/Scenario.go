package virtual_industrial_portal


type ScenarioStruct struct {
	Map     string `json:"map"`
	Missions 	[]MissionStruct `json:"missions"`
}

type MissionStruct struct {
	Timestamp string `json:"timestamp"`
	Stops     []struct {
		Name string `json:"name"`
	} `json:"stops"`
}

type Scenario struct{
	scenarioStruct ScenarioStruct
	currentMission MissionStruct
	missionChanged bool
}


func NewScenario(scenarioStruct ScenarioStruct)*Scenario{
	scenario := new(Scenario)
	scenario.scenarioStruct = scenarioStruct
	if(len(scenarioStruct.Missions) > 0 ){
		scenario.currentMission = scenarioStruct.Missions[0]
	}
	scenario.missionChanged = true
	return scenario
}

func (scenario Scenario)getStopList()[]string{
	var stopList []string
	for _, stop := range scenario.currentMission.Stops{
		stopList = append(stopList, stop.Name)
	}
	return stopList
}