package weather

import (
	"weather_station/api"
	"weather_station/component"
)

func NewWeatherStation(name string, st api.StationToolkit) component.WeatherStationComponent {
	return &WeatherStation{
		Name:       name,
		TempSensor: NewTemperatureSensor(st),
	}
}

type WeatherStation struct {
	Name       string
	TempSensor TemperatureSensor
}

func (station *WeatherStation) AddTempObserver(observer component.Observer) {
	station.TempSensor.AddObserver(observer)
}

func (station *WeatherStation) Read() {
	station.TempSensor.Read()
}
