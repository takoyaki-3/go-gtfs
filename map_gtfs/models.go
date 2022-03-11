package mapgtfs

import (
	. "github.com/takoyaki-3/go-gtfs"
)

type MapGTFS struct {
	Path          string // The path to the containing directory
	Agency        Agency
	Agencies      map[string]Agency
	Routes        map[string][]Route
	Stops         map[string][]Stop
	StopsTimes    []StopTime
	Trips         map[string][]Trip
	Calendars     []Calendar
	CalendarDates []CalendarDate
	Transfers     []Transfer
}
