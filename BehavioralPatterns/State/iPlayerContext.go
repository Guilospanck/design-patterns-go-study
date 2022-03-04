package main

type IPlayerContext interface {
	GetState() IState
	ChangeState(state IState)
	NextSong()
	PreviousSong()
	StopSong()
	GetIsPlaying() bool
	SetIsPlaying(bool)
}
