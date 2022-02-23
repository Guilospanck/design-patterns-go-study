package visitor

import (
	"fmt"
)

type Rectangle struct {
	x, y int
}

func (r *Rectangle) Move(x, y int) {
	r.x = x
	r.y = y
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing rectangle at (%d, %d)\n", r.x, r.y)
}

func (r *Rectangle) Accept(v IVisitor) {
	v.visitRectangle(r)
}

func NewRectangle() *Rectangle {
	return &Rectangle{}
}
