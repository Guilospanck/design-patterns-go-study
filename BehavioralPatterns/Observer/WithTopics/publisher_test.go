package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	topic1 = "topic1"
	topic2 = "topic2"

	uuid1 = uuid.New()
	uuid2 = uuid.New()
	uuid3 = uuid.New()
)

type PublisherSuite struct {
	suite.Suite
	pub  *Publisher
	sub1 *SubscriberMock
	sub2 *SubscriberMock
	sub3 *SubscriberMock
}

func (s *PublisherSuite) SetupSuite() {
	s.pub = NewPublisher()
	s.sub1 = NewSubscriberMock(uuid1)
	s.sub2 = NewSubscriberMock(uuid2)
	s.sub3 = NewSubscriberMock(uuid3)
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
	s.pub.AddSubscriber(topic1, s.sub1)
	require.Equal(s.T(), 1, len(s.pub.subscribers[topic1]))
	require.Equal(s.T(), 1, len(s.pub.subscribers))

	s.pub.AddSubscriber(topic1, s.sub2)
	require.Equal(s.T(), 2, len(s.pub.subscribers[topic1]))
	require.Equal(s.T(), 1, len(s.pub.subscribers))

	s.pub.AddSubscriber(topic2, s.sub2)
	require.Equal(s.T(), 2, len(s.pub.subscribers[topic1]))
	require.Equal(s.T(), 1, len(s.pub.subscribers[topic2]))
	require.Equal(s.T(), 2, len(s.pub.subscribers))
}

func (s *PublisherSuite) TestRemoveSubscriber() {
	// arrange
	s.pub.AddSubscriber(topic1, s.sub1)
	s.pub.AddSubscriber(topic1, s.sub2)
	s.pub.AddSubscriber(topic2, s.sub2)

	// act and assert
	s.pub.RemoveSubscriber(topic1, s.sub1)
	require.Equal(s.T(), 1, len(s.pub.subscribers[topic1]))
	require.Equal(s.T(), 1, len(s.pub.subscribers[topic2]))
	require.Equal(s.T(), 2, len(s.pub.subscribers))

	s.pub.RemoveSubscriber(topic1, s.sub3)
	require.Equal(s.T(), 1, len(s.pub.subscribers[topic1]))
	require.Equal(s.T(), 1, len(s.pub.subscribers[topic2]))
	require.Equal(s.T(), 2, len(s.pub.subscribers))

	s.pub.RemoveSubscriber(topic2, s.sub2)
	require.Equal(s.T(), 1, len(s.pub.subscribers[topic1]))
	require.Equal(s.T(), 0, len(s.pub.subscribers[topic2]))
	require.Equal(s.T(), 2, len(s.pub.subscribers))
}

func (s *PublisherSuite) TestNotify() {
	// arrange
	s.pub.AddSubscriber(topic1, s.sub1)
	s.pub.AddSubscriber(topic1, s.sub2)
	s.pub.AddSubscriber(topic2, s.sub3)

	// act and assert
	s.pub.Notify(topic1, s.sub2)
	require.Equal(s.T(), false, s.sub1.updateCalled)
	require.Equal(s.T(), true, s.sub2.updateCalled)
	require.Equal(s.T(), false, s.sub3.updateCalled)

	s.pub.Notify(topic1, s.sub1)
	require.Equal(s.T(), true, s.sub1.updateCalled)
	require.Equal(s.T(), true, s.sub2.updateCalled)
	require.Equal(s.T(), false, s.sub3.updateCalled)

	s.pub.Notify(topic2, s.sub3)
	require.Equal(s.T(), true, s.sub1.updateCalled)
	require.Equal(s.T(), true, s.sub2.updateCalled)
	require.Equal(s.T(), true, s.sub3.updateCalled)
}

func (s *PublisherSuite) TestNotifyAllFromTopic() {
	// arrange
	s.pub.AddSubscriber(topic1, s.sub1)
	s.pub.AddSubscriber(topic1, s.sub2)
	s.pub.AddSubscriber(topic2, s.sub3)

	// act and assert
	s.pub.NotifyAllFromTopic(topic1)
	require.Equal(s.T(), true, s.sub1.updateCalled)
	require.Equal(s.T(), true, s.sub2.updateCalled)
	require.Equal(s.T(), false, s.sub3.updateCalled)

	s.pub.NotifyAllFromTopic(topic2)
	require.Equal(s.T(), true, s.sub1.updateCalled)
	require.Equal(s.T(), true, s.sub2.updateCalled)
	require.Equal(s.T(), true, s.sub3.updateCalled)
}

func (s *PublisherSuite) TestNotifyAll() {
	// arrange
	s.pub.AddSubscriber(topic1, s.sub1)
	s.pub.AddSubscriber(topic1, s.sub2)
	s.pub.AddSubscriber(topic2, s.sub3)

	// act and assert
	s.pub.NotifyAll()
	require.Equal(s.T(), true, s.sub1.updateCalled)
	require.Equal(s.T(), true, s.sub2.updateCalled)
	require.Equal(s.T(), true, s.sub3.updateCalled)
}

func (s *PublisherSuite) TestUpdateState() {
	// arrange

	// act and assert
	require.Equal(s.T(), 0, s.pub.state)

	s.pub.UpdateState()
	require.Equal(s.T(), 1, s.pub.state)
	s.pub.UpdateState()
	require.Equal(s.T(), 2, s.pub.state)
	s.pub.UpdateState()
	require.Equal(s.T(), 3, s.pub.state)
}
