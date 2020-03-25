package weather

import (
	"weather_station/api"
	"weather_station/component"
)

func newTemperatureSensor(st api.StationToolkit) temperatureSensor {
	return temperatureSensor{
		TemperatureSensorImpl: st.MakeTemperature(),
		AlarmClock:            newAlarmClock(st),
	}
}

type temperatureSensor struct {
	TemperatureSensorImpl api.TemperatureSensorDevice
	AlarmClock            alarmClock
	itsLastReading        float32
	observers             []component.Observer
}

func (sensor *temperatureSensor) Read() {
	sensor.AlarmClock.Wakeup(5, func() {
		sensor.check()
	})
}

func (sensor *temperatureSensor) check() {
	val := sensor.TemperatureSensorImpl.Read()
	if val != sensor.itsLastReading {
		sensor.itsLastReading = val
		sensor.NotifyObserver(val)
	}
}

func (sensor *temperatureSensor) AddObserver(observer component.Observer) {
	if sensor.observers == nil {
		sensor.observers = []component.Observer{}
	}
	sensor.observers = append(sensor.observers, observer)
}

func (sensor *temperatureSensor) NotifyObserver(data float32) {
	if sensor.observers != nil {
		for _, observer := range sensor.observers {
			observer.Update(data)
		}
	}
}
