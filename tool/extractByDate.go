package tool

import (
	"time"
	"github.com/takoyaki-3/go-gtfs"
)

func ExtractByDate(g *gtfs.GTFS, date time.Time)*gtfs.GTFS{

	theDayServices := map[string]bool{}

	for _,calendar := range g.Calendars {
		// 曜日ごとの運行・運休を記録
		var isService bool
		switch date.Weekday() {
		case time.Sunday:
			isService = (calendar.Sunday==1)
		case time.Monday:
			isService = (calendar.Monday==1)
		case time.Tuesday:
			isService = (calendar.Tuesday==1)
		case time.Wednesday:
			isService = (calendar.Wednesday==1)
		case time.Thursday:
			isService = (calendar.Thursday==1)
		case time.Friday:
			isService = (calendar.Friday==1)
		case time.Saturday:
			isService = (calendar.Saturday==1)
		}
		theDayServices[calendar.ServiceID] = isService
	}

	// 臨時運行・運休情報を追加
	dateStr := date.Format("20060102")
	for _,calendarDate := range g.CalendarDates{
		if calendarDate.Date == dateStr {
			theDayServices[calendarDate.ServiceID] = (calendarDate.ExceptionType == 1)
		}
	}

	newG := gtfs.GTFS{}
	newG = *g
	newG.Trips = []gtfs.Trip{}
	newG.CalendarDates = []gtfs.CalendarDate{}
	newG.Calendars = []gtfs.Calendar{}
	newG.StopsTimes = []gtfs.StopTime{}

	trips := map[string]bool{}
	for _,trip := range g.Trips {
		if v,ok:=theDayServices[trip.ServiceID];ok{
			if v{
				newG.Trips = append(newG.Trips, trip)
				trips[trip.ID] = true
			}
		}
	}
	for _,stopTime := range g.StopsTimes {
		if _,ok:=trips[stopTime.TripID];ok {
			newG.StopsTimes = append(newG.StopsTimes, stopTime)
		}
	}
	for _,calendar := range g.Calendars {
		if v,ok:=theDayServices[calendar.ServiceID];ok{
			if v{
				newG.Calendars = append(newG.Calendars, calendar)
			}
		}
	}
	for _,calendarDate := range g.CalendarDates {
		if v,ok := theDayServices[calendarDate.ServiceID];ok{
			if v{
				newG.CalendarDates = append(newG.CalendarDates, calendarDate)
			}
		}
	}
	return &newG
}
