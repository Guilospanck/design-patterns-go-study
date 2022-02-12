package main

func main() {
	// Data
	data := Data{
		Name:    "Guilherme",
		Surname: "Any",
		Age:     25,
	}

	// Create receivers
	receiver := NewReceiver()

	// create commands and associate them with receivers if needed
	sendDataCommand := NewSendDataCommand(data, receiver)

	// create senders
	sender := NewSender(sendDataCommand)

	// actually sends data
	sender.SendData()

}
