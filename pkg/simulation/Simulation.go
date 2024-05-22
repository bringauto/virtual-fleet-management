package simulation

import (
	"log"
	"virtual-fleet-management/pkg/http"
	"virtual-fleet-management/pkg/scenario"
)

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

func (simulation *Simulation) Start() {
	if simulation.orderManager.carId == nil {
		log.Fatalf("[ERROR] Car ID is not set for: %v", simulation.simulationScenario.CarId)
	}
	log.Printf("[INFO] [%v] Starting simulation", simulation.simulationScenario.CarId)
	simulation.resetSimulation()

	// TODO sleep until the mission is finished
	if simulation.loop {
		log.Printf("[INFO] [%v] Car have finished all of its missions, starting simulation again", simulation.simulationScenario.CarId)
		simulation.resetSimulation()
	} else {
		log.Printf("[INFO] [%v] All missions have been finished", simulation.simulationScenario.CarId) // TODO check order state first
		// TODO finish simulation
	}
}

func (simulation *Simulation) resetSimulation() {
	simulation.missionManager.startMissions(simulation.simulationScenario.Missions)

}
