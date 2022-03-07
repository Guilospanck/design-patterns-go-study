package main

type PlayingStateMock struct {
	*State
}

func (s *PlayingStateMock) ClickLock() {
	s.context.ChangeState(NewLockedStateMock(s.context))
}

func (s *PlayingStateMock) ClickPlay() {
	if s.context.GetIsPlaying() {
		s.context.StopSong()
		s.context.ChangeState(NewReadyStateMock(s.context))
	} else {
		s.context.SetIsPlaying(true)
	}
}

func (s *PlayingStateMock) ClickNext() {
	s.context.NextSong()
}

func (s *PlayingStateMock) ClickPrevious() {
	s.context.PreviousSong()
}

func NewPlayingStateMock(context IPlayerContext) *PlayingStateMock {
	return &PlayingStateMock{
		State: NewState(context),
	}
}
