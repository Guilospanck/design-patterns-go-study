package visitor

import (
	mocks "base/BehavioralPatterns/mocks/visitor"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	mock    *mocks.CircleMock
	circle  *Circle
	visitor IVisitor
}

func (s *Suite) SetupSuite() {
	s.mock = mocks.NewCircleMock()
	s.circle = NewCircle()
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestMove() {
	// arrange
	x := s.mock.X
	y := s.mock.Y

	// act
	s.circle.Move(x, y)

	// assert
	require.Equal(s.T(), s.circle.x, x)
	require.Equal(s.T(), s.circle.y, y)
}

func (s *Suite) TestAccept() {
	// arrange
	s.visitor = mocks.NewExportXMLVisitorMock()

	// act
	s.circle.Accept(s.visitor)

	// assert
	require.Equal(s.T(), mocks.VisitCircleCalled, true)
}
