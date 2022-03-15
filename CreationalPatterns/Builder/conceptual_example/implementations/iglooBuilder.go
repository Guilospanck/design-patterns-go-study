package implementations

import "base/CreationalPatterns/Builder/conceptual_example/domain"

type iglooBuilder struct {
	windowType, doorType string
	floor                int
}

func (i *iglooBuilder) SetWindowType() {
	i.windowType = "Snow Window"
}

func (i *iglooBuilder) SetDoorType() {
	i.doorType = "Snow Door"
}

func (i *iglooBuilder) SetNumFloor() {
	i.floor = 1
}

func (i *iglooBuilder) GetHouse() domain.House {
	return domain.House{
		WindowType: i.windowType,
		DoorType:   i.doorType,
		Floor:      i.floor,
	}
}

func NewIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}
