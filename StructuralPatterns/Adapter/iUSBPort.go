package main

type IUSBPort interface {
	GetData() int
}

type USBPort struct {
	data int
}

func (usb *USBPort) GetData() int {
	return usb.data
}

func NewUSBPort(data int) IUSBPort {
	return &USBPort{
		data: data,
	}
}
