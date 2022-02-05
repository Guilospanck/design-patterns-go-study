package main

type USBToLightningAdapter struct {
	usb *USBPort
}

func (adapter *USBToLightningAdapter) transform() *LightningPort {
	lightningPort := NewLightningPort(float64(adapter.usb.data))

	return lightningPort
}

func NewUSBToLightningAdapter(usbPort USBPort) *USBToLightningAdapter {
	return &USBToLightningAdapter{
		usb: &usbPort,
	}
}
