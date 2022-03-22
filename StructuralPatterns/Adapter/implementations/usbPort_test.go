package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_USBPort_GetData(t *testing.T) {
	t.Run("Should return the right data", func(t *testing.T) {
		// arrange
		sut := makeUSBPortSut()

		// act
		result := sut.sut.GetData()

		// assert
		assert.Equal(t, sut.data, result)
	})
}

type usbPortSut struct {
	data int
	sut  *USBPort
}

func makeUSBPortSut() usbPortSut {
	data := 85
	sut := NewUSBPort(data)

	return usbPortSut{data, sut}
}
