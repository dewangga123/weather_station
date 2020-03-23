package api

type AlarmListener interface {
	Tic(seconds int)
}
