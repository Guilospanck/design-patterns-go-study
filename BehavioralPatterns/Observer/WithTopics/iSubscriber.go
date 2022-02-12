package main

type ISubscriber interface {
	Update(data IPublisher)
	GetID() string
}
