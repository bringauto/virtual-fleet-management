package simulation

import (
	"log"
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
	loop               bool
	simulationScenario scenario.Scenario
	currentMission     scenario.MissionStruct
	remainingMissions  []scenario.MissionStruct
	routeIds           map[string]int32
	stopIds            map[string]int32
	startTimestamp     int64
	missionTimer       CancelableTimer
	client             *http_client.Client
	carId              *int32
}

func New(simulationScenario scenario.Scenario, loop bool, client *http_client.Client) *Simulation {
	simulation := new(Simulation)
	simulation.simulationScenario = simulationScenario
	simulation.remainingMissions = simulationScenario.Missions
	simulation.loop = loop
	simulation.client = client
	simulation.routeIds = make(map[string]int32)
	simulation.stopIds = make(map[string]int32)
	simulation.missionTimer = CancelableTimer{timer: nil, cancelTimer: make(chan struct{})}

	simulation.initDatabase(simulationScenario)
	return simulation
}

func (simulation *Simulation) SetCarId(carId *int32) {
	simulation.carId = carId
}

func (simulation *Simulation) Start() {
	if simulation.carId == nil {
		log.Fatalf("[ERROR] Car ID is not set for: %v", simulation.simulationScenario.CarId)
	}
	log.Printf("[INFO] [%v] Starting simulation", simulation.simulationScenario.CarId)
	simulation.resetSimulation()
}

func (simulation *Simulation) setNextMissionTimer() {
	if len(simulation.remainingMissions) > 0 {
		missionTimeOffset, err := strconv.ParseInt(simulation.remainingMissions[0].Timestamp, 10, 64)
		if err != nil {
			log.Printf("[INFO] [%v] Next mission (%v) timestamp has wrong format(%v), defaulting to 1 minute", simulation.simulationScenario.CarId, simulation.remainingMissions[0].Name, simulation.remainingMissions[0].Timestamp)
			missionTimeOffset = 60
		}
		startNextMissionTimestamp := missionTimeOffset + simulation.startTimestamp
		calculatedTimerTime := startNextMissionTimestamp - time.Now().Unix()
		if calculatedTimerTime < 1 {
			log.Printf("[WARNING] [%v] Calculated time to next mission (%v) seems wrong, defaulting to one minute", simulation.simulationScenario.CarId, calculatedTimerTime)
			calculatedTimerTime = 60
		}
		log.Printf("[INFO] [%v] Next mission (%v) timestamp: %v, mission will start in %vs", simulation.simulationScenario.CarId, simulation.remainingMissions[0].Name, startNextMissionTimestamp, calculatedTimerTime)
		simulation.startMissionTimer(int(calculatedTimerTime))
	}
}

func (simulation *Simulation) startMissionTimer(duration int) {
	if simulation.missionTimer.timer == nil {
		simulation.missionTimer.timer = time.NewTimer(time.Duration(duration) * time.Second)

	} else {
		simulation.missionTimer.timer.Reset(time.Duration(duration) * time.Second)
	}

	go func() {
		select {
		case <-simulation.missionTimer.timer.C:
			if len(simulation.currentMission.Stops) > 0 {
				log.Printf("[WARNING] [%v] Mission (%v) timeout! Starting new mission (%v)\n", simulation.simulationScenario.CarId, simulation.currentMission.Name, simulation.remainingMissions[0].Name)
			} else {
				log.Printf("[INFO] [%v] Starting next scheduled mission (%v)\n", simulation.simulationScenario.CarId, simulation.remainingMissions[0].Name)
			}
			simulation.popNextMission()
		case <-simulation.missionTimer.cancelTimer:
			simulation.missionTimer.timer = nil
		}
	}()
}

func (simulation *Simulation) resetSimulation() {
	simulation.remainingMissions = simulation.simulationScenario.Missions
	simulation.startTimestamp = time.Now().Unix()
	if len(simulation.remainingMissions) > 0 {
		simulation.setNextMissionTimer()
	}
}

func (simulation *Simulation) popNextMission() {
	if len(simulation.remainingMissions) > 0 {
		simulation.currentMission = simulation.remainingMissions[0]
		simulation.remainingMissions = simulation.remainingMissions[1:]
		simulation.postMissionOrders()
		simulation.setNextMissionTimer()
		//simulation.missionChanged = true // TODO delete old orders?
	}

	if simulation.loop {
		log.Printf("[INFO] [%v] Car have finished all of its missions, starting simulation again", simulation.simulationScenario.CarId)
		simulation.resetSimulation()
	} else {
		log.Printf("[INFO] [%v] All missions have been finished", simulation.simulationScenario.CarId) // TODO check order state first
		// TODO finish simulation
	}
}

func (simulation *Simulation) postMissionOrders() {
	for _, stop := range simulation.currentMission.Stops {
		simulation.client.AddOrder(*simulation.carId, simulation.stopIds[stop.Name], simulation.routeIds[simulation.currentMission.Route])
	}
}
