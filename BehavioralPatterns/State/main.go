package main

func main() {

	// Context
	playerContext := NewPlayerContext() // Ready State
	playerContext.ClickLock()           // goes to locked state

	// Locked State
	playerContext.ClickNext()     // do nothing
	playerContext.ClickPlay()     // do nothing
	playerContext.ClickPrevious() // do nothing
	playerContext.ClickLock()     // go to ready state

	// ready state
	playerContext.ClickNext()     // goes to next song
	playerContext.ClickPrevious() // goes to previous song
	playerContext.ClickPlay()     // goes to playing state

	// playing state
	playerContext.ClickNext()     // goes to next song
	playerContext.ClickPrevious() // goes to previous song
	playerContext.ClickPlay()     // stops song (because isPlaying is true) and goes to ready state

	// ready state
	playerContext.ClickLock() // locks it

	// locked state

}
