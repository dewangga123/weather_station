package api

type PersistentDevice interface {
	Store(name string, data interface{}) error
	Retrive(name string) ([]byte, error)
}
