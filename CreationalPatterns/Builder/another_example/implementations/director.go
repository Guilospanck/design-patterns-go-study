package implementations

import (
	"base/CreationalPatterns/Builder/another_example/domain"
	"base/CreationalPatterns/Builder/another_example/interfaces"
)

type director struct {
	builder interfaces.ISensorsBuilder
}

func (d *director) SetBuilder(b interfaces.ISensorsBuilder) {
	d.builder = b
}

func (d *director) CreateSensor(opts domain.SensorOpts) domain.Sensor {
	d.builder.Reset()
	d.builder.SetThreshold(opts.Threshold)
	d.builder.SetHigher(opts.Higher)
	d.builder.SetLower(opts.Lower)
	d.builder.SetTSample(opts.TSample)
	d.builder.SetBorder(opts.Border)
	d.builder.SetRisingEdge(opts.RisingEdge)
	d.builder.SetFallingEdge(opts.FallingEdge)
	d.builder.SetBetween(opts.Between)
	d.builder.SetFirstThreshold(opts.FirstThreshold)
	d.builder.SetSecondThreshold(opts.SecondThreshold)
	d.builder.SetType(opts.Type)
	d.builder.SetSocialContact(opts.SocialContact)
	return d.builder.GetSensor()
}

func NewDirector(builder interfaces.ISensorsBuilder) *director {
	return &director{
		builder: builder,
	}
}
