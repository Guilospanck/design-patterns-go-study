package main

type USBToLightningAdapter struct {
	usb IUSBPort
}

func (adapter *USBToLightningAdapter) transform() ILightningPort {
	data := adapter.usb.GetData()
	lightningPort := NewLightningPort(float64(data))

	return lightningPort
}

func NewUSBToLightningAdapter(usbPort IUSBPort) *USBToLightningAdapter {
	return &USBToLightningAdapter{
		usb: usbPort,
	}
}
