package gtfs

import (
	"sort"
	"strconv"
)

// GTFS -
type GTFS struct {
	Path          string // The path to the containing directory
	Agency        Agency
	Agencies      []Agency
	Routes        []Route
	Stops         []Stop
	StopsTimes    []StopTime
	Trips         []Trip
	Calendars     []Calendar
	CalendarDates []CalendarDate
	Transfers     []Transfer
}

// Route -
type Route struct {
	ID        string `csv:"route_id"`
	AgencyID  string `csv:"agency_id"`
	ShortName string `csv:"route_short_name"`
	LongName  string `csv:"route_long_name"`
	Type      int    `csv:"route_type"`
	Desc      string `csv:"route_url"`
	URL       string `csv:"route_desc"`
	Color     string `csv:"route_color"`
	TextColor string `csv:"route_text_color"`
}

// Trip -
type Trip struct {
	ID          string `csv:"trip_id"`
	Name        string `csv:"trip_short_name"`
	RouteID     string `csv:"route_id"`
	ServiceID   string `csv:"service_id"`
	ShapeID     string `csv:"shape_id"`
	DirectionID string `csv:"direction_id"`
	Headsign    string `csv:"trip_headsign"`
}

// Stop -
type Stop struct {
	ID          string  `csv:"stop_id"`
	Code        string  `csv:"stop_code"`
	Name        string  `csv:"stop_name"`
	Description string  `csv:"stop_desc"`
	Latitude    float64 `csv:"stop_lat"`
	Longitude   float64 `csv:"stop_lon"`
	ZoneID			string  `csv:"zone_id"`
	Type        string  `csv:"location_type"`
	Parent      string  `csv:"parent_station"`
}

// StopTime -
type StopTime struct {
	StopID       string  `csv:"stop_id"`
	StopSeq      string  `csv:"stop_sequence"`
	StopHeadSign string  `csv:"stop_headsign"`
	TripID       string  `csv:"trip_id"`
	Shape        int 		 `csv:"shape_dist_traveled"`
	Departure    string  `csv:"departure_time"`
	Arrival      string  `csv:"arrival_time"`
	PickupType   int  	 `csv:"pickup_type"`
	DropOffType  int  	 `csv:"drop_off_type"`
}

// Calendar -
type Calendar struct {
	ServiceID string `csv:"service_id"`
	Monday    int    `csv:"monday"`
	Tuesday   int    `csv:"tuesday"`
	Wednesday int    `csv:"wednesday"`
	Thursday  int    `csv:"thursday"`
	Friday    int    `csv:"friday"`
	Saturday  int    `csv:"saturday"`
	Sunday    int    `csv:"sunday"`
	Start     string `csv:"start_date"`
	End       string `csv:"end_date"`
}

// CalendarDate -
type CalendarDate struct {
	ServiceID     string `csv:"service_id"`
	Date          string `csv:"date"`
	ExceptionType int    `csv:"exception_type"`
}

// Transfer -
type Transfer struct {
	FromStopID string `csv:"from_stop_id"`
	ToStopID   string `csv:"to_stop_id"`
	Type       int    `csv:"transfer_type"`
	MinTime    int    `csv:"min_transfer_time"`
}

// Agency -
type Agency struct {
	ID       string `csv:"agency_id"`
	Name     string `csv:"agency_name"`
	URL      string `csv:"agency_url"`
	Timezone string `csv:"agency_timezone"`
	Langue   string `csv:"agency_lang"`
	Phone    string `csv:"agency_phone"`
}

func (g *GTFS)Sort(){
	sort.Slice(g.Agencies,func(i, j int) bool {
		return g.Agencies[i].ID < g.Agencies[j].ID
	})
	sort.Slice(g.CalendarDates,func(i, j int) bool {
		iv := g.CalendarDates[i]
		jv := g.CalendarDates[j]
		iStr := iv.Date + ":" + iv.ServiceID + ":" + strconv.Itoa(iv.ExceptionType)
		jStr := jv.Date + ":" + jv.ServiceID + ":" + strconv.Itoa(jv.ExceptionType)
		return iStr < jStr
	})
	sort.Slice(g.Calendars,func(i, j int) bool {
		return g.Calendars[i].ServiceID < g.Calendars[j].ServiceID
	})
	sort.Slice(g.Routes,func(i, j int) bool {
		return g.Routes[i].ID < g.Routes[j].ID
	})
	sort.Slice(g.Stops,func(i, j int) bool {
		return g.Stops[i].ID < g.Stops[j].ID
	})
	sort.Slice(g.StopsTimes,func(i, j int) bool {
		iv := g.StopsTimes[i]
		jv := g.StopsTimes[j]
		iStr := iv.StopID + ":" + iv.TripID + ":" + iv.Departure
		jStr := jv.StopID + ":" + jv.TripID + ":" + jv.Departure
		return iStr < jStr
	})
	sort.Slice(g.Transfers,func(i, j int) bool {
		iv := g.Transfers[i]
		jv := g.Transfers[j]
		iStr := iv.FromStopID + iv.ToStopID + strconv.Itoa(iv.Type)
		jStr := jv.FromStopID + jv.ToStopID + strconv.Itoa(jv.Type)
		return iStr < jStr
	})
	sort.Slice(g.Trips,func(i, j int) bool {
		return g.Trips[i].ID < g.Trips[j].ID
	})
}