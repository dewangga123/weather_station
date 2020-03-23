package ui

import (
	"weather_station/component"
)

func NewTemperatureUIObserver() component.Observer {
	return TemperatureUIObserver{
		Screen: monitoringScreen{},
	}
}

type TemperatureUIObserver struct {
	Screen monitoringScreen
}

func (observer TemperatureUIObserver) Update(data float32) {
	observer.Screen.DisplayTemperature(data)
}
