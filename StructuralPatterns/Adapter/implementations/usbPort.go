package implementations

type USBPort struct {
	data int
}

func (usb *USBPort) GetData() int {
	return usb.data
}

func NewUSBPort(data int) *USBPort {
	return &USBPort{
		data: data,
	}
}
