package api

type PersistentImpl interface {
	Store(name string, data interface{}) error
	Retrive(name string) ([]byte, error)
}
