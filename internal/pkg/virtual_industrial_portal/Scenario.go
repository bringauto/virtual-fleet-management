package virtual_industrial_portal

import (
	"fmt"
	"log"
)


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
	topic string
	scenarioData ScenarioStruct
	remainingMissions []MissionStruct
	currentMission MissionStruct
	missionChanged bool
	loop bool
}


func NewScenario(scenarioStruct ScenarioStruct, topic string, loop bool)*Scenario{
	scenario := new(Scenario)
	scenario.scenarioData = scenarioStruct
	scenario.remainingMissions = scenarioStruct.Missions
	scenario.missionChanged = true
	scenario.topic = topic
	scenario.loop = loop
	scenario.updateMissionIfEmpty()
	return scenario
}

func (scenario *Scenario)getStopList()[]string{
	var stopList []string
	for _, stop := range scenario.currentMission.Stops{
		stopList = append(stopList, stop.Name)
	}
	return stopList
}

func (scenario *Scenario)markStopAsDone(stopToMark string){
	if(len(scenario.currentMission.Stops) < 1 ){
		panic(fmt.Sprintf("Vehicle %s trying to mark wrong stop as done, received when no stop is on mission: %s, should be: %s", scenario.topic, stopToMark, scenario.currentMission.Stops[0].Name))
	}

	if stopToMark == scenario.currentMission.Stops[0].Name {
			scenario.currentMission.Stops = scenario.currentMission.Stops[1:]
			scenario.missionChanged = true
	} else {
		panic(fmt.Sprintf("Vehicle %s trying to mark wrong stop as done, received: %s, should be: %s", scenario.topic, stopToMark, scenario.currentMission.Stops[0].Name))
	} 

	scenario.updateMissionIfEmpty()
}

func(scenario *Scenario)updateMissionIfEmpty(){
	if(len(scenario.currentMission.Stops) > 0){
		return
	}
	if(len(scenario.remainingMissions) > 0 ){
		scenario.currentMission = scenario.remainingMissions[0]
		scenario.remainingMissions = scenario.remainingMissions[1:]
		log.Printf("[INFO] Car %v have finished mission, starting new mission %v", scenario.topic, scenario.currentMission)
		return
	}

	if(scenario.loop){
		log.Printf("[INFO] Car %v have finished all of its missions, starting again", scenario.topic)
		scenario.remainingMissions = scenario.scenarioData.Missions
		if(len(scenario.remainingMissions) > 0){
			scenario.updateMissionIfEmpty()
		}
	}
	log.Printf("[INFO] Car %v have finnished all of its missions", scenario.topic)
}

func (scenario *Scenario)markMissionAccepted(){
	scenario.missionChanged = false;
}