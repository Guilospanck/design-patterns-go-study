package main

type ContextMock struct {
	strategy IStrategy
}

func (context *ContextMock) SetStrategy(strategy IStrategy) {
	context.strategy = strategy
}

func (context *ContextMock) DoSomething(a, b int) int {
	return context.strategy.Execute(a, b)
}

func NewContextMock() *ContextMock {
	return &ContextMock{}
}
