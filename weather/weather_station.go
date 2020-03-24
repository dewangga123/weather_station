package weather

import (
	"weather_station/api"
	"weather_station/component"
)

func NewWeatherStation(name string, st api.StationToolkit) component.WeatherStationComponent {
	return &WeatherStation{
		Name:           name,
		TempSensor:     newTemperatureSensor(st),
		HiloTempSensor: newHiloTemperatureSensor(st),
	}
}

type WeatherStation struct {
	Name           string
	TempSensor     temperatureSensor
	HiloTempSensor temperatureHiLo
}

func (station *WeatherStation) AddTempObserver(observer component.Observer) {
	station.TempSensor.AddObserver(observer)
}

func (station *WeatherStation) AddHiLoTempObserver(observer component.HiloObserver) {
	station.HiloTempSensor.AddObserver(observer)
}

func (station *WeatherStation) Read() {
	//go station.TempSensor.Read()
	station.HiloTempSensor.Read()
}
