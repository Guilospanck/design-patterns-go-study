package main

type SensorsBuilder struct {
	obj        SensorOpts
	sensorType string
}

func (sb *SensorsBuilder) reset() {
	sb.obj = SensorOpts{}
}

func (sb *SensorsBuilder) setThreshold(threshold int) {
	sb.obj.Threshold = threshold
}

func (sb *SensorsBuilder) setHigher(higher bool) {
	sb.obj.Higher = higher
}

func (sb *SensorsBuilder) setLower(lower bool) {
	sb.obj.Lower = lower
}

func (sb *SensorsBuilder) setTSample(tsample uint) {
	sb.obj.TSample = tsample
}

func (sb *SensorsBuilder) setBorder(border bool) {
	sb.obj.Border = border
}

func (sb *SensorsBuilder) setRisingEdge(rising bool) {
	sb.obj.RisingEdge = rising
}

func (sb *SensorsBuilder) setFallingEdge(falling bool) {
	sb.obj.FallingEdge = falling
}

func (sb *SensorsBuilder) setBetween(between bool) {
	sb.obj.Between = between
}

func (sb *SensorsBuilder) setFirstThreshold(threshold int) {
	sb.obj.FirstThreshold = threshold
}

func (sb *SensorsBuilder) setSecondThreshold(threshold int) {
	sb.obj.SecondThreshold = threshold
}

func (sb *SensorsBuilder) setType(whatType uint) {
	sb.obj.Type = whatType
}

func (sb *SensorsBuilder) setSocialContact(socialContact string) {
	sb.obj.SocialContact = socialContact
}

func (sb *SensorsBuilder) getSensor() Sensor {
	sensorOpts := &SensorOpts{
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

	sb.reset()

	return Sensor{
		sb.sensorType: sensorOpts,
	}
}

func NewSensorsBuilder(sensorType string) *SensorsBuilder {
	return &SensorsBuilder{
		obj:        SensorOpts{},
		sensorType: sensorType,
	}
}
