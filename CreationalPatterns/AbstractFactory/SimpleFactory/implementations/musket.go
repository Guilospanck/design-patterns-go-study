package implementations

import "base/CreationalPatterns/AbstractFactory/SimpleFactory/interfaces"

// Concrete Product
type Musket struct {
	Gun
}

func NewMusket() interfaces.IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket",
			power: 1,
		},
	}
}
