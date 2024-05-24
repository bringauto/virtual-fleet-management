package simulation

import (
	"log"
	"time"
	"virtual-fleet-management/pkg/scenario"
)

type CancelableTimer struct {
	timer       *time.Timer
	cancelTimer chan struct{}
	durationSec int
}

type MissionManager struct {
	currentMission    scenario.MissionStruct
	remainingMissions []scenario.MissionStruct
	missionTimer      CancelableTimer
	startTimestamp    int64
	carName           string
	orderManager      *OrderManager
}

func NewMissionManager(carName string, orderManager *OrderManager) *MissionManager {
	missionManager := new(MissionManager)
	missionManager.carName = carName
	missionManager.missionTimer = CancelableTimer{timer: nil, cancelTimer: make(chan struct{})}
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
		if calculatedTimerTime < 1 {
			log.Printf("[WARNING] [%v] Calculated time to next mission (%v) seems wrong, defaulting to one minute", missionManager.carName, calculatedTimerTime)
			calculatedTimerTime = 60
		}
		log.Printf("[INFO] [%v] Next mission (%v) timestamp: %v, mission will start in %vs", missionManager.carName, missionManager.remainingMissions[0].Name, startNextMissionTimestamp, calculatedTimerTime)
		missionManager.startMissionTimer(int(calculatedTimerTime))
	}
}

func (missionManager *MissionManager) startMissionTimer(duration int) {
	if missionManager.missionTimer.timer == nil {
		missionManager.missionTimer.timer = time.NewTimer(time.Duration(duration) * time.Second)

	} else {
		missionManager.missionTimer.timer.Reset(time.Duration(duration) * time.Second)
	}

	go func() {
		select {
		case <-missionManager.missionTimer.timer.C:
			if len(missionManager.currentMission.Stops) > 0 {
				log.Printf("[WARNING] [%v] Mission (%v) timeout! Starting new mission (%v)\n", missionManager.carName, missionManager.currentMission.Name, missionManager.remainingMissions[0].Name)
			} else {
				log.Printf("[INFO] [%v] Starting next scheduled mission (%v)\n", missionManager.carName, missionManager.remainingMissions[0].Name)
			}
			missionManager.popNextMission()
		case <-missionManager.missionTimer.cancelTimer:
			missionManager.missionTimer.timer = nil
		}
	}()
}

func (missionManager *MissionManager) popNextMission() {
	if len(missionManager.remainingMissions) > 0 {
		missionManager.currentMission = missionManager.remainingMissions[0]
		missionManager.remainingMissions = missionManager.remainingMissions[1:]
		missionManager.orderManager.cancelRemainingOrdersSince(missionManager.startTimestamp, missionManager.carName)
		for _, stop := range missionManager.currentMission.Stops {
			missionManager.orderManager.postOrder(stop.Name, missionManager.currentMission.Route)

		}
		missionManager.setNextMissionTimer()
	}
}
