package implementations

import "base/CreationalPatterns/AbstractFactory/SimpleFactory/interfaces"

// Concrete Product
type Ak47 struct {
	Gun
}

func NewAK47() interfaces.IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47",
			power: 4,
		},
	}
}
