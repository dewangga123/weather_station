package component

type Observable interface {
	AddObserver(observer Observer)
	NotifyObserver(data float32)
}
