package implementations

import "base/BehavioralPatterns/Command/TurningTVOnExample/interfaces"

type OffCommand struct {
	device interfaces.IDevice
}

func (command *OffCommand) Execute() {
	command.device.TurnOff()
}

func NewOffCommand(device interfaces.IDevice) *OffCommand {
	return &OffCommand{
		device: device,
	}
}
