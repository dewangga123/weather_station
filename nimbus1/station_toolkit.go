package nimbus1

import (
	"weather_station/api"
)

func NewNimbus1StationToolkit() api.StationToolkit {
	return &nimbus1StationToolkit{}
}

type nimbus1StationToolkit struct {
}

func (toolkit *nimbus1StationToolkit) MakeTemperature() api.TemperatureSensorDevice {
	return &nimbus1TemperatureSensor{}
}

func (toolkit *nimbus1StationToolkit) GetAlarmClock() api.AlarmClockDevice {
	return &nimbus1AlarmClock{}
}

func (toolkit *nimbus1StationToolkit) GetPersistent() api.PersistentDevice {
	return &nimbus1PersistentDevice{
		DB: make(map[string][]byte),
	}
}
