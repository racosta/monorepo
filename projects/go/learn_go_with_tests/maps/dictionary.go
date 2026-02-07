// Package dictionary provides a simple in-memory dictionary with basic operations.
package dictionary

const (
	// ErrNotFound means the definition could not be found for the given word
	ErrNotFound = Err("could not find the word you were looking for")

	// ErrWordExists means you are trying to add a word that is already known
	ErrWordExists = Err("cannot add word because it already exists")

	// ErrWordDoesNotExist occurs when trying to perform an operation on a word not in the dictionary
	ErrWordDoesNotExist = Err("cannot perform operation on word because it does not exist")
)

// Err are errors that can happen when interacting with the dictionary.
type Err string

// Error returns the error message.
func (e Err) Error() string {
	return string(e)
}

// Dictionary store definitions to words.
type Dictionary map[string]string

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add inserts a word and definition into the dictionary.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update changes the definition of a given word.
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

// Delete removes a word from the dictionary.
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
