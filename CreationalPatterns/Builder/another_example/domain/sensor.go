package domain

type SensorOpts struct {
	Threshold       int
	Higher          bool
	Lower           bool
	TSample         uint
	Border          bool
	RisingEdge      bool
	FallingEdge     bool
	Between         bool
	FirstThreshold  int
	SecondThreshold int
	Type            uint
	SocialContact   string
}

type Sensor map[string]*SensorOpts
