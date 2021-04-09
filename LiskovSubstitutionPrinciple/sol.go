package main

// There is no need no create a set width and height for the square since it is a rectangle
// it should have the same height and width since the beginning

type Square2 struct {
	Rectangle
}

func NewSquare2(size int) *Square2 {
	return &Square2{Rectangle{width: size, height: size}}
}

// We cannot just update one only width or height, we should update both
// we should not change SetWidth and SetHeight methods since changing those may lead breaking the code
// ChangeSize is the optimal option to change height and width of the current square
func (s *Square2) ChangeSize(size int) {
	s.height = size
	s.width = size
}
