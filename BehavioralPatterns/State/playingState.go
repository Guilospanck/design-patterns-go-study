package main

type PlayingState struct {
	*State
}

func (s *PlayingState) ClickLock() {
	s.context.ChangeState(NewLockedState(s.context))
}

func (s *PlayingState) ClickPlay() {
	if s.context.GetIsPlaying() {
		s.context.StopSong()
		s.context.ChangeState(NewReadyState(s.context))
	} else {
		s.context.SetIsPlaying(true)
	}
}

func (s *PlayingState) ClickNext() {
	s.context.NextSong()
}

func (s *PlayingState) ClickPrevious() {
	s.context.PreviousSong()
}

func NewPlayingState(context IPlayerContext) *PlayingState {
	return &PlayingState{
		State: NewState(context),
	}
}
