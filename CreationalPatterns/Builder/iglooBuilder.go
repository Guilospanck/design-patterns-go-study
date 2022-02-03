package main

type iglooBuilder struct {
	windowType, doorType string
	floor                int
}

func (i *iglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
}

func (i *iglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
}

func (i *iglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *iglooBuilder) getHouse() house {
	return house{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}

func newiglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}
