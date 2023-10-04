package dependency_injection

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Parallel()

	buffer := bytes.Buffer{}

	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris\n"

	assert.Equal(t, got, want)
}
