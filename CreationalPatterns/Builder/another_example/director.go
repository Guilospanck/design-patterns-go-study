package main

type director struct {
	builder ISensorsBuilder
}

func (d *director) setBuilder(b ISensorsBuilder) {
	d.builder = b
}

func (d *director) createSensor(opts SensorOpts) Sensor {
	d.builder.reset()
	d.builder.setThreshold(opts.Threshold)
	d.builder.setHigher(opts.Higher)
	d.builder.setLower(opts.Lower)
	d.builder.setTSample(opts.TSample)
	d.builder.setBorder(opts.Border)
	d.builder.setRisingEdge(opts.RisingEdge)
	d.builder.setFallingEdge(opts.FallingEdge)
	d.builder.setBetween(opts.Between)
	d.builder.setFirstThreshold(opts.FirstThreshold)
	d.builder.setSecondThreshold(opts.SecondThreshold)
	d.builder.setType(opts.Type)
	d.builder.setSocialContact(opts.SocialContact)
	return d.builder.getSensor()
}

func NewDirector(builder ISensorsBuilder) *director {
	return &director{
		builder: builder,
	}
}
