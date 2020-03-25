package wmsdata_impl

import "time"

type HiLoDataImpl struct {
	itsLowValue  float32
	itsHighValue float32
	itsHighTime  time.Time
	itsLowTime   time.Time
}

func (data *HiLoDataImpl) InitData(init float32, initTime time.Time) {
	data.itsHighValue = init
	data.itsLowValue = init
	data.itsLowTime = initTime
	data.itsHighTime = initTime
}

func (data *HiLoDataImpl) CurrentReading(current float32, readingTime time.Time) bool {
	if current > data.itsHighValue {
		data.itsHighValue = current
		data.itsHighTime = readingTime
		return true
	} else if current < data.itsLowValue {
		data.itsLowValue = current
		data.itsLowTime = readingTime
		return true
	}
	return false

}
func (data *HiLoDataImpl) NewDay(initial float32, dayTime time.Time) {
	data.itsHighValue = initial
	data.itsLowValue = initial
	data.itsHighTime = dayTime
	data.itsLowTime = dayTime
}

func (data *HiLoDataImpl) GetHighValue() float32 {
	return data.itsHighValue
}
func (data *HiLoDataImpl) GetHighTime() time.Time {
	return data.itsHighTime
}
func (data *HiLoDataImpl) GetLowValue() float32 {
	return data.itsLowValue
}
func (data *HiLoDataImpl) GetLowTime() time.Time {
	return data.itsLowTime
}
