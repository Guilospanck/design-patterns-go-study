package interfaces

// Abstract Product
type IGun interface {
	GetName() string
	SetName(string)
	GetPower() int
	SetPower(int)
}
