package implementations

import (
	"base/BehavioralPatterns/Command/SendDataExample/interfaces"
	"fmt"
)

type Sender struct {
	sendDataCommand interfaces.ICommand
}

func (s *Sender) SendData() {
	fmt.Printf("Sender sending data...\n")
	s.sendDataCommand.Execute()
}

func NewSender(sendDataCommand interfaces.ICommand) *Sender {
	return &Sender{
		sendDataCommand: sendDataCommand,
	}
}
