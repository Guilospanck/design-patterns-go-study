package implementations

import "base/CreationalPatterns/AbstractFactory/StoreFactory/interfaces"

// Concrete Factory (implements the abstract factory interface)
type adidas struct {
	shoe  interfaces.IShoe
	shirt interfaces.IShirt
}

func (a *adidas) MakeShoe() interfaces.IShoe {
	a.shoe.SetLogo("adidas")
	a.shoe.SetSize(14)

	return a.shoe
}

func (a *adidas) MakeShirt() interfaces.IShirt {
	a.shirt.SetLogo("adidas")
	a.shirt.SetSize(14)

	return a.shirt
}

func NewAdidas(shoe interfaces.IShoe, shirt interfaces.IShirt) *adidas {
	return &adidas{
		shoe,
		shirt,
	}
}
