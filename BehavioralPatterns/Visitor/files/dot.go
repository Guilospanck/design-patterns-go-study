package visitor

import (
	"fmt"
)

type Dot struct {
	x, y int
}

func (d *Dot) Move(x, y int) {
	d.x = x
	d.y = y
}

func (d *Dot) Draw() {
	fmt.Printf("Drawing rectangle at (%d, %d)\n", d.x, d.y)
}

func (d *Dot) Accept(v IVisitor) {
	v.visitDot(d)
}

func NewDot() *Dot {
	return &Dot{}
}
