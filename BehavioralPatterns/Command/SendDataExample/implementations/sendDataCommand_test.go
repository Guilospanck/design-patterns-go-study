package implementations

import (
	"base/BehavioralPatterns/Command/SendDataExample/domain"
	"testing"
)

func Test_SendDataCommand_Execute(t *testing.T) {
	t.Run("it should Execute method in SendDataCommand", func(t *testing.T) {
		// arrange
		sut := makeSendDataCommandSut()
		sut.receiver.On("ReceiveData", sut.data).Return()

		// act
		sut.usecase.Execute()

		// assert
		sut.receiver.AssertNumberOfCalls(t, "ReceiveData", 1)
		sut.receiver.AssertExpectations(t)

	})
}

type sendDataCommandSut struct {
	data     domain.Data
	receiver *ReceiverSpy
	usecase  *SendDataCommand
}

func makeSendDataCommandSut() sendDataCommandSut {
	data := domain.Data{
		Name:    "Guilherme",
		Surname: "Any",
		Age:     26,
	}

	receiver := NewReceiverSpy()

	usecase := NewSendDataCommand(data, receiver)

	return sendDataCommandSut{data, receiver, usecase}
}
