package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("ERR: value not exist.")
	ErrWordExists = errors.New("ERR: cannot add word because it already exists.")
)

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err == nil {
		return ErrWordExists
	}
	d[key] = value
	return nil
}
