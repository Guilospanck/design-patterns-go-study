package visitor

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DotSuite struct {
	suite.Suite

	mock    *DotMock
	dot     *Dot
	visitor IVisitor
}

func (s *DotSuite) SetupSuite() {
	s.mock = NewDotMock()
	s.dot = NewDot()
}

func (s *DotSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestDotInit(t *testing.T) {
	suite.Run(t, new(DotSuite))
}

func (s *DotSuite) TestRectangleMove() {
	// arrange
	x := s.mock.X
	y := s.mock.Y

	// act
	s.dot.Move(x, y)

	// assert
	require.Equal(s.T(), s.dot.x, x)
	require.Equal(s.T(), s.dot.y, y)
}

func (s *DotSuite) TestRectangleAccept() {
	// arrange
	s.visitor = NewExportXMLVisitorMock()

	// act
	s.dot.Accept(s.visitor)

	// assert
	require.Equal(s.T(), VisitDotCalled, true)
}
