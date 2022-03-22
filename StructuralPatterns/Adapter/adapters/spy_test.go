package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Spy_UsbPort_GetData(t *testing.T) {
	t.Run("Should get data properly", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.usbPortSpy.On("GetData").Return(sut.data)

		// act
		result := sut.usbPortSpy.GetData()

		// assert
		assert.Equal(t, sut.data, result)
		sut.usbPortSpy.AssertExpectations(t)
	})
}

type spySut struct {
	usbPortSpy *usbPortSpy
	data       int
}

func makeSpySut() spySut {
	data := 15
	usbPortSpy := NewUSBPortSpy(data)

	return spySut{usbPortSpy, data}
}
