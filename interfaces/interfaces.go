package interfaces

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}