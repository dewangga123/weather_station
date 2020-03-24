package component

type WeatherStationComponent interface {
	AddTempObserver(observer Observer)
	AddHiLoTempObserver(observer HiloObserver)
	Read()
}
