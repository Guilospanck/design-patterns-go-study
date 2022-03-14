package main

import (
	"base/BehavioralPatterns/Command/SendDataExample/domain"
	"base/BehavioralPatterns/Command/SendDataExample/implementations"
)

func main() {
	// Data
	data := domain.Data{
		Name:    "Guilherme",
		Surname: "Any",
		Age:     25,
	}

	// Create receivers
	receiver := implementations.NewReceiver()

	// create commands and associate them with receivers if needed
	sendDataCommand := implementations.NewSendDataCommand(data, receiver)

	// create senders
	sender := implementations.NewSender(sendDataCommand)

	// actually sends data
	sender.SendData()

}
