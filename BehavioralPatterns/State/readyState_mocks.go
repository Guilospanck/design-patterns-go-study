package main

type ReadyStateMock struct {
	*State
}

func (s *ReadyStateMock) ClickLock() {
	s.context.ChangeState(NewLockedStateMock(s.context))
}

func (s *ReadyStateMock) ClickPlay() {
	s.context.SetIsPlaying(true)
	s.context.ChangeState(NewPlayingStateMock(s.context))
}

func (s *ReadyStateMock) ClickNext() {
	s.context.NextSong()
}

func (s *ReadyStateMock) ClickPrevious() {
	s.context.PreviousSong()
}

func NewReadyStateMock(context IPlayerContext) *ReadyStateMock {
	return &ReadyStateMock{
		State: NewState(context),
	}
}
