package main

import "testing"

func Test_Sender_SendData(t *testing.T) {
	t.Run("it should send data", func(t *testing.T) {
		// arrange
		sut := makeSenderSut()
		sut.sendDataCommand.On("Execute").Return()

		// act
		sut.usecase.SendData()

		// assert
		sut.sendDataCommand.AssertExpectations(t)
		sut.sendDataCommand.AssertNumberOfCalls(t, "Execute", 1)
	})
}

type senderSut struct {
	usecase         *Sender
	sendDataCommand *SendDataCommandSpy
}

func makeSenderSut() senderSut {
	sendDataCommand := NewSendDataCommandSpy()
	usecase := NewSender(sendDataCommand)

	return senderSut{usecase, sendDataCommand}
}
