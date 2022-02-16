package tool

import (
	"sync"

	"github.com/takoyaki-3/go-gtfs"
	"github.com/takoyaki-3/goraph"
	"github.com/takoyaki-3/goraph/geometry"
	"github.com/takoyaki-3/goraph/geometry/h3"
	"github.com/takoyaki-3/goraph/search"
)

// connectRange: 接続する停留所間の最大距離
// walkingSpeed: 歩行速度（分速メートル）
// road:				 地図データのグラフ
// numThread:    使用するスレッド数
func MakeTransfer(g *gtfs.GTFS, connectRange float64, walkingSpeed float64, road goraph.Graph, numThread int) {

	// 地図データ読み込み
	h3index := h3.MakeH3Index(road, 9)

	wg := sync.WaitGroup{}
	wg.Add(numThread)
	type Dis struct {
		fromId string
		toId   string
		dis    float64
	}
	diss := make([][]Dis, numThread)
	for rank := 0; rank < numThread; rank++ {
		go func(rank int) {
			defer wg.Done()
			for i, stopI := range g.Stops {
				for j, stopJ := range g.Stops {
					if (i+j)%numThread != rank {
						continue
					}
					if i <= j {
						continue
					}
					dis := geometry.HubenyDistance(goraph.LatLon{
						Lat: stopI.Latitude,
						Lon: stopI.Longitude,
					}, goraph.LatLon{
						Lat: stopJ.Latitude,
						Lon: stopJ.Longitude,
					})
					if dis <= connectRange && len(road.LatLons) != 0 {
						// 道のりも計算
						route := search.Search(road, search.Query{
							From: h3.Find(road, h3index, goraph.LatLon{
								Lat: stopI.Latitude,
								Lon: stopI.Longitude,
							}, 9),
							To: h3.Find(road, h3index, goraph.LatLon{
								Lat: stopJ.Latitude,
								Lon: stopJ.Longitude,
							}, 9),
						})
						dis = route.Cost
					}
					if dis <= connectRange || stopI.Parent == stopJ.Parent {
						diss[rank] = append(diss[rank], Dis{
							fromId: stopI.ID,
							toId:   stopJ.ID,
							dis:    dis})
					}
				}
			}
		}(rank)
	}
	wg.Wait()
	for _, arr := range diss {
		for _, v := range arr {
			g.Transfers = append(g.Transfers, gtfs.Transfer{
				FromStopID: v.fromId,
				ToStopID: v.toId,
				MinTime: int(v.dis * 60.0 / walkingSpeed),
				Type: 0,
			})
			g.Transfers = append(g.Transfers, gtfs.Transfer{
				FromStopID: v.toId,
				ToStopID: v.fromId,
				MinTime: int(v.dis * 60.0 / walkingSpeed),
				Type: 0,
			})
		}
	}
}
