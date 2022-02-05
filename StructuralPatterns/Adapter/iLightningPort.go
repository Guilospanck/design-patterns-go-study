package main

type ILightningPort interface {
	GetData() float64
}

type LightningPort struct {
	data float64
}

func (lightning *LightningPort) GetData() float64 {
	return lightning.data
}

func NewLightningPort(data float64) ILightningPort {
	return &LightningPort{
		data: data,
	}
}
