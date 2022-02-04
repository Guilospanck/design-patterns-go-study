package main

import "fmt"

func main() {
	ak47, _ := newGunFactory("ak47")
	musket, _ := newGunFactory("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun iGun) {
	fmt.Printf("Gun: %s\nPower: %d\n=======\n", gun.getName(), gun.getPower())
}
