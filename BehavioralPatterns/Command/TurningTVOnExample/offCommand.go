package main

type OffCommand struct {
	device IDevice
}

func (command *OffCommand) Execute() {
	command.device.Off()
}

func NewOffCommand(device IDevice) *OffCommand {
	return &OffCommand{
		device: device,
	}
}
