package main

import (
	"base/CreationalPatterns/AbstractFactory/StoreFactory/implementations"
	"base/CreationalPatterns/AbstractFactory/StoreFactory/interfaces"
	"fmt"
)

// Client code
func main() {
	adidasFactory, _ := implementations.GetSportsFactory("adidas")
	nikeFactory, _ := implementations.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(adidasShoe)
	printShoeDetails(nikeShoe)

	printShirtDetails(adidasShirt)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s interfaces.IShoe) {
	fmt.Println("Shoe:")
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
	fmt.Println("=======================")
}

func printShirtDetails(s interfaces.IShirt) {
	fmt.Println("Shirt:")
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
	fmt.Println("=======================")
}
