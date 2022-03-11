package tool

import (
	. "github.com/takoyaki-3/go-gtfs"
)

func GetTrip(g *GTFS, tripId string) Trip {
	for _, trip := range g.Trips {
		if tripId == trip.ID {
			return trip
		}
	}
	return Trip{}
}

func GetRoute(g *GTFS, routeId string) Route {
	for _, route := range g.Routes {
		if routeId == route.ID {
			return route
		}
	}
	return Route{}
}

func GetHeadSign(g *GTFS, tripId string, stopId string) string {
	trip := GetTrip(g, tripId)
	if trip.Headsign != "" {
		return trip.Headsign
	}
	for _, stopTime := range g.StopsTimes {
		if stopTime.TripID == tripId && stopTime.StopID == stopId {
			return stopTime.StopHeadSign
		}
	}
	return ""
}
