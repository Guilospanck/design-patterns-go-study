package main

import (
	"github.com/google/uuid"
)

type SubscriberMock struct {
	id           string
	updateCalled bool
}

func (s *SubscriberMock) Update(data IPublisher) {
	s.updateCalled = true
}

func (s *SubscriberMock) GetID() string {
	return s.id
}

func NewSubscriberMock(optionalId ...uuid.UUID) *SubscriberMock {
	id := uuid.New()
	if len(optionalId) > 0 {
		id = optionalId[0]
	}

	return &SubscriberMock{
		id:           id.String(),
		updateCalled: false,
	}
}
