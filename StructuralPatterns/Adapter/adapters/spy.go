package adapters

import "github.com/stretchr/testify/mock"

type usbPortSpy struct {
	mock.Mock
	data int
}

func (spy *usbPortSpy) GetData() int {
	spy.Called()
	return spy.data
}

func NewUSBPortSpy(data int) *usbPortSpy {
	return &usbPortSpy{data: data}
}
