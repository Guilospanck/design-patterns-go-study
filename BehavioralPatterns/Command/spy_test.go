package main

import "testing"

func Test_Execute(t *testing.T) {
	t.Run("should execute mock correctly", func(t *testing.T) {
		// arrange
		sut := NewSendDataCommandSpy()
		sut.On("Execute").Return()

		// act
		sut.Execute()

		// assert
		sut.AssertExpectations(t)
	})
}

func Test_ReceiveData(t *testing.T) {
	t.Run("should execute mock correctly", func(t *testing.T) {
		// arrange
		sut := NewReceiverSpy()

		data := Data{
			Name:    "Guilherme",
			Surname: "Any",
			Age:     26,
		}

		sut.On("ReceiveData", data).Return()

		// act
		sut.ReceiveData(data)

		// assert
		sut.AssertExpectations(t)
	})
}
