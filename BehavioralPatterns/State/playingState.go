package main

type PlayingState struct {
	*State
}

func (s *PlayingState) ClickLock() {
	s.context.ChangeState(NewLockedState(s.context))
}

func (s *PlayingState) ClickPlay() {
	if s.context.IsPlaying {
		s.context.StopSong()
		s.context.ChangeState(NewReadyState(s.context))
	} else {
		s.context.IsPlaying = true
	}
}

func (s *PlayingState) ClickNext() {
	s.context.NextSong()
}

func (s *PlayingState) ClickPrevious() {
	s.context.PreviousSong()
}

func NewPlayingState(context *PlayerContext) *PlayingState {
	return &PlayingState{
		State: NewState(context),
	}
}
