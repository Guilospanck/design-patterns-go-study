package implementations

import "base/BehavioralPatterns/Command/TurningTVOnExample/interfaces"

type OnCommand struct {
	device interfaces.IDevice
}

func (command *OnCommand) Execute() {
	command.device.On()
}

func NewOnCommand(device interfaces.IDevice) *OnCommand {
	return &OnCommand{
		device: device,
	}
}
