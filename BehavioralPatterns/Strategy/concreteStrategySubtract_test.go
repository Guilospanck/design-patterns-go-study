package main

type ConcreteStrategySubtractMock struct{}

var SubtractStrategyCalled bool = false

func (strategy *ConcreteStrategySubtractMock) Execute(a, b int) int {
	SubtractStrategyCalled = true
	return a - b
}

func NewConcreteStrategySubtractMock() *ConcreteStrategySubtractMock {
	return &ConcreteStrategySubtractMock{}
}
