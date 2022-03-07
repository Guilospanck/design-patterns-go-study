package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LockedStateSuite struct {
	suite.Suite

	contextMock IPlayerContext
	mock        *LockedStateMock
	state       *LockedState
}

func (s *LockedStateSuite) SetupSuite() {
	s.contextMock = NewPlayerContextMock()
	s.mock = NewLockedStateMock(s.contextMock)
	s.state = NewLockedState(s.contextMock)
}

func (s *LockedStateSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestLockedStateInit(t *testing.T) {
	suite.Run(t, new(LockedStateSuite))
}

func (s *LockedStateSuite) TestClickLock() {
	// arrange

	// act
	s.state.ClickLock()

	// assert
	require.Equal(s.T(), "*main.ReadyState", fmt.Sprintf("%T", s.state.context.GetState()))
}

func (s *LockedStateSuite) TestClickPlay() {
	// arrange

	// act
	s.state.ClickPlay()

	// assert
	require.Equal(s.T(), "*main.ReadyStateMock", fmt.Sprintf("%T", s.state.context.GetState()))
}

func (s *LockedStateSuite) TestClickNext() {
	// arrange

	// act
	s.state.ClickNext()

	// assert
	require.Equal(s.T(), "*main.ReadyStateMock", fmt.Sprintf("%T", s.state.context.GetState()))
}

func (s *LockedStateSuite) TestClickPrevious() {
	// arrange

	// act
	s.state.ClickPrevious()

	// assert
	require.Equal(s.T(), "*main.ReadyStateMock", fmt.Sprintf("%T", s.state.context.GetState()))
}
