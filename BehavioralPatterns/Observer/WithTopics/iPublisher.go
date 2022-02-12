package main

type IPublisher interface {
	AddSubscriber(topic string, subscriber ISubscriber)
	RemoveSubscriber(topic string, subscriber ISubscriber)
	Notify(topic string, subscriber ISubscriber)
	NotifyAllFromTopic(topic string)
	NotifyAll()
}
