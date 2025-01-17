package gtfs

import (
	"time"
)

// 対象日のGTFSのみに絞り込む
func (g *GTFS) ExtractByDate(date time.Time) *GTFS {

	theDayServices := map[string]bool{}

	for _, calendar := range g.Calendars {
		// 曜日ごとの運行・運休を記録
		var isService bool
		switch date.Weekday() {
		case time.Sunday:
			isService = (calendar.Sunday == 1)
		case time.Monday:
			isService = (calendar.Monday == 1)
		case time.Tuesday:
			isService = (calendar.Tuesday == 1)
		case time.Wednesday:
			isService = (calendar.Wednesday == 1)
		case time.Thursday:
			isService = (calendar.Thursday == 1)
		case time.Friday:
			isService = (calendar.Friday == 1)
		case time.Saturday:
			isService = (calendar.Saturday == 1)
		}
		theDayServices[calendar.ServiceID] = isService
	}

	// 臨時運行・運休情報を追加
	dateStr := date.Format("20060102")
	for _, calendarDate := range g.CalendarDates {
		if calendarDate.Date == dateStr {
			theDayServices[calendarDate.ServiceID] = (calendarDate.ExceptionType == 1)
		}
	}

	newG := GTFS{}
	newG = *g
	newG.Trips = []Trip{}
	newG.CalendarDates = []CalendarDate{}
	newG.Calendars = []Calendar{}
	newG.StopsTimes = []StopTime{}

	trips := map[string]bool{}
	for _, trip := range g.Trips {
		if v, ok := theDayServices[trip.ServiceID]; ok {
			if v {
				newG.Trips = append(newG.Trips, trip)
				trips[trip.ID] = true
			}
		}
	}
	for _, stopTime := range g.StopsTimes {
		if _, ok := trips[stopTime.TripID]; ok {
			newG.StopsTimes = append(newG.StopsTimes, stopTime)
		}
	}
	for _, calendar := range g.Calendars {
		if v, ok := theDayServices[calendar.ServiceID]; ok {
			if v {
				newG.Calendars = append(newG.Calendars, calendar)
			}
		}
	}
	for _, calendarDate := range g.CalendarDates {
		if v, ok := theDayServices[calendarDate.ServiceID]; ok {
			if v {
				newG.CalendarDates = append(newG.CalendarDates, calendarDate)
			}
		}
	}
	return &newG
}

func ArrayIn(target string, array []string) bool {
	for _,v:=range array {
		if v == target {
			return true
		}
	}
	return false
}

// routeIDを基に絞り込む
func (g *GTFS) ExtractByRouteIDs(ids []string) *GTFS {

	newG := GTFS{}
	newG = *g
	newG.Trips = []Trip{}
	newG.StopsTimes = []StopTime{}

	trips := map[string]bool{}
	for _, trip := range g.Trips {
		if ArrayIn(trip.RouteID, ids) {
			newG.Trips = append(newG.Trips, trip)
			trips[trip.ID] = true
		}
	}
	for _, stopTime := range g.StopsTimes {
		if _, ok := trips[stopTime.TripID]; ok {
			newG.StopsTimes = append(newG.StopsTimes, stopTime)
		}
	}
	return &newG
}

// agencyIDを基に絞り込む
func (g *GTFS) ExtractByAgencyID(ids []string) *GTFS {

	newG := GTFS{}
	newG = *g
	newG.Trips = []Trip{}
	newG.StopsTimes = []StopTime{}

	routes := map[string]bool{}
	for _, route := range g.Routes {
		if ArrayIn(route.AgencyID, ids) {
			routes[route.ID] = true
		}
	}

	trips := map[string]bool{}
	for _, trip := range g.Trips {
		if _,ok:=routes[trip.RouteID];ok {
			newG.Trips = append(newG.Trips, trip)
			trips[trip.ID] = true
		}
	}
	for _, stopTime := range g.StopsTimes {
		if _, ok := trips[stopTime.TripID]; ok {
			newG.StopsTimes = append(newG.StopsTimes, stopTime)
		}
	}

	return &newG
}
