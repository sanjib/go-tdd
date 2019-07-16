package maps

const (
	ErrorWordNotFound = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists   = DictionaryErr("cannot add word; already exists")
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
