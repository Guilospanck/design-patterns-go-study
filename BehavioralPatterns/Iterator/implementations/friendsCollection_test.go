package implementations

import (
	"base/BehavioralPatterns/Iterator/interfaces"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FriendsCollectionSuite struct {
	suite.Suite

	friendsIteratorMock *FriendsIteratorMocks
	friendsCollection   *FriendsCollection
	friendsMocks        []*Friend
}

func (s *FriendsCollectionSuite) SetupSuite() {
	friend1 := NewFriend("Friend1", "Neighborhood1", "City1", 1)
	friend2 := NewFriend("Friend2", "Neighborhood2", "City2", 2)
	friend3 := NewFriend("Friend3", "Neighborhood3", "City3", 3)
	s.friendsMocks = []*Friend{friend1, friend2, friend3}

	s.friendsIteratorMock = NewFriendsIteratorMocks(s.friendsMocks)
	s.friendsCollection = NewFriendsCollection(s.friendsMocks)
}

func (s *FriendsCollectionSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestFriendsCollectionInit(t *testing.T) {
	suite.Run(t, new(FriendsCollectionSuite))
}

func (s *FriendsCollectionSuite) TestCreateIterator() {
	// return NewFriendsIterator(collection.Friends)

	// act
	result := s.friendsCollection.CreateIterator()

	// assert
	require.Implements(s.T(), (*interfaces.IIterator)(nil), result)
}

func (s *FriendsCollectionSuite) TestAddFriend() {
	// arrange
	friend := NewFriend("Friend4", "Neighborhood4", "City4", 4)

	// act
	s.friendsCollection.AddFriend(*friend)

	// assert
	require.Equal(s.T(), 4, len(s.friendsCollection.Friends))
	require.Equal(s.T(), friend, s.friendsCollection.Friends[3])
}