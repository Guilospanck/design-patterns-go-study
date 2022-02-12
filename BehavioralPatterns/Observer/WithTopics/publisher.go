package main

import "log"

type Publisher struct {
	subscribers map[string][]ISubscriber
	state       int
}

func (p *Publisher) AddSubscriber(topic string, subscriber ISubscriber) {
	log.Printf("[PUBLISHER] Adding subscriber %s to topic %s.\n", subscriber.GetID(), topic)

	if _, found := p.subscribers[topic]; found {
		p.subscribers[topic] = append(p.subscribers[topic], subscriber)
	} else {
		subs := []ISubscriber{subscriber}
		p.subscribers[topic] = subs
	}
}

func (p *Publisher) RemoveSubscriber(topic string, subscriber ISubscriber) {
	log.Printf("[PUBLISHER] Removing subscriber %s from topic %s.\n", subscriber.GetID(), topic)

	if _, found := p.subscribers[topic]; !found {
		return
	}

	subs := p.subscribers[topic]
	p.subscribers[topic] = _removeSubscriberFromArray(subs, subscriber)
}

func (p *Publisher) Notify(topic string, subscriber ISubscriber) {
	log.Printf("[PUBLISHER] Notifying subscriber %s from topic %s.\n", subscriber.GetID(), topic)

	if _, found := p.subscribers[topic]; !found {
		return
	}

	subs := p.subscribers[topic]
	for _, sub := range subs {
		if sub.GetID() == subscriber.GetID() {
			sub.Update(p)
		}
	}
}

func (p *Publisher) NotifyAllFromTopic(topic string) {
	log.Printf("[PUBLISHER] Notifying all subscribers from topic %s.\n", topic)

	if _, found := p.subscribers[topic]; !found {
		return
	}

	subs := p.subscribers[topic]
	for _, sub := range subs {
		sub.Update(p)
	}
}

func (p *Publisher) NotifyAll() {
	log.Printf("[PUBLISHER] Notifying all subscribers from all topic.\n")

	for _, subs := range p.subscribers {
		for _, sub := range subs {
			sub.Update(p)
		}
	}
}

func _removeSubscriberFromArray(subs []ISubscriber, subscriber ISubscriber) []ISubscriber {
	subsLength := len(subs)

	for index, sub := range subs {
		if sub.GetID() == subscriber.GetID() {
			subs[index], subs[subsLength-1] = subs[subsLength-1], subs[index]
			return subs[:subsLength-1]
		}
	}

	return subs
}

// some state change
func (p *Publisher) UpdateState() {
	p.state++
}

func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[string][]ISubscriber),
		state:       0,
	}
}
