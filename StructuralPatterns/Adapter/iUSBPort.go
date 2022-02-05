package main

type IUSBPort interface {
	getData() int
}

type USBPort struct {
	data int
}

func (usb *USBPort) getData() int {
	return usb.data
}

func NewUSBPort(data int) *USBPort {
	return &USBPort{
		data: data,
	}
}
