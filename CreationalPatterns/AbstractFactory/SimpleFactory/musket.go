package main

// Concrete Product
type musket struct {
	Gun
}

func newMusket() iGun {
	return &musket{
		Gun: Gun{
			name:  "Musket",
			power: 1,
		},
	}
}
