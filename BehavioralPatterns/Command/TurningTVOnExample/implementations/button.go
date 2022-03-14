package implementations

import "base/BehavioralPatterns/Command/TurningTVOnExample/interfaces"

// Invoker (Sender)
type Button struct {
	command interfaces.ICommand
}

func (b *Button) Press() {
	b.command.Execute()
}

func NewButton(command interfaces.ICommand) *Button {
	return &Button{
		command: command,
	}
}
