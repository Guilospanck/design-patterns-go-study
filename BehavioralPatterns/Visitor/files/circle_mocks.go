package visitor

type CircleMock struct {
	X, Y int
}

func (c *CircleMock) Move(x, y int) {

}

func (c *CircleMock) Draw() {

}

func (c *CircleMock) Accept(v IVisitor) {

}

func NewCircleMock() *CircleMock {
	return &CircleMock{
		X: 10,
		Y: 20,
	}
}
