package main

import "fmt"

type client struct{}

func (c *client) transferData(lightningPort ILightningPort) bool {
	data := lightningPort.getData()
	fmt.Printf("Data sent: %.2f\n", data)
	return true
}

func NewClient() *client {
	return &client{}
}
