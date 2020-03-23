package nimbus1

import (
	"time"
	"weather_station/api"
)

type nimbus1AlarmClock struct {
}

func (alarm *nimbus1AlarmClock) Register(alarmListener api.AlarmListener) {
	for {
		time.Sleep(time.Second)
		alarmListener.Tic(1)
	}
}
