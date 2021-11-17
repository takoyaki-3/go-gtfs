package stoppattern

import (
	. "github.com/takoyaki-3/go-gtfs"
	. "github.com/takoyaki-3/go-gtfs/trip_timetable"
)

type RoutePattern struct {
	Trips []TripTimetable
	// Stops []Stop
	// Route Route
	// RouteId string
}

func GetRoutePatterns(g *GTFS)(patterns []RoutePattern){

	tripTimetables := GetTripTimetables(g)

	routePatterns := map[string][]TripTimetable{}

	for _,trip := range tripTimetables{
		stopPattern := trip.Properties.RouteID
		for _,stop := range trip.StopTimes {
			stopPattern += ":" + stop.StopID
		}
		routePatterns[stopPattern] = append(routePatterns[stopPattern], trip)
	}

	for _,trip:=range routePatterns {
		patterns = append(patterns, RoutePattern{
			Trips: trip,
		})
	}

	return patterns
}