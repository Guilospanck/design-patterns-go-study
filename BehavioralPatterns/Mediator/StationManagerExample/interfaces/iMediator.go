package interfaces

// IMediator
type IMediator interface {
	CanArrive(ITrain) bool
	NotifyAboutDeparture()
}
