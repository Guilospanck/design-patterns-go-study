package visitor

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RectangleSuite struct {
	suite.Suite

	mock      *RectangleMock
	rectangle *Rectangle
	visitor   IVisitor
}

func (s *RectangleSuite) SetupSuite() {
	s.mock = NewRectangleMock()
	s.rectangle = NewRectangle()
}

func (s *RectangleSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestRectangleInit(t *testing.T) {
	suite.Run(t, new(RectangleSuite))
}

func (s *RectangleSuite) TestMove() {
	// arrange
	x := s.mock.X
	y := s.mock.Y

	// act
	s.rectangle.Move(x, y)

	// assert
	require.Equal(s.T(), s.rectangle.x, x)
	require.Equal(s.T(), s.rectangle.y, y)
}

func (s *RectangleSuite) TestAccept() {
	// arrange
	s.visitor = NewExportXMLVisitorMock()

	// act
	s.rectangle.Accept(s.visitor)

	// assert
	require.Equal(s.T(), VisitRectangleCalled, true)
}
