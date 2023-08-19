package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()
	for _, tc := range []sumInput{
		newSumInput([]int{1, 2, 3, 4, 5}, 15),
		newSumInput([]int{1, 1, 1, 1, 1}, 5),
	} {
		tc := tc
		t.Run(fmt.Sprintf("summing %v and expecting %d", tc.Numbers, tc.Sum), func(t *testing.T) {
			t.Parallel()
			got := Sum(tc.Numbers)
			want := tc.Sum
			assert.Equal(t, got, want)
		})
	}
}

func TestSumAll(t *testing.T) {
	t.Parallel()
	for _, tc := range []sumAllInput{
		newSumAllInput([][]int{{1, 2}, {0, 9}}, []int{3, 9}),
		newSumAllInput([][]int{{1, 2}, {0, 9}, {1, 2}}, []int{3, 9, 3}),
	} {
		tc := tc
		t.Run(fmt.Sprintf("summing %v and expecting %d", tc.NumberSlices, tc.SumSlice), func(t *testing.T) {
			t.Parallel()
			got := SumAll(tc.NumberSlices...)
			want := tc.SumSlice
			assert.Equal(t, got, want)
		})
	}
}

type sumAllInput struct {
	NumberSlices [][]int
	SumSlice 	 []int
}

func newSumAllInput(numberSlices [][]int, sumSlice []int) sumAllInput {
	return sumAllInput{
		NumberSlices: numberSlices,
		SumSlice: sumSlice,
	}
}

type sumInput struct {
	Numbers []int
	Sum     int
}

func newSumInput(numbers []int, sum int) sumInput {
	return sumInput{
		Numbers: numbers, Sum: sum,
	}
}
