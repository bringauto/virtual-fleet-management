package http

import (
	"context"
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"log"
	"net/url"
)

type Client struct {
	apiClient *openapi.APIClient
	auth      context.Context
}

// createConfiguration Create configuration for client
func createConfiguration(host string) *openapi.Configuration {
	u, err := url.Parse(host)
	if err != nil {
		log.Fatal("[ERROR] ", err)
	}
	config := openapi.NewConfiguration()
	config.Host = u.Host
	config.Scheme = u.Scheme
	return config
}

// CreateClient Create client for communication with server
func CreateClient(host string, key string) *Client {
	apiClient := openapi.NewAPIClient(createConfiguration(host))

	auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"APIKeyAuth": {Key: key},
		},
	)
	log.Printf("[INFO] Checking access to API '%v'", host)
	_, err := apiClient.ApiAPI.CheckApiIsAlive(auth).Execute()
	if err != nil {
		log.Fatal("[ERROR] ", err)
	}
	return &Client{
		apiClient: apiClient,
		auth:      auth,
	}
}

func (c *Client) GetStops() []openapi.Stop {
	stopData, _, err := c.apiClient.StopAPI.GetStops(c.auth).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'StopAPI.GetStops': `, err)
	}
	return stopData
}

// AddStop Add stop to database and return stopId
func (c *Client) AddStop(stop *openapi.Stop) (stopId *int32) {
	stopList := []openapi.Stop{*stop}
	stopData, _, err := c.apiClient.StopAPI.CreateStops(c.auth).Stop(stopList).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'StopApi.CreateStop' with stop: '`, stop.Name, `' error: `, err)
	}
	return stopData[0].Id
}

func (c *Client) GetRoutes() []openapi.Route {
	routeData, _, err := c.apiClient.RouteAPI.GetRoutes(c.auth).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'RouteAPI.GetRoutes': `, err)
	}
	return routeData
}

func (c *Client) AddRoute(route *openapi.Route) (routeId *int32) {
	routeList := []openapi.Route{*route}
	routeData, _, err := c.apiClient.RouteAPI.CreateRoutes(c.auth).Route(routeList).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'RouteAPI.CreateRoute':`, err)
	}
	return routeData[0].Id
}

func (c *Client) GetCars() []openapi.Car {
	carData, _, err := c.apiClient.CarAPI.GetCars(c.auth).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'CarAPI.GetCars': `, err)
	}
	return carData
}

func (c *Client) AddOrders(carId int32, stopIds []int32, routeId int32) {
	allOrders := make([]openapi.Order, len(stopIds))
	for i, stopId := range stopIds {
		allOrders[i] = *openapi.NewOrder(carId, stopId, routeId)
	}
	_, _, err := c.apiClient.OrderAPI.CreateOrders(c.auth).Order(allOrders).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'OrderAPI.CreateOrder': `, err)
	}
}

func (c *Client) GetOrdersForCar(carId int32) []openapi.Order {
	orders, _, err := c.apiClient.OrderAPI.GetCarOrders(c.auth, carId).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'OrderAPI.GetCarOrders': `, err)

	}
	return orders
}

func (c *Client) CancelOrders(orderIds []int32) {
	if len(orderIds) <= 0 {
		return
	}
	allOrderStates := make([]openapi.OrderState, len(orderIds))
	for i, orderId := range orderIds {
		allOrderStates[i] = *openapi.NewOrderState(openapi.CANCELED, orderId)
	}
	_, r, err := c.apiClient.OrderStateAPI.CreateOrderStates(c.auth).OrderState(allOrderStates).Execute()
	if err != nil {
		log.Fatal(`[ERROR] cancelling order with 'OrderStateAPI.CreateOrderState': `, r.Status, err)
	}
}
