package main

import (
	"base/StructuralPatterns/Adapter/interfaces"
	"fmt"
)

type client struct{}

func (c *client) TransferData(lightningPort interfaces.ILightningPort) bool {
	data := lightningPort.GetData()
	fmt.Printf("Data sent: %.2f\n", data)
	return true
}

func NewClient() *client {
	return &client{}
}
