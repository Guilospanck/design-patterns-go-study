package implementations

import "fmt"

// Concrete receiver
type TV struct {
	isRunning bool
}

func (tv *TV) TurnOn() {
	tv.isRunning = true
	fmt.Println("TV is turning ON")
}

func (tv *TV) TurnOff() {
	tv.isRunning = false
	fmt.Println("TV is turning OFF")
}

func NewTV() *TV {
	return &TV{}
}
