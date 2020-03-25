package weather

import (
	"time"
	"weather_station/api"
	"weather_station/component"
	"weather_station/wmsdata"
)

func newHiloTemperatureSensor(st api.StationToolkit, dt wmsdata.DataToolkit) temperatureHiLo {
	return temperatureHiLo{
		TemperatureSensorImpl: st.MakeTemperature(),
		AlarmClock:            newAlarmClock(st),
		ItsLastReading:        dt.GetTempHiLoData(st, "temp", time.Now(), 14.0, time.Now()),
	}
}

type temperatureHiLo struct {
	TemperatureSensorImpl api.TemperatureSensorDevice
	AlarmClock            alarmClock
	observers             []component.HiloObserver
	ItsLastReading        wmsdata.HiLoData
}

func (sensor *temperatureHiLo) Read() {
	sensor.AlarmClock.Wakeup(5, func() {
		sensor.check()
	})
}

func (sensor *temperatureHiLo) check() {
	val := sensor.TemperatureSensorImpl.Read()
	if sensor.ItsLastReading != nil {
		sensor.ItsLastReading.CurrentReading(val, time.Now())
		sensor.NotifyObserver(sensor.ItsLastReading)
	}
}

func (sensor *temperatureHiLo) AddObserver(observer component.HiloObserver) {
	if sensor.observers == nil {
		sensor.observers = []component.HiloObserver{}
	}
	sensor.observers = append(sensor.observers, observer)
}

func (sensor *temperatureHiLo) NotifyObserver(data wmsdata.HiLoData) {
	if sensor.observers != nil {
		for _, observer := range sensor.observers {
			observer.Update(data)
		}
	}
}
