package weather


type Observer interface {
	Update(data float32)
}