package nimbus1

import (
	"encoding/json"
	"errors"
)

type nimbus1PersistentDevice struct {
	DB map[string][]byte
}

func (persistent *nimbus1PersistentDevice) Store(name string, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	persistent.DB[name] = dataBytes
	return nil
}
func (persistent *nimbus1PersistentDevice) Retrive(name string) ([]byte, error) {
	if data, ok := persistent.DB[name]; ok {
		return data, nil
	} else {
		return data, errors.New("not found")
	}
}
