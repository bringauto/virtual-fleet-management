package http

import (
	"context"
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"log"
	"net/url"
	"virtual-fleet-management/pkg/scenario"
)

type Client struct {
	apiClient *openapi.APIClient
	auth      context.Context
	userId    int32
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
	if err != nil { // TODO should loop for a while?
		log.Fatal("[ERROR] ", err)
	}
	// TODO: get userId from management api
	return &Client{
		apiClient: apiClient,
		auth:      auth,
		userId:    1, // TODO not implemented in management api yet
	}
}

// convertPositionToGnss Convert scenario.Position to openapi.GNSSPosition
func convertPositionToGnss(position scenario.Position) openapi.GNSSPosition {
	gnss := openapi.NewGNSSPosition()
	gnss.Latitude = &position.Latitude
	gnss.Longitude = &position.Longitude
	gnss.Altitude = &position.Altitude
	return *gnss
}

func (c *Client) GetStops() []openapi.Stop {
	stopData, _, err := c.apiClient.StopAPI.GetStops(c.auth).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'StopAPI.GetStops': `, err)
	}
	return stopData
}

// AddStop Add stop to database and return stopId
func (c *Client) AddStop(stationStruct scenario.StationStruct) (stopId *int32) {
	newStop := openapi.NewStop(stationStruct.Name, convertPositionToGnss(stationStruct.Position))
	stopData, _, err := c.apiClient.StopAPI.CreateStop(c.auth).Stop(*newStop).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'StopApi.CreateStop' with stop: '`, newStop.Name, `' error: `, err)
	}
	return stopData.Id
}

func (c *Client) GetRoutes() []openapi.Route {
	routeData, _, err := c.apiClient.RouteAPI.GetRoutes(c.auth).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'RouteAPI.GetRoutes': `, err)
	}
	return routeData
}

func (c *Client) AddRoute(route *openapi.Route) (routeId *int32) {
	routeData, _, err := c.apiClient.RouteAPI.CreateRoute(c.auth).Route(*route).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'RouteAPI.CreateRoute':`, err)
	}
	return routeData.Id
}

func (c *Client) GetCars() []openapi.Car {
	carData, _, err := c.apiClient.CarAPI.GetCars(c.auth).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'CarAPI.GetCars': `, err)
	}
	return carData
}

func (c *Client) AddOrder(carId int32, stopId int32, routeId int32) {
	order := openapi.NewOrder(c.userId, carId, stopId, routeId)
	_, _, err := c.apiClient.OrderAPI.CreateOrder(c.auth).Order(*order).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'OrderAPI.CreateOrder': `, err)
	}
}

func (c *Client) GetOrdersForCar(carId int32, since int64) []openapi.Order {
	orders, _, err := c.apiClient.OrderAPI.GetCarOrders(c.auth, carId).Since(since).Execute()
	if err != nil {
		log.Fatal(`[ERROR] calling 'OrderAPI.GetCarOrders': `, err)

	}
	return orders
}

func (c *Client) CancelOrder(orderId int32) {
	orderStatus := *openapi.NewOrderState(openapi.CANCELED, orderId)
	s, r, err := c.apiClient.OrderStateAPI.CreateOrderState(c.auth).OrderState(orderStatus).Execute()
	if err != nil {
		log.Fatal(`[ERROR] cancelling order with 'OrderStateAPI.CreateOrderState': `, err, r, s)
	}
}
