package edgetimetable

import (
	"sort"

	"github.com/takoyaki-3/go-gtfs"
)

func GTFS2TimeTableEdges(g *gtfs.GTFS)(et *EdgeTimetable){

	trips := map[string][]gtfs.StopTime{}

	for _,stopTime := range g.StopsTimes{
		trips[stopTime.TripID] = append(trips[stopTime.TripID], stopTime)
	}
	for tripId,_ := range trips{
		sort.Slice(trips[tripId],func(i, j int) bool {
			return trips[tripId][i].Departure < trips[tripId][j].Departure
		})
		for i:=1;i<len(trips[tripId]);i++{
			et.Edges = append(et.Edges, TimetableEdge{
				TripId: tripId,
				DepartureTime: trips[tripId][i-1].Departure,
				ArrivalTime: trips[tripId][i].Arrival,
				FromStop: trips[tripId][i-1].Departure,
				ToStop: trips[tripId][i].Arrival,
				StopHeadSign: trips[tripId][i-1].StopHeadSign,
				// PickupType: trips[tripId][i-1].,
			})
		}
	}
	routes := map[string]gtfs.Route{}
	for _,route := range g.Routes{
		routes[route.ID] = route
	}
	for _,trip := range g.Trips{
		route := routes[trip.RouteID]
		et.Properties = append(et.Properties, TimetableEdgeProperty{
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
		})
	}

	return et
}
