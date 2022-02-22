package mocks

type DotMock struct {
	X, Y int
}

func NewDotMock() *DotMock {
	return &DotMock{
		X: 10,
		Y: 20,
	}
}
