package weather

import "weather_station/api"

func newAlarmClock(st api.StationToolkit) alarmClock {
	return alarmClock{
		AlarmClockImpl: st.GetAlarmClock(),
	}
}

type alarmClock struct {
	interval       int
	cacheInterval  int
	listener       func()
	AlarmClockImpl api.AlarmClockImpl
}

func (alarm *alarmClock) Tic(milliseconds int) {
	alarm.cacheInterval += milliseconds
	if alarm.cacheInterval >= alarm.interval {
		alarm.cacheInterval = 0
		if alarm.listener != nil {
			alarm.listener()
		}
	}
}

func (alarm *alarmClock) Wakeup(interval int, listener func()) {
	alarm.interval = interval
	alarm.listener = listener
	alarm.AlarmClockImpl.Register(alarm)
}
