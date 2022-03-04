package main

import "fmt"

// Context
type PlayerContext struct {
	IsPlaying bool
	State     IState
}

func (context *PlayerContext) GetState() IState {
	return context.State
}

func (context *PlayerContext) ChangeState(state IState) {
	context.State = state
}

func (context *PlayerContext) ClickLock() {
	fmt.Printf("[CLICK LOCK] => State: %T\n", context.State)
	context.State.ClickLock()
}

func (context *PlayerContext) ClickPlay() {
	context.State.ClickPlay()
}

func (context *PlayerContext) ClickNext() {
	context.State.ClickNext()
}

func (context *PlayerContext) ClickPrevious() {
	context.State.ClickPrevious()
}

// State may call some methods in the context
func (context *PlayerContext) NextSong() {
	fmt.Println("Next song...")
}

func (context *PlayerContext) PreviousSong() {
	fmt.Println("Previous song...")
}

func (context *PlayerContext) StopSong() {
	fmt.Println("Stop song...")
}

func (context *PlayerContext) GetIsPlaying() bool {
	return context.IsPlaying
}

func (context *PlayerContext) SetIsPlaying(isPlaying bool) {
	context.IsPlaying = isPlaying
}

func NewPlayerContext() *PlayerContext {
	playerContext := &PlayerContext{}
	readyState := NewReadyState(playerContext)
	playerContext.State = readyState
	return playerContext
}
