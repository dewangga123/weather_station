package api

type AlarmClockDevice interface {
	Register(alarmListener AlarmListener)
}
