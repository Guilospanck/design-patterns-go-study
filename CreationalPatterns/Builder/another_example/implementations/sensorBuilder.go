package implementations

import "base/CreationalPatterns/Builder/another_example/domain"

type SensorsBuilder struct {
	obj        domain.SensorOpts
	sensorType string
}

func (sb *SensorsBuilder) Reset() {
	sb.obj = domain.SensorOpts{}
}

func (sb *SensorsBuilder) SetThreshold(threshold int) {
	sb.obj.Threshold = threshold
}

func (sb *SensorsBuilder) SetHigher(higher bool) {
	sb.obj.Higher = higher
}

func (sb *SensorsBuilder) SetLower(lower bool) {
	sb.obj.Lower = lower
}

func (sb *SensorsBuilder) SetTSample(tsample uint) {
	sb.obj.TSample = tsample
}

func (sb *SensorsBuilder) SetBorder(border bool) {
	sb.obj.Border = border
}

func (sb *SensorsBuilder) SetRisingEdge(rising bool) {
	sb.obj.RisingEdge = rising
}

func (sb *SensorsBuilder) SetFallingEdge(falling bool) {
	sb.obj.FallingEdge = falling
}

func (sb *SensorsBuilder) SetBetween(between bool) {
	sb.obj.Between = between
}

func (sb *SensorsBuilder) SetFirstThreshold(threshold int) {
	sb.obj.FirstThreshold = threshold
}

func (sb *SensorsBuilder) SetSecondThreshold(threshold int) {
	sb.obj.SecondThreshold = threshold
}

func (sb *SensorsBuilder) SetType(whatType uint) {
	sb.obj.Type = whatType
}

func (sb *SensorsBuilder) SetSocialContact(socialContact string) {
	sb.obj.SocialContact = socialContact
}

func (sb *SensorsBuilder) GetSensor() domain.Sensor {
	sensorOpts := &domain.SensorOpts{
		Threshold:       sb.obj.Threshold,
		Higher:          sb.obj.Higher,
		Lower:           sb.obj.Lower,
		TSample:         sb.obj.TSample,
		Border:          sb.obj.Border,
		RisingEdge:      sb.obj.RisingEdge,
		FallingEdge:     sb.obj.FallingEdge,
		Between:         sb.obj.Between,
		FirstThreshold:  sb.obj.FirstThreshold,
		SecondThreshold: sb.obj.SecondThreshold,
		Type:            sb.obj.Type,
		SocialContact:   sb.obj.SocialContact,
	}

	sb.Reset()

	return domain.Sensor{
		sb.sensorType: sensorOpts,
	}
}

func NewSensorsBuilder(sensorType string) *SensorsBuilder {
	return &SensorsBuilder{
		obj:        domain.SensorOpts{},
		sensorType: sensorType,
	}
}
