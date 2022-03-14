package implementations

import "fmt"

// Another concrete receiver
type Radio struct {
	isRunning bool
}

func (r *Radio) TurnOn() {
	r.isRunning = true
	fmt.Println("Radio is turning on...")
}

func (r *Radio) TurnOff() {
	r.isRunning = false
	fmt.Println("Radio is turning off...")
}

func NewRadio() *Radio {
	return &Radio{}
}
