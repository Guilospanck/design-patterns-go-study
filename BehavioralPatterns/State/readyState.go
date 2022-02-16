package main

type ReadyState struct {
	*State
}

func (s *ReadyState) ClickLock() {
	s.context.ChangeState(NewLockedState(s.context))
}

func (s *ReadyState) ClickPlay() {
	s.context.IsPlaying = true
	s.context.ChangeState(NewPlayingState(s.context))
}

func (s *ReadyState) ClickNext() {
	s.context.NextSong()
}

func (s *ReadyState) ClickPrevious() {
	s.context.PreviousSong()
}

func NewReadyState(context *PlayerContext) *ReadyState {
	return &ReadyState{
		State: NewState(context),
	}
}
