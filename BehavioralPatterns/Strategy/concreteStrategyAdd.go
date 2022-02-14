package main

type ConcreteStrategyAdd struct{}

func (strategy *ConcreteStrategyAdd) Execute(a, b int) int {
	return a + b
}

func NewConcreteStrategyAdd() *ConcreteStrategyAdd {
	return &ConcreteStrategyAdd{}
}
