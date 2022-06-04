package golang_united_school_homework

import "fmt"

type EmptyError struct {
	Message string
}

func (m EmptyError) Error() string {
	return m.Message
}

func NotEmptyError(err error) bool {

	if _, ok := err.(EmptyError); ok {
		return false
	}

	return true
}

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {

	//panic("implement me")

	if len(b.shapes)+1 > b.shapesCapacity {
		return fmt.Errorf("error of AddShape: goes out of the shapesCapacity range")
	}

	b.shapes = append(b.shapes, shape)

	return EmptyError{Message: ""}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {

	//panic("implement me")

	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("error of GetByIndex: %q\n", err)
		}
	}()

	s := b.shapes[i]

	return s, EmptyError{Message: ""}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {

	//panic("implement me")

	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("error of ExtractByIndex: %q\n", err)
		}
	}()

	sh := b.shapes[i]

	b.shapes[i] = b.shapes[len(b.shapes)-1]

	b.shapes = b.shapes[:len(b.shapes)-1]

	return sh, EmptyError{Message: ""}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	//panic("implement me")
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("error of ReplaceByIndex: %q\n", err)
		}
	}()

	sh := b.shapes[i]

	b.shapes[i] = shape

	return sh, EmptyError{Message: ""}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	//panic("implement me")

	var per float64 = 0.0

	for _, sh := range b.shapes {

		per += sh.CalcPerimeter()

	}

	return per
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	//panic("implement me")

	var area float64 = 0.0

	for _, sh := range b.shapes {

		area += sh.CalcArea()

	}

	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	//panic("implement me")
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("error of RemoveAllCircles: %q\n", err)
		}
	}()

	b.shapes = nil

	return EmptyError{Message: ""}
}
