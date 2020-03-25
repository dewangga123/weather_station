package main

import (
	"weather_station/nimbus1"
	"weather_station/persistent"
	"weather_station/ui"
	"weather_station/weather"
)

func main() {
	st := nimbus1.NewNimbus1StationToolkit()
	dt := persistent.NewDataToolkitImpl()
	weatherStation := weather.NewWeatherStation("Station", st, dt)
	tempObserver := ui.NewTemperatureUIObserver()
	tempHiLoObserver := ui.NewTempHiloUIObserver()
	weatherStation.AddTempObserver(tempObserver)
	weatherStation.AddHiLoTempObserver(tempHiLoObserver)
	weatherStation.Read()
}
