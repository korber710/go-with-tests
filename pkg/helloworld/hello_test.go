package hello

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloToPeople(t *testing.T) {
	t.Parallel()
	for _, name := range []string{"Steve", "Bob", "Jane"} {
		name := name
		t.Run(fmt.Sprintf("saying hello to %s", name), func(t *testing.T) {
			t.Parallel()
			got := Hello(name, "")
			want := fmt.Sprintf("Hello, %s", name)
			assert.Equal(t, got, want)
		})
	}
}

func TestHelloWithEmptyString(t *testing.T) {
	t.Parallel()
	t.Run("saying hello to world", func(t *testing.T) {
		t.Parallel()
		got := Hello("", "")
		want := "Hello, World"
		assert.Equal(t, got, want)
	})
}

func TestHelloWithNonEnglishGreeting(t *testing.T) {
	t.Parallel()
	for _, tc := range []helloInput{
		newHelloInput("Steve", "French", "Bonjour, "),
		newHelloInput("Bob", "Spanish", "Hola, "),
		newHelloInput("Jane", "German", "Hello, "),
	} {
		tc := tc
		t.Run(fmt.Sprintf("saying hello to %s in %s", tc.Name, tc.Language), func(t *testing.T) {
			t.Parallel()
			got := Hello(tc.Name, tc.Language)
			want := fmt.Sprintf("%s%s", tc.Greeting, tc.Name)
			assert.Equal(t, got, want)
		})
	}
}

type helloInput struct {
	Name     string
	Language string
	Greeting string
}

func newHelloInput(name string, language string, greeting string) helloInput {
	return helloInput{
		Name: name, Language: language, Greeting: greeting,
	}
}
