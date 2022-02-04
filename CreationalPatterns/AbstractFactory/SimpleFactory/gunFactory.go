package main

import "fmt"

// Look how we didn't have to have a "IGunFactory" interface
// because we're dealing with only different products, not different brands.
func newGunFactory(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAK47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("gun type is not known")
	}
}
