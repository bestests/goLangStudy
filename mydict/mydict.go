package mydict

import "errors"

var errNotFound = errors.New("ERROR - Not Found")

// Dictionary type
/*
	struct 외로 이렇게 type 선언 가능
*/
type Dictionary map[string]string

// Search dictionary search
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}
