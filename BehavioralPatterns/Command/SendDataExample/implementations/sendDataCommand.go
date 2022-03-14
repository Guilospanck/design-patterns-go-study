package implementations

import (
	"base/BehavioralPatterns/Command/SendDataExample/domain"
	"base/BehavioralPatterns/Command/SendDataExample/interfaces"
	"fmt"
)

type SendDataCommand struct {
	data     domain.Data
	receiver interfaces.IReceiver
}

func (command *SendDataCommand) Execute() {
	fmt.Printf("Command sending data to receiver...\n")
	command.receiver.ReceiveData(command.data)
}

func NewSendDataCommand(data domain.Data, receiver interfaces.IReceiver) *SendDataCommand {
	return &SendDataCommand{
		data:     data,
		receiver: receiver,
	}
}
