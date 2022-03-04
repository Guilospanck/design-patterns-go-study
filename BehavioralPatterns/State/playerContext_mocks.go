package main

import "fmt"

// Context
type PlayerContextMock struct {
	IsPlaying bool
	State     IState
}

func (context *PlayerContextMock) GetState() IState {
	return context.State
}

func (context *PlayerContextMock) ChangeState(state IState) {
	context.State = state
}

func (context *PlayerContextMock) ClickLock() {
	fmt.Printf("[CLICK LOCK] => State: %T\n", context.State)
	context.State.ClickLock()
}

func (context *PlayerContextMock) ClickPlay() {
	context.State.ClickPlay()
}

func (context *PlayerContextMock) ClickNext() {
	context.State.ClickNext()
}

func (context *PlayerContextMock) ClickPrevious() {
	context.State.ClickPrevious()
}

// State may call some methods in the context
func (context *PlayerContextMock) NextSong() {
	fmt.Println("Next song...")
}

func (context *PlayerContextMock) PreviousSong() {
	fmt.Println("Previous song...")
}

func (context *PlayerContextMock) StopSong() {
	fmt.Println("Stop song...")
}

func (context *PlayerContextMock) GetIsPlaying() bool {
	return context.IsPlaying
}

func (context *PlayerContextMock) SetIsPlaying(isPlaying bool) {
	context.IsPlaying = isPlaying
}

func NewPlayerContextMock() *PlayerContextMock {
	playerContext := &PlayerContextMock{}
	readyState := NewReadyStateMock(playerContext)
	playerContext.State = readyState
	return playerContext
}
