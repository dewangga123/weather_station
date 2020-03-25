package api

type TemperatureSensorDevice interface {
	Read() float32
}
