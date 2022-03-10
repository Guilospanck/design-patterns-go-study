package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SubscriberSuite struct {
	suite.Suite
	sub        *Subscriber
	mockedUUID uuid.UUID
	pub        *Publisher
}

func (s *SubscriberSuite) SetupSuite() {
	s.mockedUUID = uuid.New()
	s.sub = NewSubscriber(s.mockedUUID)
	s.pub = NewPublisher()
}

func (s *SubscriberSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestSubscriberSuiteInit(t *testing.T) {
	suite.Run(t, new(SubscriberSuite))
}

func (s *SubscriberSuite) TestUpdate() {
	// arrange

	// act
	s.sub.Update(s.pub)

	// assert
	require.NoError(s.T(), nil)
}

func (s *SubscriberSuite) TestGetID() {
	// arrange

	// act
	result := s.sub.GetID()

	// assert
	require.Equal(s.T(), s.mockedUUID.String(), result)
}
