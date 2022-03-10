package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var id = "guid"

type CustomerSubscriberSuite struct {
	suite.Suite
	subscriber *CustomerSubscriber
}

func (s *CustomerSubscriberSuite) SetupSuite() {
	s.subscriber = NewCustomerSubscriber(id)
}

func (s *CustomerSubscriberSuite) AfterTests(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestCustomerSubscriberSuiteInit(t *testing.T) {
	suite.Run(t, new(CustomerSubscriberSuite))
}

func (s *CustomerSubscriberSuite) TestUpdate() {
	// arrange
	data := "I'm data"

	// act
	s.subscriber.Update(data)

	// assert
	require.NoError(s.T(), nil)
}

func (s *CustomerSubscriberSuite) TestGetID() {
	// arrange

	// act
	result := s.subscriber.GetID()

	// assert
	require.Equal(s.T(), id, result)
}
