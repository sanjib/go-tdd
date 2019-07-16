package maps

const (
	ErrorWordNotFound           = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists             = DictionaryErr("cannot add word; already exists")
	UpdateErrorWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if definition, ok := d[word]; ok {
		return definition, nil
	}
	return "", ErrorWordNotFound
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	// ErrorWordNotFound is a good thing, means we can
	// add the word
	case ErrorWordNotFound:
		d[word] = definition
	// nil means Search successful word was found,
	// so we actually shouldn't try to add it as
	// it already exists
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

// Update is opposite of Add. In Add we add only if word doesn't
// exist. In Update we need to update only if word exists.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	// ErrorWordNotFound is a bad thing, means we are trying
	// to update non-existing word
	case ErrorWordNotFound:
		return UpdateErrorWordDoesNotExist
	// nil means a good thing, that word exists and we should
	// be fine to update it
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
