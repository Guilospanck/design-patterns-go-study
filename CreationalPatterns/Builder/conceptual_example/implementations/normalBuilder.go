package implementations

import "base/CreationalPatterns/Builder/conceptual_example/domain"

type normalBuilder struct {
	windowType, doorType string
	floor                int
}

func (n *normalBuilder) SetWindowType() {
	n.windowType = "Wooden Window"
}

func (n *normalBuilder) SetDoorType() {
	n.doorType = "Wooden Door"
}

func (n *normalBuilder) SetNumFloor() {
	n.floor = 2
}

func (n *normalBuilder) GetHouse() domain.House {
	return domain.House{
		WindowType: n.windowType,
		DoorType:   n.doorType,
		Floor:      n.floor,
	}
}

func NewNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}
