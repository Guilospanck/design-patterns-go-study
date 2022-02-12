package main

import "fmt"

type Sender struct {
	sendDataCommand ICommand
}

func (s *Sender) SendData() {
	fmt.Printf("Sender sending data...\n")
	s.sendDataCommand.Execute()
}

func NewSender(sendDataCommand ICommand) *Sender {
	return &Sender{
		sendDataCommand: sendDataCommand,
	}
}
