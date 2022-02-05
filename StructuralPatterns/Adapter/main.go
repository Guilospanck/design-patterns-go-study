package main

import "fmt"

func main() {
	lightningPort := NewLightningPort(6.5)
	usbPort := NewUSBPort(65)

	client := NewClient()

	// OK
	fmt.Println("Sending data from a lightning port:")
	client.transferData(lightningPort)

	// ERROR
	// client.transferData(usbPort)

	// Using Adapter
	usbAdapter := NewUSBToLightningAdapter(*usbPort)
	lightningAdapted := usbAdapter.transform()

	fmt.Println("Sending data from a usb adapted to a lightning port:")
	client.transferData(lightningAdapted)
}
