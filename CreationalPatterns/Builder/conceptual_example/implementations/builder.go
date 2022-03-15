package implementations

import "base/CreationalPatterns/Builder/conceptual_example/interfaces"

func GetBuilder(builderType string) interfaces.IBuilder {
	switch builderType {
	case "normal":
		return NewNormalBuilder()
	case "igloo":
		return NewIglooBuilder()
	default:
		return nil
	}
}
