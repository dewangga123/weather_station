package persistent

import (
	"time"
	"weather_station/api"
	"weather_station/wmsdata"
)

func NewDataToolkitImpl() wmsdata.DataToolkit {
	return &dataToolkitImpl{}
}

type dataToolkitImpl struct{}

func (toolkit *dataToolkitImpl) GetTempHiLoData(st api.StationToolkit, dataType string, theDate time.Time, init float32, initTime time.Time) wmsdata.HiLoData {
	data := &hiLoDataProxy{}
	data.InitData(st, dataType, theDate, init, initTime)
	return data
}
