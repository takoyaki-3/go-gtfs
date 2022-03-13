package gtfs

import (
	"fmt"
	"path"
	"sort"

	csvtag "github.com/takoyaki-3/go-csv-tag/v3"
)

func (g *GTFS) GTFS2TimeTableEdges() (et *EdgeTimetable) {
	et = &EdgeTimetable{}

	trips := map[string][]StopTime{}

	for _, stopTime := range g.StopsTimes {
		trips[stopTime.TripID] = append(trips[stopTime.TripID], stopTime)
	}
	for tripId, _ := range trips {
		sort.Slice(trips[tripId], func(i, j int) bool {
			return trips[tripId][i].Departure < trips[tripId][j].Departure
		})
		for i := 1; i < len(trips[tripId]); i++ {
			et.Edges = append(et.Edges, TimetableEdge{
				TripId:        tripId,
				DepartureTime: trips[tripId][i-1].Departure,
				ArrivalTime:   trips[tripId][i].Arrival,
				FromStop:      trips[tripId][i-1].StopID,
				ToStop:        trips[tripId][i].StopID,
				StopHeadSign:  trips[tripId][i-1].StopHeadSign,
				PickupType:    trips[tripId][i-1].PickupType,
				DropOffType:   trips[tripId][i].DropOffType,
			})
		}
	}
	routes := map[string]Route{}
	for _, route := range g.Routes {
		routes[route.ID] = route
	}
	for _, trip := range g.Trips {
		route := routes[trip.RouteID]
		et.Properties = append(et.Properties, TimetableEdgeProperty{
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
		})
	}
	et.Stops = g.Stops

	return et
}

func DumpEdgeTimetable(edgeTimetable *EdgeTimetable, dirPath string, filter map[string]bool) error {

	files := map[string]interface{}{
		"ptedges.csv":       edgeTimetable.Edges,
		"pteproperties.csv": edgeTimetable.Properties,
		"stops.csv":         edgeTimetable.Stops,
	}

	for file, src := range files {
		if filter != nil && !filter[file[:len(file)-4]] {
			continue
		}
		if src == nil {
			continue
		}
		filePath := path.Join(dirPath, file)

		err := csvtag.DumpToFile(src, filePath)
		if err != nil {
			return fmt.Errorf("Error dumping file %v: %v", file, err)
		}
	}
	return nil
}
