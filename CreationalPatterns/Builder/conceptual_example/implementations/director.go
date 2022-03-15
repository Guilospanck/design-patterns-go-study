package implementations

import (
	"base/CreationalPatterns/Builder/conceptual_example/domain"
	"base/CreationalPatterns/Builder/conceptual_example/interfaces"
)

type director struct {
	builder interfaces.IBuilder
}

func (d *director) SetBuilder(b interfaces.IBuilder) {
	d.builder = b
}

func (d *director) BuildHouse() domain.House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}

func NewDirector(builder interfaces.IBuilder) *director {
	return &director{
		builder: builder,
	}
}
