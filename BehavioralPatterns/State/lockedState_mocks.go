package main

import "fmt"

type LockedStateMock struct {
	*State
}

func (s *LockedStateMock) ClickLock() {
	if s.context.IsPlaying {
		s.context.ChangeState(NewPlayingState(s.context))
	} else {
		s.context.ChangeState(NewReadyState(s.context))
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

func NewLockedStateMock(context *PlayerContext) *LockedStateMock {
	return &LockedStateMock{
		State: NewState(context),
	}
}
