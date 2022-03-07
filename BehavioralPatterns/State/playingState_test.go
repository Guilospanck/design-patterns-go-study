package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PlayingStateSuite struct {
	suite.Suite

	mock    *PlayingStateMock
	state   *PlayingState
	context IPlayerContext
}

func (s *PlayingStateSuite) SetupSuite() {
	NextSongClicked = false
	PreviousSongClicked = false
	StopSongClicked = false

	s.context = NewPlayerContextMock()
	s.mock = NewPlayingStateMock(s.context)
	s.state = NewPlayingState(s.context)
}

func (s *PlayingStateSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestPlayingStateSuiteInit(t *testing.T) {
	suite.Run(t, new(PlayingStateSuite))
}

func (s *PlayingStateSuite) TestClickLock() {
	// arrange

	// act
	s.state.ClickLock()

	// assert
	require.Equal(s.T(), "*main.LockedState", fmt.Sprintf("%T", s.state.context.GetState()))
}

func (s *PlayingStateSuite) TestClickPlay() {
	// arrange
	s.context.SetIsPlaying(true)

	// act
	s.state.ClickPlay()

	// assert
	require.Equal(s.T(), "*main.ReadyState", fmt.Sprintf("%T", s.state.context.GetState()))
	require.Equal(s.T(), true, StopSongClicked)
}

func (s *PlayingStateSuite) TestClickPlayWhenIsNotPlaying() {
	// arrange
	s.context.SetIsPlaying(false)

	// act
	s.state.ClickPlay()

	// assert
	require.Equal(s.T(), "*main.ReadyStateMock", fmt.Sprintf("%T", s.state.context.GetState()))
	require.Equal(s.T(), false, StopSongClicked)
	require.Equal(s.T(), true, s.context.GetIsPlaying())
}

func (s *PlayingStateSuite) TestClickNext() {
	// arrange

	// act
	s.state.ClickNext()

	// assert
	require.Equal(s.T(), "*main.ReadyStateMock", fmt.Sprintf("%T", s.state.context.GetState()))
	require.Equal(s.T(), true, NextSongClicked)
}

func (s *PlayingStateSuite) TestClickPrevious() {
	// arrange

	// act
	s.state.ClickPrevious()

	// assert
	require.Equal(s.T(), "*main.ReadyStateMock", fmt.Sprintf("%T", s.state.context.GetState()))
	require.Equal(s.T(), true, PreviousSongClicked)
}
