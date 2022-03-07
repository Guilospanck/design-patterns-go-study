package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ReadyStateSuite struct {
	suite.Suite

	mock    *ReadyStateMock
	state   *ReadyState
	context IPlayerContext
}

func (s *ReadyStateSuite) SetupSuite() {
	s.context = NewPlayerContextMock()
	s.mock = NewReadyStateMock(s.context)
	s.state = NewReadyState(s.context)
}

func (s *ReadyStateSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestReadyStateSuiteInit(t *testing.T) {
	suite.Run(t, new(ReadyStateSuite))
}

func (s *ReadyStateSuite) TestClickLock() {
	// arrange

	// act
	s.state.ClickLock()

	// assert
	require.Equal(s.T(), "*main.LockedState", fmt.Sprintf("%T", s.state.context.GetState()))
}

func (s *ReadyStateSuite) TestClickPlay() {
	// arrange

	// act
	s.state.ClickPlay()

	// assert
	require.Equal(s.T(), "*main.PlayingState", fmt.Sprintf("%T", s.state.context.GetState()))
	require.Equal(s.T(), true, s.state.context.GetIsPlaying())
}

func (s *ReadyStateSuite) TestClickNext() {
	// arrange

	// act
	s.state.ClickNext()

	// assert
	require.Equal(s.T(), true, NextSongClicked)
}

func (s *ReadyStateSuite) TestClickPrevious() {
	// arrange

	// act
	s.state.ClickPrevious()

	// assert
	require.Equal(s.T(), true, PreviousSongClicked)
}
