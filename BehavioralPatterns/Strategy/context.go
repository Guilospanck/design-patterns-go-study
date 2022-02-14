package main

type Context struct {
	strategy IStrategy
}

func (context *Context) SetStrategy(strategy IStrategy) {
	context.strategy = strategy
}

func (context *Context) DoSomething(a, b int) int {
	return context.strategy.Execute(a, b)
}

func NewContext() *Context {
	return &Context{}
}
