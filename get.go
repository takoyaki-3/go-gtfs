package gtfs

import (
	"errors"
)

func (g *GTFS) GetRoute(routeId string) Route {
	for _, r := range g.Routes {
		if r.ID == routeId {
			return r
		}
	}
	return Route{}
}

func (g *GTFS) GetTrip(tripId string) Trip {
	for _, t := range g.Trips {
		if t.ID == tripId {
			return t
		}
	}
	return Trip{}
}

func (g *GTFS) GetStop(stopID string) Stop {
	for _, s := range g.Stops {
		if s.ID == stopID {
			return s
		}
	}
	return Stop{}
}

func (g *GTFS) GetHeadSign(tripId string, stopId string) string {
	trip := g.GetTrip(tripId)
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

func (g *GTFS) GetFareAttribute(fareId string) (FareAttribute, error) {
	for _, v := range g.FareAttributes {
		if v.FareId == fareId {
			return v, nil
		}
	}
	return FareAttribute{}, errors.New("muched fare_attribute not found.")
}

func (g *GTFS) GetFareAttributeFromOD(originId string, destinationId string, routeId string) (FareAttribute, error) {

	for _, v := range g.FareRules {

		// 条件に一致するか判定
		isOriginOk := (originId == v.OriginId)
		isDestinationOk := (destinationId == v.DestinationId) || ("*" == v.DestinationId) || ("" == v.DestinationId)
		isRouteIdOk := (v.RouteId == routeId) || (v.RouteId == "*") || (v.RouteId == "")

		if isOriginOk && isDestinationOk && isRouteIdOk {
			return g.GetFareAttribute(v.FareId)
		}
	}

	return FareAttribute{}, errors.New("muched fare_rule not found.")
}
