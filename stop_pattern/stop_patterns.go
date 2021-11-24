package stoppattern

import (
	"sort"

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
		sort.Slice(trip,func(i, j int) bool {
			return trip[i].StopTimes[0].Departure < trip[j].StopTimes[0].Departure
		})
		patterns = append(patterns, RoutePattern{
			Trips: trip,
		})
	}

	return patterns
}