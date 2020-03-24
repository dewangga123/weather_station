package nimbus1

import (
	"math/rand"
	"time"
)

type nimbus1TemperatureSensor struct {
}

func (sensor *nimbus1TemperatureSensor) Read() float32 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return 25 * r.Float32()
}
   