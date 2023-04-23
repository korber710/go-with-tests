package iteration

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatFiveTimes(t *testing.T) {
	t.Parallel()
	for _, tc := range []repeatInput{
		newRepeatInput("a", 5),
		newRepeatInput("b", 10),
		newRepeatInput("z", 99),
	} {
		tc := tc
		t.Run(fmt.Sprintf("character %s repeat %d times", tc.Character, tc.RepeatCount), func(t *testing.T) {
			t.Parallel()
			got := Repeat(tc.Character, tc.RepeatCount)
			want := strings.Repeat(tc.Character, tc.RepeatCount)
			assert.Equal(t, got, want)
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

type repeatInput struct {
	Character   string
	RepeatCount int
}

func newRepeatInput(character string, repeatCount int) repeatInput {
	return repeatInput{
		Character: character, RepeatCount: repeatCount,
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))

	// Output: aaaaa
}
