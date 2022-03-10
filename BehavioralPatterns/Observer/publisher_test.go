package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	sub1 = NewCustomerSubscriberMock("Customer1")
	sub2 = NewCustomerSubscriberMock("Customer2")
	sub3 = NewCustomerSubscriberMock("Customer3")
)

type PublisherSuite struct {
	suite.Suite
	pub      *Publisher
	dataMock *Data
}

func (s *PublisherSuite) SetupSuite() {
	s.dataMock = &Data{
		Name: "Product1",
		Size: 48,
	}
	s.pub = NewPublisher(s.dataMock.Marshal())
}

func (s *PublisherSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestPublisherSuiteInit(t *testing.T) {
	suite.Run(t, new(PublisherSuite))
}

func (s *PublisherSuite) TestAddSubscriber() {
	// arrange

	// act and assert
	require.Equal(s.T(), 0, len(s.pub.subscribers))

	s.pub.AddSubscriber(sub1)
	require.Equal(s.T(), 1, len(s.pub.subscribers))
	require.Equal(s.T(), sub1.GetID(), s.pub.subscribers[0].GetID())

	s.pub.AddSubscriber(sub2)
	require.Equal(s.T(), 2, len(s.pub.subscribers))
	require.Equal(s.T(), sub2.GetID(), s.pub.subscribers[1].GetID())

	s.pub.AddSubscriber(sub3)
	require.Equal(s.T(), 3, len(s.pub.subscribers))
	require.Equal(s.T(), sub3.GetID(), s.pub.subscribers[2].GetID())
}

func (s *PublisherSuite) TestRemoveSubscriber() {
	// arrange
	s.pub.AddSubscriber(sub1)
	s.pub.AddSubscriber(sub2)

	// act and assert
	s.pub.RemoveSubscriber(sub1)
	require.Equal(s.T(), 1, len(s.pub.subscribers))
	require.Equal(s.T(), sub2.GetID(), s.pub.subscribers[0].GetID())

	s.pub.RemoveSubscriber(sub2)
	require.Equal(s.T(), 0, len(s.pub.subscribers))

	s.pub.RemoveSubscriber(sub3)
	require.Equal(s.T(), 0, len(s.pub.subscribers))
}

func (s *PublisherSuite) TestNotifyAll() {
	// arrange
	s.pub.AddSubscriber(sub1)
	s.pub.AddSubscriber(sub2)

	// act
	s.pub.NotifyAll()

	// assert
	require.Equal(s.T(), true, sub1.updateCalled)
	require.Equal(s.T(), true, sub2.updateCalled)
	require.Equal(s.T(), false, sub3.updateCalled)
}

func (s *PublisherSuite) TestUpdateAvailabilityOfData() {
	// arrange
	s.pub.AddSubscriber(sub1)
	s.pub.AddSubscriber(sub2)

	// act
	s.pub.updateAvailabilityOfData()

	// assert
	require.Equal(s.T(), true, sub1.updateCalled)
	require.Equal(s.T(), true, sub2.updateCalled)
	require.Equal(s.T(), false, sub3.updateCalled)
}
