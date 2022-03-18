package implementations

import "base/CreationalPatterns/AbstractFactory/StoreFactory/interfaces"

// Concrete Factory (implements the abstract factory interface)
type nike struct {
	shoe  interfaces.IShoe
	shirt interfaces.IShirt
}

func (n *nike) MakeShoe() interfaces.IShoe {
	n.shoe.SetLogo("nike")
	n.shoe.SetSize(14)

	return n.shoe
}

func (n *nike) MakeShirt() interfaces.IShirt {
	n.shirt.SetLogo("nike")
	n.shirt.SetSize(14)

	return n.shirt
}

func NewNike(shoe interfaces.IShoe, shirt interfaces.IShirt) *nike {
	return &nike{
		shoe,
		shirt,
	}
}
