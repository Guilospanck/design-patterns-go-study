package implementations

import "base/CreationalPatterns/AbstractFactory/StoreFactory/interfaces"

// Concrete Factory (implements the abstract factory interface)
type adidas struct{}

func (a *adidas) MakeShoe() interfaces.IShoe {
	return &adidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) MakeShirt() interfaces.IShirt {
	return &adidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
