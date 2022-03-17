package main

import (
	"base/CreationalPatterns/AbstractFactory/SimpleFactory/implementations"
	"base/CreationalPatterns/AbstractFactory/SimpleFactory/interfaces"
	"fmt"
)

func main() {
	ak47, _ := implementations.NewGunFactory("ak47")
	musket, _ := implementations.NewGunFactory("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun interfaces.IGun) {
	fmt.Printf("Gun: %s\nPower: %d\n=======\n", gun.GetName(), gun.GetPower())
}
