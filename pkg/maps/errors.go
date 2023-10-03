package maps

const (
	ErrWordNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrWordAlreadyExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
