package weather

import "weather_station/api"

func NewWeatherStation(name string, st api.StationToolkit) WeatherStation {
	return WeatherStation{
		Name:       name,
		TempSensor: NewTemperatureSensor(st),
	}
}

type WeatherStation struct {
	Name       string
	TempSensor TemperatureSensor
}

func (station *WeatherStation) AddTempObserver(observer Observer) {
	station.TempSensor.AddObserver(observer)
}

func (station *WeatherStation) Read() {
	station.TempSensor.Read()
}
