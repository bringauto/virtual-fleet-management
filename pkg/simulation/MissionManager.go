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
		missionTimeOffset := int64(missionManager.remainingMissions[0].DelaySeconds)
		startNextMissionTimestamp := missionTimeOffset + missionManager.startTimestamp
		calculatedTimerTime := startNextMissionTimestamp - time.Now().Unix()
		if calculatedTimerTime < 0 {
			log.Printf("[WARNING] [%v] Calculated time to next mission (%v) seems wrong, defaulting to one minute", missionManager.carName, calculatedTimerTime)
			calculatedTimerTime = 60
		}
		log.Printf("[INFO] [%v] Next mission (%v) timestamp: %v, mission will start in %vs", missionManager.carName, missionManager.remainingMissions[0].Name, startNextMissionTimestamp, calculatedTimerTime)
		missionManager.startMissionTimer(int(calculatedTimerTime))
	}
}

func (missionManager *MissionManager) startMissionTimer(duration int) {
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
		//startInMilli := missionManager.startTimestamp * 1000
		missionManager.orderManager.cancelRemainingOrdersSince(0, missionManager.carName)
		for _, stop := range missionManager.currentMission.Stops {
			missionManager.orderManager.postOrder(stop.Name, missionManager.currentMission.Route)

		}
		missionManager.setNextMissionTimer()
	}
}
