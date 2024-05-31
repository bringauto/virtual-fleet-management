package simulation

import (
	"log"
	"sync"
	"time"
	"virtual-fleet-management/pkg/http"
	"virtual-fleet-management/pkg/scenario"
)

const sleepTime = 10

type Simulation struct {
	loop               bool
	simulationScenario scenario.Scenario
	missionManager     *MissionManager
	orderManager       *OrderManager
}

func New(simulationScenario scenario.Scenario, loop bool, client *http.Client) *Simulation {
	simulation := new(Simulation)
	simulation.simulationScenario = simulationScenario
	simulation.loop = loop
	routeIds, stopIds := simulation.initDatabase(simulationScenario, client)
	simulation.orderManager = newOrderManager(client, routeIds, stopIds)
	simulation.missionManager = NewMissionManager(simulationScenario.CarId, simulation.orderManager)
	return simulation
}

// SetCarId sets the car ID for the simulation, the carId is taken from management API database
func (simulation *Simulation) SetCarId(carId *int32) {
	simulation.orderManager.SetCarId(carId)
}

func (simulation *Simulation) Start(wg *sync.WaitGroup) {
	if simulation.orderManager.carId == nil {
		log.Fatalf("[ERROR] Car ID is not set for: %v", simulation.simulationScenario.CarId)
	}
	log.Printf("[INFO] [%v] Starting simulation", simulation.simulationScenario.CarId)
	finished := false
	for !finished {
		simulation.resetSimulation()
		// sleep until the minimum delay of mission + sleepTime to avoid race condition
		time.Sleep(time.Duration(simulation.simulationScenario.GetTotalDelay()+sleepTime) * time.Second)
		for {
			if simulation.orderManager.AreAllCarOrdersDone() {
				break
			}
			time.Sleep(sleepTime * time.Second)
		}
		if simulation.loop {
			log.Printf("[INFO] [%v] Car have finished all of its missions, starting simulation again", simulation.simulationScenario.CarId)
		} else {
			finished = true
		}
	}
	log.Printf("[INFO] [%v] All missions have been finished", simulation.simulationScenario.CarId)
	defer wg.Done()
}

func (simulation *Simulation) resetSimulation() {
	simulation.getCarToStartingState()
	simulation.missionManager.startMissions(simulation.simulationScenario.Missions)
}

func (simulation *Simulation) getCarToStartingState() {
	log.Printf("[INFO] [%v] Cancelling all remaining orders", simulation.simulationScenario.CarId)
	simulation.orderManager.cancelRemainingOrdersSince(0, simulation.simulationScenario.CarId)
	log.Printf("[INFO] [%v] Ordering car to the starting station", simulation.simulationScenario.CarId)
	simulation.orderManager.postOrder(simulation.simulationScenario.StartingStation, simulation.simulationScenario.StartingRoute)
}
