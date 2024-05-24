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

func (orderManager *OrderManager) SetCarId(carId *int32) {
	orderManager.carId = carId
}

func (orderManager *OrderManager) postOrder(stopName string, routeName string) {
	orderManager.client.AddOrder(*orderManager.carId, orderManager.stopIds[stopName], orderManager.routeIds[routeName])
}

func (orderManager *OrderManager) cancelRemainingOrdersSince(since int64, carName string) {
	orders := orderManager.client.GetOrdersForCar(*orderManager.carId, since)
	for _, order := range orders {
		if order.LastState.Status != openapi.DONE && order.LastState.Status != openapi.CANCELED {
			log.Printf("[INFO] [%v] cancelling order id: %v", carName, *order.Id) // TODO do reverse lookup in map for stopName?
			orderManager.client.CancelOrder(*order.Id)
		}
	}
}

func (orderManager *OrderManager) AreAllCarOrdersDone() bool {
	orders := orderManager.client.GetOrdersForCar(*orderManager.carId, 0)
	for _, order := range orders {
		if order.LastState.Status != openapi.DONE && order.LastState.Status != openapi.CANCELED {
			return false
		}
	}
	return true
}
