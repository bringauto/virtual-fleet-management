package simulation

import (
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"log"
	"math"
	"reflect"
	"virtual-fleet-management/pkg/http"
	"virtual-fleet-management/pkg/scenario"
)

const gpsEqualityThreshold = 1e-6

func gpsEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) <= gpsEqualityThreshold
}

func isPositionEqual(position1 scenario.Position, position2 openapi.GNSSPosition) bool {
	return gpsEqual(position1.Latitude, *position2.Latitude) && gpsEqual(position1.Longitude, *position2.Longitude)
}
func findStationId(station scenario.StationStruct, existingStations []openapi.Stop) *int32 {
	for _, existingStation := range existingStations {
		if station.Name == existingStation.Name {
			if isPositionEqual(station.Position, existingStation.Position) {
				return existingStation.Id
			} else {
				log.Fatalf("[ERROR] Station %v already exists, but with different position", station.Name)
				return nil
			}
		}
	}
	return nil
}

func findRouteId(route *openapi.Route, existingRoutes []openapi.Route) *int32 {
	for _, existingRoute := range existingRoutes {
		if route.Name == existingRoute.Name {
			if reflect.DeepEqual(route.StopIds, existingRoute.StopIds) {
				return existingRoute.Id
			} else {
				log.Fatalf("[ERROR] Route %v already exists, but with different stops", route.Name)
				return nil
			}
		}
	}
	return nil
}

// convertPositionToGnss Convert scenario.Position to openapi.GNSSPosition
func convertPositionToGnss(position scenario.Position) openapi.GNSSPosition {
	gnss := openapi.NewGNSSPosition()
	gnss.Latitude = &position.Latitude
	gnss.Longitude = &position.Longitude
	gnss.Altitude = &position.Altitude
	return *gnss
}

func (simulation *Simulation) initDatabase(scenario2 scenario.Scenario, client *http.Client) (map[string]int32, map[string]int32) {
	stopIds := make(map[string]int32)
	routeIds := make(map[string]int32)
	existingStations := client.GetStops()
	for _, route := range scenario2.Routes {
		var routeStopIds []int32
		for _, station := range route.Stations {
			stationId := findStationId(station, existingStations)
			if stationId == nil {
				newStop := openapi.NewStop(station.Name, convertPositionToGnss(station.Position))
				stationId = client.AddStop(newStop)
				newStop.Id = stationId
				existingStations = append(existingStations, *newStop)
			}
			routeStopIds = append(routeStopIds, *stationId)
			stopIds[station.Name] = *stationId
		}
		newRoute := openapi.NewRoute(route.Name)
		newRoute.SetStopIds(routeStopIds)
		routeId := findRouteId(newRoute, client.GetRoutes())
		if routeId == nil {
			routeId = client.AddRoute(newRoute)
		}
		routeIds[route.Name] = *routeId
	}
	return routeIds, stopIds
}
