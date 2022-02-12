package main

import "fmt"

// Another concrete receiver
type Radio struct {
	isRunning bool
}

func (r *Radio) On() {
	r.isRunning = true
	fmt.Println("Radio is turning on...")
}

func (r *Radio) Off() {
	r.isRunning = false
	fmt.Println("Radio is turning off...")
}

func NewRadio() *Radio {
	return &Radio{}
}
