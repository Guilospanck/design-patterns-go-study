package visitor

import (
	"fmt"

	"base/BehavioralPatterns/Visitor/interfaces"
)

type Circle struct {
	x, y int
}

func (c *Circle) Move(x, y int) {
	c.x = x
	c.y = y
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing circle at (%d, %d)\n", c.x, c.y)
}

func (c *Circle) Accept(v interfaces.IVisitor) {
	v.visitCircle(c)
}

func NewCircle() *Circle {
	return &Circle{}
}
