package wmsdata_impl

import (
	"encoding/json"
	"time"
	"weather_station/api"
	"weather_station/wmsdata"
)

func NewHiLoDataImpl(st api.StationToolkit,
	dataType string,
	theDate time.Time,
	init float32,
	initTime time.Time) wmsdata.HiLoData {
	data := &hiLoDataImpl{}
	data.InitData(st, dataType, theDate, init, initTime)
	return data
}

type hiLoDataImpl struct {
	itsLowValue   float32
	itsHighValue  float32
	itsHighTime   time.Time
	itsLowTime    time.Time
	itsType       string
	itsPI         api.PersistentImpl
	itsStorageKey string
}

func (data *hiLoDataImpl) InitData(st api.StationToolkit,
	dataType string,
	theDate time.Time,
	init float32,
	initTime time.Time) {
	data.itsPI = st.GetPersistent()
	data.itsType = dataType
	data.itsStorageKey = data.calculateStorageKey(theDate)
	dataBytes, err := data.itsPI.Retrive(data.itsStorageKey)
	if err == nil {
		json.Unmarshal(dataBytes, data)
		data.CurrentReading(init, initTime)
	} else {
		data.itsHighValue = init
		data.itsLowValue = init
		data.itsHighTime = initTime
		data.itsLowTime = initTime
	}
}

func (data *hiLoDataImpl) calculateStorageKey(inputTime time.Time) string {
	return data.itsType + inputTime.Format("01022006")
}

func (data *hiLoDataImpl) store() error {
	return data.itsPI.Store(data.itsStorageKey, data.itsStorageKey)
}
func (data *hiLoDataImpl) CurrentReading(current float32, readingTime time.Time) {
	if current > data.itsHighValue {
		data.itsHighValue = current
		data.itsHighTime = readingTime
		data.store()
	} else if current < data.itsLowValue {
		data.itsLowValue = current
		data.itsLowTime = readingTime
		data.store()
	}

}
func (data *hiLoDataImpl) NewDay(initial float32, dayTime time.Time) {
	data.store()
	data.itsHighValue = initial
	data.itsLowValue = initial
	data.itsHighTime = dayTime
	data.itsLowTime = dayTime
	data.itsStorageKey = data.calculateStorageKey(dayTime)
	data.store()
}
func (data *hiLoDataImpl) GetHighValue() float32 {
	return data.itsHighValue
}
func (data *hiLoDataImpl) GetHighTime() time.Time {
	return data.itsHighTime
}
func (data *hiLoDataImpl) GetLowValue() float32 {
	return data.itsLowValue
}
func (data *hiLoDataImpl) GetLowTime() time.Time {
	return data.itsLowTime
}
