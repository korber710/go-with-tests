package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	for _, tc := range []int{1, 2, 3} {
		tc := tc
		t.Run(fmt.Sprintf("adding %d with %d", tc, tc), func(t *testing.T) {
			t.Parallel()
			got := Add(tc, tc)
			want := tc * 2
			assert.Equal(t, want, got)
		})
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

	// Output: 6
}
