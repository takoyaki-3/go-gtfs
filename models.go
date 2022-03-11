package gtfs

import (
	"sort"
	"strconv"
)

// GTFS -
type GTFS struct {
	Path          string         `json:"path"` // The path to the containing directory
	Agency        Agency         `json:"agency"`
	Agencies      []Agency       `json:"agencies"`
	Routes        []Route        `json:"routes"`
	Stops         []Stop         `json:"stops"`
	StopsTimes    []StopTime     `json:"stop_times"`
	Trips         []Trip         `json:"trips"`
	Calendars     []Calendar     `json:"calendars"`
	CalendarDates []CalendarDate `json:"calendar_dates"`
	Transfers     []Transfer     `json:"transfer"`
}

// Route -
type Route struct {
	ID        string `csv:"route_id" json:"route_id"`
	AgencyID  string `csv:"agency_id" json:"agency_id"`
	ShortName string `csv:"route_short_name" json:"route_short_name"`
	LongName  string `csv:"route_long_name" json:"route_long_name"`
	Type      int    `csv:"route_type" json:"route_type"`
	Desc      string `csv:"route_url" json:"route_url"`
	URL       string `csv:"route_desc" json:"route_desc"`
	Color     string `csv:"route_color" json:"route_color"`
	TextColor string `csv:"route_text_color" json:"route_text_color"`
}

// Trip -
type Trip struct {
	ID          string `csv:"trip_id" json:"trip_id"`
	Name        string `csv:"trip_short_name" json:"trip_short_name"`
	RouteID     string `csv:"route_id" json:"route_id"`
	ServiceID   string `csv:"service_id" json:"service_id"`
	ShapeID     string `csv:"shape_id" json:"shape_id"`
	DirectionID string `csv:"direction_id" json:"direction_id"`
	Headsign    string `csv:"trip_headsign" json:"trip_headsign"`
}

// Stop -
type Stop struct {
	ID          string  `csv:"stop_id" json:"trip_id"`
	Code        string  `csv:"stop_code" json:"stop_code"`
	Name        string  `csv:"stop_name" json:"stop_name"`
	Description string  `csv:"stop_desc" json:"stop_desc"`
	Latitude    float64 `csv:"stop_lat" json:"stop_lat"`
	Longitude   float64 `csv:"stop_lon" json:"stop_lon"`
	ZoneID      string  `csv:"zone_id" json:"zone_id"`
	Type        string  `csv:"location_type" json:"location_type"`
	Parent      string  `csv:"parent_station" json:"parent_station"`
}

// StopTime -
type StopTime struct {
	StopID       string `csv:"stop_id" json:"stop_id"`
	StopSeq      string `csv:"stop_sequence" json:"stop_sequence"`
	StopHeadSign string `csv:"stop_headsign" json:"stop_headsign"`
	TripID       string `csv:"trip_id" json:"trip_id"`
	Shape        int    `csv:"shape_dist_traveled" json:"shape_dist_traveled"`
	Departure    string `csv:"departure_time" json:"departure_time"`
	Arrival      string `csv:"arrival_time" json:"arrival_time"`
	PickupType   int    `csv:"pickup_type" json:"pickup_type"`
	DropOffType  int    `csv:"drop_off_type" json:"drop_off_type"`
}

// Calendar -
type Calendar struct {
	ServiceID string `csv:"service_id" json:"trip_id"`
	Monday    int    `csv:"monday" json:"monday"`
	Tuesday   int    `csv:"tuesday" json:"tuesday"`
	Wednesday int    `csv:"wednesday" json:"wednesday"`
	Thursday  int    `csv:"thursday" json:"thursday"`
	Friday    int    `csv:"friday" json:"friday"`
	Saturday  int    `csv:"saturday" json:"saturday"`
	Sunday    int    `csv:"sunday" json:"sunday"`
	Start     string `csv:"start_date" json:"start_date"`
	End       string `csv:"end_date" json:"end_date"`
}

// CalendarDate -
type CalendarDate struct {
	ServiceID     string `csv:"service_id" json:"service_id"`
	Date          string `csv:"date" json:"date"`
	ExceptionType int    `csv:"exception_type" json:"exception_type"`
}

// Transfer -
type Transfer struct {
	FromStopID string `csv:"from_stop_id" json:"trip_id"`
	ToStopID   string `csv:"to_stop_id" json:"to_stop_id"`
	Type       int    `csv:"transfer_type" json:"transfer_type"`
	MinTime    int    `csv:"min_transfer_time" json:"min_transfer_time"`
}

// Agency -
type Agency struct {
	ID       string `csv:"agency_id" json:"trip_id"`
	Name     string `csv:"agency_name" json:"agency_name"`
	URL      string `csv:"agency_url" json:"agency_url"`
	Timezone string `csv:"agency_timezone" json:"agency_timezone"`
	Langue   string `csv:"agency_lang" json:"agency_lang"`
	Phone    string `csv:"agency_phone" json:"agency_phone"`
}

func (g *GTFS) Sort() {
	sort.Slice(g.Agencies, func(i, j int) bool {
		return g.Agencies[i].ID < g.Agencies[j].ID
	})
	sort.Slice(g.CalendarDates, func(i, j int) bool {
		iv := g.CalendarDates[i]
		jv := g.CalendarDates[j]
		iStr := iv.Date + ":" + iv.ServiceID + ":" + strconv.Itoa(iv.ExceptionType)
		jStr := jv.Date + ":" + jv.ServiceID + ":" + strconv.Itoa(jv.ExceptionType)
		return iStr < jStr
	})
	sort.Slice(g.Calendars, func(i, j int) bool {
		return g.Calendars[i].ServiceID < g.Calendars[j].ServiceID
	})
	sort.Slice(g.Routes, func(i, j int) bool {
		return g.Routes[i].ID < g.Routes[j].ID
	})
	sort.Slice(g.Stops, func(i, j int) bool {
		return g.Stops[i].ID < g.Stops[j].ID
	})
	sort.Slice(g.StopsTimes, func(i, j int) bool {
		iv := g.StopsTimes[i]
		jv := g.StopsTimes[j]
		iStr := iv.StopID + ":" + iv.TripID + ":" + iv.Departure
		jStr := jv.StopID + ":" + jv.TripID + ":" + jv.Departure
		return iStr < jStr
	})
	sort.Slice(g.Transfers, func(i, j int) bool {
		iv := g.Transfers[i]
		jv := g.Transfers[j]
		iStr := iv.FromStopID + iv.ToStopID + strconv.Itoa(iv.Type)
		jStr := jv.FromStopID + jv.ToStopID + strconv.Itoa(jv.Type)
		return iStr < jStr
	})
	sort.Slice(g.Trips, func(i, j int) bool {
		return g.Trips[i].ID < g.Trips[j].ID
	})
}
