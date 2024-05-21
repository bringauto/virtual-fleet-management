package simulation

import (
	openapi "github.com/bringauto/fleet-management-http-client-go"
	"log"
	"math"
	"reflect"
	"virtual_fleet_management/pkg/scenario"
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
func (simulation *Simulation) initDatabase(scenario2 scenario.Scenario) {
	existingStations := simulation.client.GetStops()
	for _, route := range scenario2.Routes {
		var stopIds []int32
		for _, station := range route.Stations {
			stationId := findStationId(station, existingStations)
			if stationId == nil {
				stationId = simulation.client.AddStop(station)
			}
			stopIds = append(stopIds, *stationId)
			simulation.stopIds[station.Name] = *stationId

		}
		newRoute := openapi.NewRoute(route.Name)
		newRoute.SetStopIds(stopIds)
		routeId := findRouteId(newRoute, simulation.client.GetRoutes())
		if routeId == nil {
			routeId = simulation.client.AddRoute(newRoute)
		}
		simulation.routeIds[route.Name] = *routeId
	}
}
