package main

// Concrete Product
type ak47 struct {
	Gun
}

func newAK47() iGun {
	return &ak47{
		Gun: Gun{
			name:  "AK47",
			power: 4,
		},
	}
}
