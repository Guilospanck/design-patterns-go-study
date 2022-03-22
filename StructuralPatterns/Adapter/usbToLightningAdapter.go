package main

import (
	"base/StructuralPatterns/Adapter/implementations"
	"base/StructuralPatterns/Adapter/interfaces"
)

type USBToLightningAdapter struct {
	usb interfaces.IUSBPort
}

func (adapter *USBToLightningAdapter) transform() interfaces.ILightningPort {
	data := adapter.usb.GetData()
	lightningPort := implementations.NewLightningPort(float64(data))

	return lightningPort
}

func NewUSBToLightningAdapter(usbPort interfaces.IUSBPort) *USBToLightningAdapter {
	return &USBToLightningAdapter{
		usb: usbPort,
	}
}
