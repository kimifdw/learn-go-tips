package dictionaries

// Dictionary ：map为引用类型
type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(key string) (string, error) {
	definition, ok := dictionary[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (dictionary Dictionary) Add(key, value string) error {

	_, err := dictionary.Search(key)
	switch err {
	case ErrNotFound:
		dictionary[key] = value
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
		return ErrWordDoesNotExists
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
