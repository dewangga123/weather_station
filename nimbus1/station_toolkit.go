package nimbus1

import "weather_station/api"

func NewNimbus1StationToolkit() api.StationToolkit {
	return &nimbus1StationToolkit{}
}

type nimbus1StationToolkit struct {
}

func (toolkit *nimbus1StationToolkit) MakeTemperature() api.TemperatureSensorImpl {
	return &nimbus1TemperatureSensor{}
}

func (toolkit *nimbus1StationToolkit) GetAlarmClock() api.AlarmClockImpl {
	return &nimbus1AlarmClock{}
}
