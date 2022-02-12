package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Subscriber struct {
	id string
}

func (s *Subscriber) Update(data IPublisher) {
	pub, ok := data.(*Publisher)
	if ok {
		log.Printf("[SUBSCRIBER %s] Data received: %s.\n", s.id, fmt.Sprintf("%d", pub.state))
	}
}

func (s *Subscriber) GetID() string {
	return s.id
}

func NewSubscriber() *Subscriber {
	id := uuid.New()

	return &Subscriber{
		id: id.String(),
	}
}
