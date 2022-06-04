package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {

	return t.Side * float64(3)
}

func (t Triangle) CalcArea() float64 {

	return math.Pow(t.Side, float64(2)) * math.Sqrt(3) / 4
}
