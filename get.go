package gtfs

func (g *GTFS)GetRoute(routeId string)Route{
	for _,r := range g.Routes{
		if r.ID == routeId {
			return r
		}
	}
	return Route{}
}

func (g *GTFS)GetTrip(tripId string)Trip{
	for _,t:=range g.Trips{
		if t.ID == tripId {
			return t
		}
	}
	return Trip{}
}
