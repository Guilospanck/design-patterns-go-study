package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Client_TransferData(t *testing.T) {
	t.Run("Should get data from lightning port properly and send data, returning true", func(t *testing.T) {
		// arrange
		sut := makeClientSut()
		sut.lightningPortSpy.On("GetData").Return(sut.data)

		// act
		result := sut.client.TransferData(sut.lightningPortSpy)

		// assert
		assert.True(t, result)
		sut.lightningPortSpy.AssertExpectations(t)
		sut.lightningPortSpy.AssertCalled(t, "GetData")
		sut.lightningPortSpy.AssertNumberOfCalls(t, "GetData", 1)

	})
}

type clientSut struct {
	data             float64
	lightningPortSpy *lightningPortSpy
	client           *client
}

func makeClientSut() clientSut {
	data := 45.5
	lightningPortSpy := NewLightningPortSpy(data)

	client := NewClient()

	return clientSut{data, lightningPortSpy, client}
}
