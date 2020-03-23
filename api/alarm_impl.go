package api

type AlarmClockImpl interface {
	Register(alarmListener AlarmListener)
}
