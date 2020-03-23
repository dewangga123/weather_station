package component

type WeatherStationComponent interface {
	AddTempObserver(observer Observer)
	Read()
}
