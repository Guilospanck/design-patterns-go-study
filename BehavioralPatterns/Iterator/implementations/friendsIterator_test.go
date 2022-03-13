package implementations

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FriendsIteratorSuite struct {
	suite.Suite

	iterator     *FriendsIterator
	friendsMocks []*Friend
}

func (s *FriendsIteratorSuite) SetupSuite() {
	s.friendsMocks = []*Friend{FriendMock1, FriendMock2, FriendMock3}

	s.iterator = NewFriendsIterator(s.friendsMocks)
}

func (s *FriendsIteratorSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestFriendsIteratorSuiteInit(t *testing.T) {
	suite.Run(t, new(FriendsIteratorSuite))
}

func (s *FriendsIteratorSuite) TestHasNext() {
	// arrange

	// act
	result := s.iterator.HasNext()

	// assert
	require.Equal(s.T(), true, result)
	require.Equal(s.T(), 0, s.iterator.index)
}

func (s *FriendsIteratorSuite) TestGetNext() {
	// arrange
	friend1MockByte, _ := json.Marshal(FriendMock1)
	friend1MockStr := string(friend1MockByte[:])

	// act
	result := s.iterator.GetNext()

	// assert
	require.Equal(s.T(), friend1MockStr, result)
	require.Equal(s.T(), 1, s.iterator.index)
}
