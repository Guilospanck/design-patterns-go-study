package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Receiver_ReceiveData(t *testing.T) {
	t.Run("it should receive data", func(t *testing.T) {
		// arrange
		sut := makeReceiverSut()

		// act
		sut.usecase.ReceiveData(sut.data)

		// assert
		assert.NoError(t, nil)
	})
}

type receiverSut struct {
	usecase *Receiver
	data    Data
}

func makeReceiverSut() receiverSut {
	usecase := NewReceiver()

	data := Data{
		Name:    "Guilherme",
		Surname: "Any",
		Age:     26,
	}

	return receiverSut{usecase, data}
}
