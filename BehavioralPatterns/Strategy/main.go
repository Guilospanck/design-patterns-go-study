package main

import "fmt"

func main() {
	// Create strategies
	addStrategy := NewConcreteStrategyAdd()
	subtractStrategy := NewConcreteStrategySubtract()
	multiplyStrategy := NewConcreteStrategyMultiply()

	// get context
	context := NewContext()

	// do something for each different strategy
	context.SetStrategy(addStrategy)
	fmt.Printf("%d + %d = %d\n", 40, 40, context.DoSomething(40, 40))

	context.SetStrategy(subtractStrategy)
	fmt.Printf("%d - %d = %d\n", 40, 40, context.DoSomething(40, 40))

	context.SetStrategy(multiplyStrategy)
	fmt.Printf("%d * %d = %d\n", 40, 40, context.DoSomething(40, 40))

}
