package implementations

import (
	"base/CreationalPatterns/AbstractFactory/StoreFactory/interfaces"
	"fmt"
)

func GetSportsFactory(brand string) (interfaces.ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{
			shoe:  &adidasShoe{},
			shirt: &adidasShirt{},
		}, nil
	case "nike":
		return &nike{
			shoe:  &NikeShoe{},
			shirt: &NikeShirt{},
		}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}
