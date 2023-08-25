package structs_methods_and_interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatingRectanglePerimeter(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		r         Rectangle
		perimeter float64
	}{
		{NewRectangle(10.0, 10.0), 40.0},
		{NewRectangle(5.0, 10.0), 30.0},
	} {
		tc := tc
		t.Run(fmt.Sprintf("calculate perimeter for %.2f and %.2f and expecting %.2f", tc.r.length, tc.r.width, tc.perimeter), func(t *testing.T) {
			t.Parallel()
			got := tc.r.Perimeter()
			want := tc.perimeter
			assert.Equal(t, got, want)
		})
	}
}

func TestCalculatingRectangleArea(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		r    Rectangle
		area float64
	}{
		{NewRectangle(10.0, 10.0), 100.0},
		{NewRectangle(10.0, 5.0), 50.0},
	} {
		tc := tc
		t.Run(fmt.Sprintf("calculate area for %.2f and %.2f and expecting %.2f", tc.r.length, tc.r.width, tc.area), func(t *testing.T) {
			t.Parallel()
			checkArea(t, tc.r, tc.area)
		})
	}
}

func TestCalculatingCircleArea(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		c    Circle
		area float64
	}{
		{NewCircle(1.0), 3.141592653589793},
		{NewCircle(5.0), 78.53981633974483},
	} {
		tc := tc
		t.Run(fmt.Sprintf("calculate area for %.2f and expecting %.2f", tc.c.radius, tc.area), func(t *testing.T) {
			t.Parallel()
			checkArea(t, tc.c, tc.area)
		})
	}
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	assert.Equal(t, got, want)
}
