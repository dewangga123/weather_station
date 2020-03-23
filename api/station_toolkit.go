package api

type StationToolkit interface {
	MakeTemperature() TemperatureSensorImpl
	GetAlarmClock() AlarmClockImpl
}
