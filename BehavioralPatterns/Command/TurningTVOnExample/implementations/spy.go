package implementations

import "github.com/stretchr/testify/mock"

type CommandSpy struct {
	mock.Mock
}

func (spy *CommandSpy) Execute() {
	spy.Called()
}

func NewCommandSpy() *CommandSpy {
	return new(CommandSpy)
}

//
type DeviceSpy struct {
	mock.Mock
}

func (spy *DeviceSpy) TurnOn() {
	spy.Called()
}

func (spy *DeviceSpy) TurnOff() {
	spy.Called()
}

func NewDeviceSpy() *DeviceSpy {
	return new(DeviceSpy)
}
