package implementations

import "base/BehavioralPatterns/Command/TurningTVOnExample/interfaces"

type OffCommand struct {
	device interfaces.IDevice
}

func (command *OffCommand) Execute() {
	command.device.Off()
}

func NewOffCommand(device interfaces.IDevice) *OffCommand {
	return &OffCommand{
		device: device,
	}
}
