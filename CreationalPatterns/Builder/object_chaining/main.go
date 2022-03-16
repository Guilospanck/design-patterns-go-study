package main

import (
	"base/CreationalPatterns/Builder/object_chaining/implementations"
	"fmt"
)

func main() {
	personBuilder := implementations.NewPersonBuilder()

	// first person
	personBuilder.
		Named("Guilherme").
		Lives().
		At("Brazil").
		WithPostalCode("35600-000").
		Works().
		As("Software Engineer").
		For("IBM").
		In("Brazil").
		WithSalary(40000)

	// builds first Person
	firstPerson := personBuilder.Build()
	fmt.Printf("%+v\n", firstPerson)

	// resets person
	personBuilder.Reset()

	// builds second person
	personBuilder.
		Named("Larry").
		Lives().
		At("Europe").
		WithPostalCode("888888-000").
		Works().
		As("Agricultor").
		For("Embrapa").
		In("Germany").
		WithSalary(457894)

	// builds second person
	secPerson := personBuilder.Build()
	fmt.Printf("%+v\n", secPerson)

}
