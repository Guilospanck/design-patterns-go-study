package implementations

type LightningPort struct {
	data float64
}

func (lightning *LightningPort) GetData() float64 {
	return lightning.data
}

func NewLightningPort(data float64) *LightningPort {
	return &LightningPort{
		data: data,
	}
}
