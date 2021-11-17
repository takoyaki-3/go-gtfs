package edgetimetable

import (
	"fmt"
	"path"

	. "github.com/takoyaki-3/go-gtfs"
	csvtag "github.com/artonge/go-csv-tag/v2"
)

func Dump(edgeTimetable *EdgeTimetable, dirPath string, filter map[string]bool) error {

	files := map[string]interface{}{
		"ptedges.csv":        edgeTimetable.Edges,
		"pteproperties.csv":  edgeTimetable.Properties,
		"stops.csv":					edgeTimetable.Stops,
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
