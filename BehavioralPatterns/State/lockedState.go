package main

import "fmt"

type LockedState struct {
	*State
}

func (s *LockedState) ClickLock() {
	if s.context.GetIsPlaying() {
		s.context.ChangeState(NewPlayingState(s.context))
	} else {
		s.context.ChangeState(NewReadyState(s.context))
	}
}

func (s *LockedState) ClickPlay() {
	// locked
	fmt.Println("Doing nothing...")
}

func (s *LockedState) ClickNext() {
	// locked
	fmt.Println("Doing nothing...")
}

func (s *LockedState) ClickPrevious() {
	// locked
	fmt.Println("Doing nothing...")
}

func NewLockedState(context IPlayerContext) *LockedState {
	return &LockedState{
		State: NewState(context),
	}
}
