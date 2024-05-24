package simulation

import (
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"log"
	"virtual-fleet-management/pkg/http"
)

type OrderManager struct {
	client   *http.Client
	routeIds map[string]int32
	stopIds  map[string]int32
	carId    *int32
}

func newOrderManager(client *http.Client, routeIds map[string]int32, stopIds map[string]int32) *OrderManager {
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

func (simulation *OrderManager) cancelRemainingOrdersSince(since int64, carName string) {
	orders := simulation.client.GetOrdersForCar(*simulation.carId, since)
	for _, order := range orders {
		if order.LastState.Status != openapi.DONE && order.LastState.Status != openapi.CANCELED {
			log.Printf("[INFO] [%v] cancelling order id: %v", carName, order.Id) // TODO do reverse lookup in map for stopName?
			simulation.client.CancelOrder(*order.Id)
		}
	}
}
