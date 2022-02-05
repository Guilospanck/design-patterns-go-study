package main

type ILightningPort interface {
	getData() float64
}

type LightningPort struct {
	data float64
}

func (lightning *LightningPort) getData() float64 {
	return lightning.data
}

func NewLightningPort(data float64) *LightningPort {
	return &LightningPort{
		data: data,
	}
}
