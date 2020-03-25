package wmsdata

import (
	"time"
	"weather_station/api"
)

type DataToolkit interface {
	GetTempHiLoData(st api.StationToolkit, dataType string, theDate time.Time, init float32, initTime time.Time) HiLoData
}
