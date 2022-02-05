package main

type IClient interface {
	transferData(lightningPort ILightningPort) bool
}
