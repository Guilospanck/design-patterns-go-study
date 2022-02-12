package main

type OnCommand struct {
	device IDevice
}

func (command *OnCommand) Execute() {
	command.device.On()
}

func NewOnCommand(device IDevice) *OnCommand {
	return &OnCommand{
		device: device,
	}
}
