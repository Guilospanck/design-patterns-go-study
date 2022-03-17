package implementations

type Gun struct {
	name  string
	power int
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetPower() int {
	return g.power
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func NewGun() *Gun {
	return &Gun{}
}
