package simulation

import (
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"log"
	"virtual-fleet-management/pkg/http"
	"virtual-fleet-management/pkg/scenario"
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

func (orderManager *OrderManager) postMissionOrders(mission scenario.MissionStruct) {
	orderStops := make([]int32, len(mission.Stops))
	for i, stop := range mission.Stops {
		orderStops[i] = orderManager.stopIds[stop.Name]
	}
	orderManager.client.AddOrders(*orderManager.carId, orderStops, orderManager.routeIds[mission.Route])
}

func (orderManager *OrderManager) postOrder(stopName string, routeName string) {
	orderManager.client.AddOrders(*orderManager.carId, []int32{orderManager.stopIds[stopName]}, orderManager.routeIds[routeName])
}

func (orderManager *OrderManager) cancelRemainingOrders(carName string) {
	orders := orderManager.client.GetOrdersForCar(*orderManager.carId)
	var ordersToCancel []int32
	for _, order := range orders {
		if order.LastState.Status != openapi.DONE && order.LastState.Status != openapi.CANCELED {
			ordersToCancel = append(ordersToCancel, *order.Id)
		}
	}
	if len(ordersToCancel) > 0 {
		log.Printf("[INFO] [%v] cancelling orders id: %v", carName, ordersToCancel) // TODO do reverse lookup in map for stopName?
		orderManager.client.CancelOrders(ordersToCancel)
	}
}

func (orderManager *OrderManager) AreAllCarOrdersDone() bool {
	orders := orderManager.client.GetOrdersForCar(*orderManager.carId)
	for _, order := range orders {
		if !isOrderDone(order) {
			return false
		}
	}
	return true
}

func isOrderDone(order openapi.Order) bool {
	return order.LastState.Status == openapi.DONE || order.LastState.Status == openapi.CANCELED
}
