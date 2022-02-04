package main

// Abstract Product
type iGun interface {
	getName() string
	setName(string)
	getPower() int
	setPower(int)
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getPower() int {
	return g.power
}

func (g *Gun) setPower(power int) {
	g.power = power
}
