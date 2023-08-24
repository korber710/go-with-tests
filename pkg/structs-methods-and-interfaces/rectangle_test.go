package structs_methods_and_interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		length    float64
		width     float64
		perimeter float64
	}{
		{10.0, 10.0, 40.0},
		{5.0, 10.0, 30.0},
	} {
		tc := tc
		t.Run(fmt.Sprintf("calculate perimeter for %.2f and %.2f and expecting %.2f", tc.length, tc.width, tc.perimeter), func(t *testing.T) {
			t.Parallel()
			got := Perimeter(tc.length, tc.width)
			want := tc.perimeter
			assert.Equal(t, got, want)
		})
	}
}

func TestArea(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		length float64
		width  float64
		area   float64
	}{
		{10.0, 10.0, 100.0},
		{5.0, 10.0, 50.0},
	} {
		tc := tc
		t.Run(fmt.Sprintf("calculate area for %.2f and %.2f and expecting %.2f", tc.length, tc.width, tc.area), func(t *testing.T) {
			t.Parallel()
			got := Area(tc.length, tc.width)
			want := tc.area
			assert.Equal(t, got, want)
		})
	}
}
