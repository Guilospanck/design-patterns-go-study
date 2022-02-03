package main

import "fmt"

// Client code
func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(adidasShoe)
	printShoeDetails(nikeShoe)

	printShirtDetails(adidasShirt)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s iShoe) {
	fmt.Println("Shoe:")
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
	fmt.Println("=======================")
}

func printShirtDetails(s iShirt) {
	fmt.Println("Shirt:")
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
	fmt.Println("=======================")
}
