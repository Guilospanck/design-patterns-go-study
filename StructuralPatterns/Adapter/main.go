package main

import (
	"base/StructuralPatterns/Adapter/adapters"
	"base/StructuralPatterns/Adapter/domain"
	"base/StructuralPatterns/Adapter/implementations"
	"fmt"
)

func main() {
	lightningPort := implementations.NewLightningPort(6.5)
	usbPort := implementations.NewUSBPort(65)

	client := domain.NewClient()

	// OK
	fmt.Println("Sending data from a lightning port:")
	client.TransferData(lightningPort)

	// ERROR
	// client.transferData(usbPort)

	// Using Adapter
	usbAdapter := adapters.NewUSBToLightningAdapter(usbPort)
	lightningAdapted := usbAdapter.Transform()

	fmt.Println("Sending data from a usb adapted to a lightning port:")
	client.TransferData(lightningAdapted)
}
