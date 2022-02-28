package main

type ConcreteStrategyMultiplyMock struct{}

var StrategyMultiplyCalled bool = false

func (strategy *ConcreteStrategyMultiplyMock) Execute(a, b int) int {
	StrategyMultiplyCalled = true
	return a * b
}

func NewConcreteStrategyMultiplyMock() *ConcreteStrategyMultiplyMock {
	return &ConcreteStrategyMultiplyMock{}
}
