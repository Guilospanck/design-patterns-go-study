package implementations

import (
	"base/BehavioralPatterns/Iterator/interfaces"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FriendsCollectionSuite struct {
	suite.Suite

	friendsCollection *FriendsCollection
	friendsMocks      []*Friend
}

func (s *FriendsCollectionSuite) SetupSuite() {
	s.friendsMocks = []*Friend{FriendMock1, FriendMock2, FriendMock3}

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
