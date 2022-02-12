package main

// Publisher (or Subject) interface
type IPublisher interface {
	AddSubscriber(subscriber ISubscriber)
	RemoveSubscriber(subscriber ISubscriber)
	NotifyAll()
}
