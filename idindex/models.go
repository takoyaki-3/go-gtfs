package idindex

import (
	. "github.com/takoyaki-3/go-gtfs"
)

type Id2Index map[string]int

type GTFSId2Index struct {
	Agencies      Id2Index
	Routes        Id2Index
	Stops         Id2Index
	Trips         Id2Index
}

func MakeId2Index(g *GTFS)(*GTFSId2Index){
	d := &GTFSId2Index{
		Agencies: newId2Index(),
		Routes: newId2Index(),
		Stops: newId2Index(),
		Trips: newId2Index(),
	}
	for i,v:=range g.Agencies{
		d.Agencies[v.ID] = i
	}
	for i,v:=range g.Routes{
		d.Routes[v.ID] = i
	}
	for i,v:=range g.Stops{
		d.Stops[v.ID] = i
	}
	for i,v:=range g.Trips{
		d.Trips[v.ID] = i
	}
	return d
}

func newId2Index()Id2Index{
	return Id2Index{}
}
