package mydict

import "errors"

/*
var errNotFound = errors.New("ERROR - Not Found")
var errWordExists = errors.New("ERROR - That Word is Exists")
var errCantUpdate = errors.New("ERROR - Cant Update Non-exists word")
*/

var (
	errNotFound   = errors.New("ERROR - Not Found")
	errWordExists = errors.New("ERROR - That Word is Exists")
	errCantUpdate = errors.New("ERROR - Cant Update Non-exists word")
	errCantDelete = errors.New("ERROR - Cant Delete Non-exists word")
)

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

// Add x Dictionary - add words
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExists
		}

		return nil
	*/
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

// Update x Dictionary - Update word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = def
	}

	return nil
}

// Delete x Dictionary - Delete word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errCantDelete
	case nil:
		delete(d, word)
	}

	return nil
}
