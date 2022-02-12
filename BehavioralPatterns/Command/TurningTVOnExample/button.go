package main

// Invoker (Sender)
type Button struct {
	command ICommand
}

func (b *Button) Press() {
	b.command.Execute()
}

func NewButton(command ICommand) *Button {
	return &Button{
		command: command,
	}
}
