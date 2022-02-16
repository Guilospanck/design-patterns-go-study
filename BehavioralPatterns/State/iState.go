package main

type IState interface {
	ClickLock()
	ClickPlay()
	ClickNext()
	ClickPrevious()
}

type State struct {
	context *PlayerContext
}

func NewState(context *PlayerContext) *State {
	return &State{
		context: context,
	}
}
