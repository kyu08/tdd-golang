package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	var result string = d[key]
	var err error
	if result == "" {
		err = errors.New("ERR: value not exist.")
	}

	return result, err
}
