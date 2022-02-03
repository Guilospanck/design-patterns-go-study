package main

type director struct {
	builder iBuilder
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func newDirector(builder iBuilder) *director {
	return &director{
		builder: builder,
	}
}
