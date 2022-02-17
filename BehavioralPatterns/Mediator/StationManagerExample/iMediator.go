package main

// IMediator
type IMediator interface {
	CanArrive(ITrain) bool
	NotifyAboutDeparture()
}
