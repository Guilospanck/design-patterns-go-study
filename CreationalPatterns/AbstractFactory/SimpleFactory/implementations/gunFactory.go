package implementations

import (
	"base/CreationalPatterns/AbstractFactory/SimpleFactory/interfaces"
	"fmt"
)

// Look how we didn't have to have a "IGunFactory" interface
// because we're dealing with only different products, not different brands.
func NewGunFactory(gunType string) (interfaces.IGun, error) {
	switch gunType {
	case "ak47":
		return NewAK47(), nil
	case "musket":
		return NewMusket(), nil
	default:
		return nil, fmt.Errorf("gun type is not known")
	}
}
