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
			hasPerimeter := tc.r.Perimeter()
			wantedPerimeter := tc.perimeter
			assert.Equal(t, hasPerimeter, wantedPerimeter)
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

func TestCalculatingTriangleArea(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		t    Triangle
		area float64
	}{
		{NewTriangle(1.0, 1.0), 0.5},
		{NewTriangle(5.0, 10.0), 25.0},
	} {
		tc := tc
		t.Run(fmt.Sprintf("calculate area for %.2f and %.2f and expecting %.2f", tc.t.base, tc.t.height, tc.area), func(t *testing.T) {
			t.Parallel()
			checkArea(t, tc.t, tc.area)
		})
	}
}

func checkArea(t testing.TB, shape Shape, wantedArea float64) {
	t.Helper()
	hasArea := shape.Area()
	assert.Equal(t, hasArea, wantedArea)
}
