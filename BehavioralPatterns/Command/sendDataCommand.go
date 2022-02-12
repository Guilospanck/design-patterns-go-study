package main

import "fmt"

type Data struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type SendDataCommand struct {
	data     Data
	receiver IReceiver
}

func (command *SendDataCommand) Execute() {
	fmt.Printf("Command sending data to receiver...\n")
	command.receiver.ReceiveData(command.data)
}

func NewSendDataCommand(data Data, receiver IReceiver) *SendDataCommand {
	return &SendDataCommand{
		data:     data,
		receiver: receiver,
	}
}
