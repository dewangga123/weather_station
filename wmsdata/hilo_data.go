package wmsdata

import "time"

type HiLoData interface {
	CurrentReading(value float32, readingTime time.Time)
	NewDay(initial float32, dayTime time.Time)
	GetHighValue() float32
	GetHighTime() time.Time
	GetLowValue() float32
	GetLowTime() time.Time
}
