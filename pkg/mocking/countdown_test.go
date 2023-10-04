package mocking

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Parallel()

	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	assert.Equal(t, got, want)
}
