// Package structs defines the shapes and their methods.
package structs

import "math"

// Shape interface defines the methods that all shapes must implement.
type Shape interface {
	Area() float64
}

// Rectangle represents a rectangle shape.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle shape.
type Circle struct {
	Radius float64
}

// Perimeter calculates the perimeter (circumference) of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area calculates the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents a triangle shape.
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a triangle.
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
