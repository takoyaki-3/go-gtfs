package gtfs

type EdgeTimetable struct {
	Edges      []TimetableEdge
	Properties []TimetableEdgeProperty
	Stops      []Stop
}

type TimetableEdge struct {
	TripId        string `csv:"trip_id"`
	FromStop      string `csv:"from_stop_id"`
	ToStop        string `csv:"to_stop_id"`
	DepartureTime string `csv:"departure_time"`
	ArrivalTime   string `csv:"arrival_time"`
	PickupType    int    `csv:"pickup_type"`
	DropOffType   int    `csv:"drop_off_type"`
	StopHeadSign  string `csv:"stop_head_sign"`
}

type TimetableEdgeProperty struct {
	TripID      string `csv:"trip_id"`
	Name        string `csv:"trip_short_name"`
	RouteID     string `csv:"route_id"`
	ServiceID   string `csv:"service_id"`
	ShapeID     string `csv:"shape_id"`
	DirectionID string `csv:"direction_id"`
	Headsign    string `csv:"trip_headsign"`
	AgencyID    string `csv:"agency_id"`
	ShortName   string `csv:"route_short_name"`
	LongName    string `csv:"route_long_name"`
	Type        int    `csv:"route_type"`
	Desc        string `csv:"route_url"`
	URL         string `csv:"route_desc"`
	Color       string `csv:"route_color"`
	TextColor   string `csv:"route_text_color"`
}
