package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Spy_LightningPort_GetData(t *testing.T) {
	t.Run("Should get data properly", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.lightningPortSpy.On("GetData").Return(sut.data)

		// act
		result := sut.lightningPortSpy.GetData()

		// assert
		assert.Equal(t, sut.data, result)
		sut.lightningPortSpy.AssertExpectations(t)
		sut.lightningPortSpy.AssertCalled(t, "GetData")
		sut.lightningPortSpy.AssertNumberOfCalls(t, "GetData", 1)
	})
}

type spySut struct {
	data             float64
	lightningPortSpy *lightningPortSpy
}

func makeSpySut() spySut {
	data := 45.5
	lightningPortSpy := NewLightningPortSpy(data)

	return spySut{data, lightningPortSpy}
}
