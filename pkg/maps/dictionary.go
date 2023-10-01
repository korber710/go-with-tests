package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}

	return value, nil
}

var ErrWordNotFound = errors.New("cannot withdraw, insufficient funds")
