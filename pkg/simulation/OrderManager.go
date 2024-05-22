package simulation

import (
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"virtual_fleet_management/pkg/http_client"
)

type OrderManager struct {
	client   *http_client.Client
	routeIds map[string]int32
	stopIds  map[string]int32
	carId    *int32
}

func newOrderManager(client *http_client.Client, routeIds map[string]int32, stopIds map[string]int32) *OrderManager {
	return &OrderManager{
		client:   client,
		routeIds: routeIds,
		stopIds:  stopIds,
	}
}

func (simulation *OrderManager) SetCarId(carId *int32) {
	simulation.carId = carId
}

func (simulation *OrderManager) postOrder(stopName string, routeName string) {
	simulation.client.AddOrder(*simulation.carId, simulation.stopIds[stopName], simulation.routeIds[routeName])
}

func (simulation *OrderManager) deleteRemainingOrdersSince(since int64) {
	orders := simulation.client.GetOrdersForCar(*simulation.carId, since)
	for _, order := range orders {
		if order.LastState.Status != openapi.DONE && order.LastState.Status != openapi.CANCELED {
			//log.Printf("[INFO] [%v] cancelling order: %v", simulation.simulationScenario.CarId, order.Id) TODO move
			simulation.client.CancelOrder(*order.Id)
		}
	}
}
