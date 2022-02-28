package main

type ConcreteStrategyAddMock struct{}

var AddStrategyCalled bool = false

func (strategy *ConcreteStrategyAddMock) Execute(a, b int) int {
	AddStrategyCalled = true
	return a + b
}

func NewConcreteStrategyAddMock() *ConcreteStrategyAddMock {
	return &ConcreteStrategyAddMock{}
}
