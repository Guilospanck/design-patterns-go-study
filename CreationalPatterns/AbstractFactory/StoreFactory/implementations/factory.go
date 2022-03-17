package implementations

import (
	"base/CreationalPatterns/AbstractFactory/StoreFactory/interfaces"
	"fmt"
)

func GetSportsFactory(brand string) (interfaces.ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}
