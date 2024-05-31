package simulation

import (
	"log"
	"time"
	"virtual-fleet-management/pkg/scenario"
)

type MissionManager struct {
	currentMission    scenario.MissionStruct
	remainingMissions []scenario.MissionStruct
	missionTimer      *time.Timer
	startTimestamp    int64
	carName           string
	orderManager      *OrderManager
}

func NewMissionManager(carName string, orderManager *OrderManager) *MissionManager {
	missionManager := new(MissionManager)
	missionManager.carName = carName
	missionManager.orderManager = orderManager
	return missionManager
}

func (missionManager *MissionManager) isMissionFinished() bool {
	return len(missionManager.remainingMissions) == 0
}

func (missionManager *MissionManager) startMissions(missions []scenario.MissionStruct) {
	missionManager.remainingMissions = missions
	missionManager.startTimestamp = time.Now().Unix()
	if len(missionManager.remainingMissions) > 0 {
		missionManager.setNextMissionTimer()
	}
}

func (missionManager *MissionManager) setNextMissionTimer() {
	if len(missionManager.remainingMissions) > 0 {
		missionTimeOffset := missionManager.remainingMissions[0].DelaySeconds
		log.Printf("[INFO] [%v] Next mission (%v) will start in %vs", missionManager.carName, missionManager.remainingMissions[0].Name, missionTimeOffset)
		missionManager.startMissionTimer(missionTimeOffset)
	}
}

func (missionManager *MissionManager) startMissionTimer(duration int32) {
	if missionManager.missionTimer == nil {
		missionManager.missionTimer = time.NewTimer(time.Duration(duration) * time.Second)

	} else {
		missionManager.missionTimer.Reset(time.Duration(duration) * time.Second)
	}

	go func() {
		<-missionManager.missionTimer.C
		if len(missionManager.currentMission.Stops) > 0 {
			log.Printf("[WARNING] [%v] Mission (%v) timeout!\n", missionManager.carName, missionManager.currentMission.Name)
		}
		log.Printf("[INFO] [%v] Starting next scheduled mission (%v)\n", missionManager.carName, missionManager.remainingMissions[0].Name)
		missionManager.popNextMission()
	}()
}

func (missionManager *MissionManager) popNextMission() {
	if len(missionManager.remainingMissions) > 0 {
		missionManager.currentMission = missionManager.remainingMissions[0]
		missionManager.remainingMissions = missionManager.remainingMissions[1:]
		missionManager.orderManager.cancelRemainingOrders(missionManager.carName)
		missionManager.orderManager.postMissionOrders(missionManager.currentMission)

		missionManager.setNextMissionTimer()
	}
}
