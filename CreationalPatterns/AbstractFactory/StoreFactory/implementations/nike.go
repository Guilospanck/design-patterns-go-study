package implementations

import "base/CreationalPatterns/AbstractFactory/StoreFactory/interfaces"

// Concrete Factory (implements the abstract factory interface)
type nike struct{}

func (n *nike) MakeShoe() interfaces.IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) MakeShirt() interfaces.IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}
