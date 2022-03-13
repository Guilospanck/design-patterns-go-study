package implementations

import (
	"base/BehavioralPatterns/Iterator/interfaces"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserCollectionSuite struct {
	suite.Suite

	collection *UserCollection
	userMocks  []*User
}

func (s *UserCollectionSuite) SetupSuite() {
	s.userMocks = []*User{UserMock1, UserMock2, UserMock3}
	s.collection = NewUserCollection(s.userMocks)
}

func (s *UserCollectionSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestUserCollectionInit(t *testing.T) {
	suite.Run(t, new(UserCollectionSuite))
}

func (s *UserCollectionSuite) TestCreateIterator() {
	// arrange

	// act
	result := s.collection.CreateIterator()

	// assert
	require.Implements(s.T(), (*interfaces.IIterator)(nil), result)
}
