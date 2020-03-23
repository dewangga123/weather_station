package weather

import "weather_station/api"

func NewTemperatureSensor(st api.StationToolkit) TemperatureSensor {
	return TemperatureSensor{
		TemperatureSensorImpl: st.MakeTemperature(),
		AlarmClock:            newAlarmClock(st),
	}
}

type TemperatureSensor struct {
	TemperatureSensorImpl api.TemperatureSensorImpl
	AlarmClock            alarmClock
	itsLastReading        float32
	observers             []Observer
}

func (sensor *TemperatureSensor) Read() {
	sensor.AlarmClock.Wakeup(5, func() {
		sensor.check()
	})
}

func (sensor *TemperatureSensor) check() {
	val := sensor.TemperatureSensorImpl.Read()
	if val != sensor.itsLastReading {
		sensor.itsLastReading = val
		sensor.NotifyObserver(val)
	}
}

func (sensor *TemperatureSensor) AddObserver(observer Observer) {
	if sensor.observers == nil {
		sensor.observers = []Observer{}
	}
	sensor.observers = append(sensor.observers, observer)
}

func (sensor *TemperatureSensor) NotifyObserver(data float32) {
	if sensor.observers != nil {
		for _, observer := range sensor.observers {
			observer.Update(data)
		}
	}
}
