package triptimetable

import (
	"sort"
	"github.com/takoyaki-3/go-gtfs"
	"github.com/takoyaki-3/go-gtfs/edgetimetable"
)

type TripTimetable struct {	
	StopTimes []gtfs.StopTime
	Properties edgetimetable.TimetableEdgeProperty
}

func GetTripTimetables(g gtfs.GTFS)(tripTimetables []TripTimetable){

	trips := map[string][]gtfs.StopTime{}

	for _,stopTime := range g.StopsTimes{
		trips[stopTime.TripID] = append(trips[stopTime.TripID], stopTime)
	}
	routes := map[string]gtfs.Route{}
	for _,route := range g.Routes{
		routes[route.ID] = route
	}
	Properties := map[string]edgetimetable.TimetableEdgeProperty{}
	for _,trip := range g.Trips{
		route := routes[trip.RouteID]
		Properties[trip.ID] = edgetimetable.TimetableEdgeProperty{
			TripID: trip.ID,
			Name: trip.Name,
			RouteID: trip.RouteID,
			ServiceID: trip.ServiceID,
			ShapeID: trip.ShapeID,
			DirectionID: trip.DirectionID,
			Headsign: trip.Headsign,
			AgencyID: route.AgencyID,
			ShortName: route.ShortName,
			LongName: route.LongName,
			Type: route.Type,
			Desc: route.Desc,
			URL: route.URL,
			Color: route.Color,
			TextColor: route.TextColor,
		}
	}

	for tripId,stopTimes := range trips{
		sort.Slice(trips[tripId],func(i, j int) bool {
			return trips[tripId][i].Departure < trips[tripId][j].Departure
		})
		tripTimetables = append(tripTimetables,TripTimetable{
			Properties: Properties[tripId],
			StopTimes: stopTimes,
		})
	}

	return tripTimetables
}
