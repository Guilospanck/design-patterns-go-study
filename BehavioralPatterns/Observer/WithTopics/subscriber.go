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

func NewSubscriber(optionalId ...uuid.UUID) *Subscriber {
	id := uuid.New()
	if len(optionalId) > 0 {
		id = optionalId[0]
	}

	return &Subscriber{
		id: id.String(),
	}
}
