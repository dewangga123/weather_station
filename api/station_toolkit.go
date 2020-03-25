package api

type StationToolkit interface {
	MakeTemperature() TemperatureSensorDevice
	GetAlarmClock() AlarmClockDevice
	GetPersistent() PersistentDevice
}
