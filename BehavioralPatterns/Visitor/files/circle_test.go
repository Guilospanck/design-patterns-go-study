package visitor

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CircleSuite struct {
	suite.Suite

	mock    *CircleMock
	circle  *Circle
	visitor IVisitor
}

func (s *CircleSuite) SetupSuite() {
	s.mock = NewCircleMock()
	s.circle = NewCircle()
}

func (s *CircleSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestCircleInit(t *testing.T) {
	suite.Run(t, new(CircleSuite))
}

func (s *CircleSuite) TestMove() {
	// arrange
	x := s.mock.X
	y := s.mock.Y

	// act
	s.circle.Move(x, y)

	// assert
	require.Equal(s.T(), s.circle.x, x)
	require.Equal(s.T(), s.circle.y, y)
}

func (s *CircleSuite) TestAccept() {
	// arrange
	s.visitor = NewExportXMLVisitorMock()

	// act
	s.circle.Accept(s.visitor)

	// assert
	require.Equal(s.T(), VisitCircleCalled, true)
}
