package main

import "fmt"

// Concrete receiver
type TV struct {
	isRunning bool
}

func (tv *TV) On() {
	tv.isRunning = true
	fmt.Println("TV is turning ON")
}

func (tv *TV) Off() {
	tv.isRunning = false
	fmt.Println("TV is turning OFF")
}

func NewTV() *TV {
	return &TV{}
}
