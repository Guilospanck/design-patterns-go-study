package main

type normalBuilder struct {
	windowType, doorType string
	floor                int
}

func (n *normalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *normalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *normalBuilder) setNumFloor() {
	n.floor = 2
}

func (n *normalBuilder) getHouse() house {
	return house{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}
