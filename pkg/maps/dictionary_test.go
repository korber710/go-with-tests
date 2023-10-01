package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")

	want := "this is just a test"
	assert.Equal(t, got, want)
}
