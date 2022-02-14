package main

type ConcreteStrategyMultiply struct{}

func (strategy *ConcreteStrategyMultiply) Execute(a, b int) int {
	return a * b
}

func NewConcreteStrategyMultiply() *ConcreteStrategyMultiply {
	return &ConcreteStrategyMultiply{}
}
