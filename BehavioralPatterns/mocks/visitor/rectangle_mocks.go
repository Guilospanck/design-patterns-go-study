package mocks

type RectangleMock struct {
	X, Y int
}

func NewRectangleMock() *RectangleMock {
	return &RectangleMock{
		X: 10,
		Y: 20,
	}
}
