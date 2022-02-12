package main

import "fmt"

type Publisher struct {
	data        Data
	subscribers []ISubscriber
}

/*
 Concrete implementations of IPublisher
*/
func (p *Publisher) AddSubscriber(subscriber ISubscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p *Publisher) RemoveSubscriber(subscriber ISubscriber) {
	p.subscribers = _removeFromSlice(p.subscribers, subscriber)
}

func (p *Publisher) NotifyAll() {
	for _, subscriber := range p.subscribers {
		subscriber.Update(p.data)
	}
}

func _removeFromSlice(subscribers []ISubscriber, subscriber ISubscriber) []ISubscriber {
	length := len(subscribers)

	for index, sub := range subscribers {
		if sub.GetID() == subscriber.GetID() {
			subscribers[length-1], subscribers[index] = subscribers[index], subscribers[length-1]
			return subscribers[:length-1]
		}
	}

	return subscribers
}

/*
  Function of this publisher
*/
func (p *Publisher) updateAvailabilityOfData() {
	fmt.Printf("Informing availability of data %+v...\n", p.data)
	p.NotifyAll()
}

func NewPublisher(data Data) *Publisher {
	return &Publisher{
		data: data,
	}
}
