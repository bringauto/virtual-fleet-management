package virtual_industrial_portal

import (
	"fmt"
	"log"
	"strconv"
	"time"
)


type ScenarioStruct struct {
	Map     string `json:"map"`
	Missions 	[]MissionStruct `json:"missions"`
}

type MissionStruct struct {
	Timestamp string `json:"timestamp"`
	Name string `json:"name"`
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
	startTimestamp int64
	missionTimer CancelableTimer
}


func NewScenario(scenarioStruct ScenarioStruct, topic string, loop bool)*Scenario{
	scenario := new(Scenario)
	scenario.scenarioData = scenarioStruct
	scenario.topic = topic
	scenario.loop = loop
	scenario.missionTimer = CancelableTimer{timer: nil, cancelTimer: make(chan struct{})}
	scenario.resetScenario()
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
		panic(fmt.Sprintf("[%v] Vehicle trying to mark wrong stop (%v) as done, received when no stop is on mission %v", scenario.topic, stopToMark, scenario.currentMission.Name))
	}

	if stopToMark == scenario.currentMission.Stops[0].Name {
			scenario.currentMission.Stops = scenario.currentMission.Stops[1:]
			scenario.missionChanged = true
	} else {
		panic(fmt.Sprintf("[%v] Vehicle trying to mark wrong stop as done, received: %s, should be: %s, mission: %v", scenario.topic, stopToMark, scenario.currentMission.Stops[0].Name, scenario.currentMission.Name))
	} 

	scenario.updateMissionIfEmpty()
}

func(scenario *Scenario)updateMissionIfEmpty(){
	if(len(scenario.currentMission.Stops) > 0){
		return
	}

	if(len(scenario.remainingMissions) > 0 ){
		log.Printf("[INFO] [%v] Car have finished mission %v, waiting for next scheduled mission", scenario.topic, scenario.currentMission.Name)
		return
	}

	if(scenario.loop){
		log.Printf("[INFO] [%v] Car have finished all of its missions, starting scenario again", scenario.topic)
		scenario.resetScenario()
		return
	}

	log.Printf("[INFO] [%v] All missions have been finished", scenario.topic)
}

func (scenario *Scenario)markMissionAccepted(){
	scenario.missionChanged = false;
}

func (scenario *Scenario)setNextMissionTimer(){
	if(len(scenario.remainingMissions) > 0){
		missionTimeOffset, err := strconv.ParseInt(scenario.remainingMissions[0].Timestamp, 10, 64)
		if err != nil {
			log.Printf("[INFO] [%v] Next mission (%v) timestamp has wrong format(%v), defaulting to 1 minute", scenario.topic,scenario.remainingMissions[0].Name ,scenario.remainingMissions[0].Timestamp)
			missionTimeOffset = 60
		}
		startNextMissionTimestamp := missionTimeOffset + scenario.startTimestamp
		calculatedTimerTime := startNextMissionTimestamp - time.Now().Unix()
		if(calculatedTimerTime < 1 ){
			log.Printf("[WARNING] [%v] Calculated time to next mission (%v) seems wrong, defaulting to one minute", scenario.topic, calculatedTimerTime)
			calculatedTimerTime = 60
		}
		log.Printf("[INFO] [%v] Next mission (%v) timestamp: %v, mission will start in %vs", scenario.topic,scenario.remainingMissions[0].Name, startNextMissionTimestamp, calculatedTimerTime)
		scenario.startMissionTimer(int(calculatedTimerTime))
	}
}

func (scenario *Scenario)startMissionTimer(duration int){
	if scenario.missionTimer.timer == nil {
		scenario.missionTimer.timer = time.NewTimer(time.Duration(duration) * time.Second)

	} else {
		scenario.missionTimer.timer.Reset(time.Duration(duration) * time.Second)
	}

	go func() {
		select {
		case <-scenario.missionTimer.timer.C:
			if(len(scenario.currentMission.Stops) > 0){
				log.Printf("[WARNING] [%v] Mission (%v) timeout! Starting new mission (%v)\n", scenario.topic, scenario.currentMission.Name, scenario.remainingMissions[0].Name)
			}else{
				log.Printf("[INFO] [%v] Starting next scheduled mission (%v)\n", scenario.topic, scenario.remainingMissions[0].Name)
			}
			scenario.popNextMission()
		case <-scenario.missionTimer.cancelTimer:
			scenario.missionTimer.timer = nil
		}
	}()
}

func(scenario *Scenario)resetScenario(){
	scenario.remainingMissions = scenario.scenarioData.Missions
	scenario.startTimestamp = time.Now().Unix()
	if(len(scenario.remainingMissions) > 0){
		scenario.popNextMission()
	}
}

func (scenario *Scenario)popNextMission(){
	if(len(scenario.remainingMissions) > 0){
		scenario.currentMission = scenario.remainingMissions[0]
		scenario.remainingMissions = scenario.remainingMissions[1:]
		scenario.setNextMissionTimer()
		scenario.missionChanged = true
	}
}