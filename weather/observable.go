package weather

type Observable interface {
	AddObserver(observer Observer)
	NotifyObserver(data float32)
}
