package main

type ConcreteStrategySubtract struct{}

func (strategy *ConcreteStrategySubtract) Execute(a, b int) int {
	return a - b
}

func NewConcreteStrategySubtract() *ConcreteStrategySubtract {
	return &ConcreteStrategySubtract{}
}
