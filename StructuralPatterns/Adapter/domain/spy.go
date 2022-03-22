package domain

import "github.com/stretchr/testify/mock"

type lightningPortSpy struct {
	mock.Mock
	data float64
}

func (spy *lightningPortSpy) GetData() float64 {
	spy.Called()
	return spy.data
}

func NewLightningPortSpy(data float64) *lightningPortSpy {
	return &lightningPortSpy{data: data}
}
