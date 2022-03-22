package interfaces

type IClient interface {
	TransferData(lightningPort ILightningPort) bool
}
