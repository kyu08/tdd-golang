package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("ERR: value not exist.")
	ErrWordExists       = DictionaryErr("ERR: cannot add word because it already exists.")
	ErrWordDoesNotExist = DictionaryErr("ERR: cannot update word because it does not exist.")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
