package main

import "fmt"

// Liskov Substitution Principle -- This is a variation of this principle since it deals with inheritance, something
// that golang does not have

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width  int
	height int
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// Here is the problem,  breaks the UseIt function, we should not break behaviour that works
// with the parent type, in this case rectangle, check sol file for solution
// ********************************************************************************

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

// ********************************************************************************

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected area of ", expectedArea, ", but got", actualArea)
}

func main() {
	rec := &Rectangle{2, 3}
	UseIt(rec)

	sq := NewSquare(5)
	UseIt(sq)

	// Checking if solution works

	sq2 := NewSquare2(5)
	UseIt(sq2)
	sq2.ChangeSize(20)
	UseIt(sq2)
	sq2.ChangeSize(50)
	UseIt(sq2)
}
