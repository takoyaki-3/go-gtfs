package gtfs

import (
	"sort"
	"strconv"
)

type RoutePattern struct {
	Trips []TripTimetable
	// Stops []Stop
	// Route Route
	// RouteId string
}

func (g *GTFS) GetRoutePatterns() (patterns []RoutePattern) {

	tripTimetables := GetTripTimetables(g)

	routePatterns := map[string][]TripTimetable{}

	for _, trip := range tripTimetables {
		stopPattern := trip.Properties.RouteID
		for _, stop := range trip.StopTimes {
			stopPattern += ":" + stop.StopID
		}
		routePatterns[stopPattern] = append(routePatterns[stopPattern], trip)
	}

	for _, trip := range routePatterns {
		sort.Slice(trip, func(i, j int) bool {
			return trip[i].StopTimes[0].Departure < trip[j].StopTimes[0].Departure
		})
		patterns = append(patterns, RoutePattern{
			Trips: trip,
		})
	}

	return patterns
}

type TripTimetable struct {
	StopTimes  []StopTime
	Properties TimetableEdgeProperty
}

func GetTripTimetables(g *GTFS) (tripTimetables []TripTimetable) {

	trips := map[string][]StopTime{}

	for _, stopTime := range g.StopsTimes {
		trips[stopTime.TripID] = append(trips[stopTime.TripID], stopTime)
	}
	routes := map[string]Route{}
	for _, route := range g.Routes {
		routes[route.ID] = route
	}
	Properties := map[string]TimetableEdgeProperty{}
	for _, trip := range g.Trips {
		route := routes[trip.RouteID]
		Properties[trip.ID] = TimetableEdgeProperty{
			TripID:      trip.ID,
			Name:        trip.Name,
			RouteID:     trip.RouteID,
			ServiceID:   trip.ServiceID,
			ShapeID:     trip.ShapeID,
			DirectionID: trip.DirectionID,
			Headsign:    trip.Headsign,
			AgencyID:    route.AgencyID,
			ShortName:   route.ShortName,
			LongName:    route.LongName,
			Type:        route.Type,
			Desc:        route.Desc,
			URL:         route.URL,
			Color:       route.Color,
			TextColor:   route.TextColor,
		}
	}

	for tripId, stopTimes := range trips {
		sort.Slice(trips[tripId], func(i, j int) bool {
			is, _ := strconv.Atoi(trips[tripId][i].StopSeq)
			js, _ := strconv.Atoi(trips[tripId][j].StopSeq)
			if is == js {
				return trips[tripId][i].Departure < trips[tripId][j].Departure
			}
			return is < js
		})
		tripTimetables = append(tripTimetables, TripTimetable{
			Properties: Properties[tripId],
			StopTimes:  stopTimes,
		})
	}

	return tripTimetables
}
