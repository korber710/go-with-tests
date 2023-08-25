package structs_methods_and_interfaces

import "math"

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	Area() float64
}

func NewRectangle(length float64, width float64) Rectangle {
	return Rectangle{length, width}
}

func NewCircle(radius float64) Circle {
	return Circle{radius}
}

func NewTriangle(base float64, height float64) Triangle {
	return Triangle{base, height}
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
