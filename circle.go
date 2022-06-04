package golang_united_school_homework

import "math"

const (
	pi float64 = math.Pi
)

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {

	return float64(2) * pi * c.Radius
}

func (c Circle) CalcArea() float64 {

	return pi * math.Pow(c.Radius, float64(2))
}
