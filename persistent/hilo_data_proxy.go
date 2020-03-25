package persistent

import (
	"encoding/json"
	"time"
	"weather_station/api"
	"weather_station/wmsdata_impl"
)

type hiLoDataProxy struct {
	itsType       string
	itsPI         api.PersistentDevice
	itsStorageKey string
	itsData       *wmsdata_impl.HiLoDataImpl
}

func (data *hiLoDataProxy) InitData(st api.StationToolkit, dataType string,
	theDate time.Time,
	init float32,
	initTime time.Time) {
	data.itsPI = st.GetPersistent()
	data.itsType = dataType
	data.itsStorageKey = data.calculateStorageKey(theDate)
	data.itsData = &wmsdata_impl.HiLoDataImpl{}
	dataBytes, err := data.itsPI.Retrive(data.itsStorageKey)
	if err != nil {
		data.itsData.InitData(init, initTime)
	} else {
		json.Unmarshal(dataBytes, data.itsData)
	}
}

func (data *hiLoDataProxy) CurrentReading(current float32, readingTime time.Time) bool {
	change := data.itsData.CurrentReading(current, readingTime)
	if change {
		data.store()
	}
	return change
}
func (data *hiLoDataProxy) NewDay(initial float32, dayTime time.Time) {
	data.store()
	data.itsData.NewDay(initial, dayTime)
	data.itsStorageKey = data.calculateStorageKey(dayTime)
	data.store()
}

func (data *hiLoDataProxy) calculateStorageKey(inputTime time.Time) string {
	return data.itsType + inputTime.Format("01022006")
}

func (data *hiLoDataProxy) store() error {
	return data.itsPI.Store(data.itsStorageKey, data.itsStorageKey)
}

func (data *hiLoDataProxy) GetHighValue() float32 {
	return data.itsData.GetHighValue()
}
func (data *hiLoDataProxy) GetHighTime() time.Time {
	return data.itsData.GetHighTime()
}
func (data *hiLoDataProxy) GetLowValue() float32 {
	return data.itsData.GetLowValue()
}
func (data *hiLoDataProxy) GetLowTime() time.Time {
	return data.itsData.GetLowTime()
}
