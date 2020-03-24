package ui

import (
	"fmt"
	"weather_station/component"
)

func NewTemperatureUIObserver() component.Observer {
	return TemperatureUIObserver{}
}

type TemperatureUIObserver struct{}

func (observer TemperatureUIObserver) Update(data float32) {
	fmt.Println("Temperature Value", data)
}
