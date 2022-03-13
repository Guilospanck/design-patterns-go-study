package implementations

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserIteratorSuite struct {
	suite.Suite

	iterator  *UserIterator
	userMocks []*User
}

func (s *UserIteratorSuite) SetupSuite() {
	s.userMocks = []*User{UserMock1, UserMock2, UserMock3}

	s.iterator = &UserIterator{users: s.userMocks}
}

func (s *UserIteratorSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestUserIteratorSuiteInit(t *testing.T) {
	suite.Run(t, new(UserIteratorSuite))
}

func (s *UserIteratorSuite) TestHasNext() {
	// act
	result := s.iterator.HasNext()

	// assert
	require.Equal(s.T(), true, result)
}

func (s *UserIteratorSuite) TestGetNext() {
	// arrange
	userByte, _ := json.Marshal(UserMock1)
	userStr := string(userByte[:])

	// act
	result := s.iterator.GetNext()

	// act
	require.Equal(s.T(), userStr, result)
	require.Equal(s.T(), 1, s.iterator.index)
}
