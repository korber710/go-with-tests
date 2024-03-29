package maps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchReturnsExpectedStringWithKnownKey(t *testing.T) {
	t.Parallel()
	for _, tc := range []dictionaryTestCase{
		newDictionaryTestCase("test", "this is just a test"),
		newDictionaryTestCase("hello", "world"),
	} {
		tc := tc
		t.Run(fmt.Sprintf("looking for value '%s' from key '%s'", tc.value, tc.key), func(t *testing.T) {
			t.Parallel()
			got, _ := tc.d.Search(tc.key)
			want := tc.value
			assert.Equal(t, got, want)
		})
	}
}

func TestSearchReportsErrorWithUnknownKey(t *testing.T) {
	t.Parallel()
	for _, tc := range []dictionaryTestCase{
		newDictionaryTestCase("test", "this is just a test"),
		newDictionaryTestCase("hello", "world"),
	} {
		tc := tc
		t.Run(fmt.Sprintf("expecting error when looking for fake key on %v", tc.d), func(t *testing.T) {
			t.Parallel()
			fakeKey := fmt.Sprintf("fake-%s", tc.key)
			_, err := tc.d.Search(fakeKey)
			want := ErrWordNotFound
			assert.Error(t, err)
			assert.ErrorIs(t, err, want)
		})
	}
}

func TestAddingNewWord(t *testing.T) {
	t.Parallel()
	dictionary := Dictionary{"test": "this is just a test"}
	_ = dictionary.Add("hello", "world")
	got, _ := dictionary.Search("hello")
	want := "world"
	assert.Equal(t, got, want)
}

func TestAddingExistingWordReportsError(t *testing.T) {
	t.Parallel()
	dictionary := Dictionary{"test": "this is just a test"}
	err := dictionary.Add("test", "new definition")
	want := ErrWordAlreadyExists
	assert.Error(t, err)
	assert.ErrorIs(t, err, want)
}

func TestDeleteExistingWord(t *testing.T) {
	t.Parallel()
	word := "test"
	dictionary := Dictionary{word: "this is just a test"}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assert.Error(t, err)
}

type dictionaryTestCase struct {
	key   string
	value string
	d     Dictionary
}

func newDictionaryTestCase(key, value string) dictionaryTestCase {
	return dictionaryTestCase{
		key:   key,
		value: value,
		d:     Dictionary{key: value},
	}
}
