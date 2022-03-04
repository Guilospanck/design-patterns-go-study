package main

type IState interface {
	ClickLock()
	ClickPlay()
	ClickNext()
	ClickPrevious()
}

type State struct {
	context IPlayerContext
}

func NewState(context IPlayerContext) *State {
	return &State{
		context: context,
	}
}
