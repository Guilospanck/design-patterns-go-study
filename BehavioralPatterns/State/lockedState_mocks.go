package main

import "fmt"

type LockedStateMock struct {
	*State
}

func (s *LockedStateMock) ClickLock() {
	if s.context.GetIsPlaying() {
		s.context.ChangeState(NewPlayingStateMock(s.context))
	} else {
		s.context.ChangeState(NewReadyStateMock(s.context))
	}
}

func (s *LockedStateMock) ClickPlay() {
	// locked
	fmt.Println("Doing nothing...")
}

func (s *LockedStateMock) ClickNext() {
	// locked
	fmt.Println("Doing nothing...")
}

func (s *LockedStateMock) ClickPrevious() {
	// locked
	fmt.Println("Doing nothing...")
}

func NewLockedStateMock(context IPlayerContext) *LockedStateMock {
	return &LockedStateMock{
		State: NewState(context),
	}
}
