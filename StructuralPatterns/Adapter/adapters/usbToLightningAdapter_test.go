package adapters

import (
	"base/StructuralPatterns/Adapter/interfaces"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UsbToLightningAdapter_Transform(t *testing.T) {
	t.Run("Should return instance of lightning port", func(t *testing.T) {
		// arrange
		sut := makeUSBToLightningAdapterSut()
		sut.usbPortSpy.On("GetData").Return(sut.data)

		// act
		result := sut.adapter.Transform()

		// assert
		assert.Implements(t, (*interfaces.ILightningPort)(nil), result)
		sut.usbPortSpy.AssertCalled(t, "GetData")
		sut.usbPortSpy.AssertNumberOfCalls(t, "GetData", 1)
		sut.usbPortSpy.AssertExpectations(t)

	})
}

type usbToLightningAdapterSut struct {
	usbPortSpy *usbPortSpy
	adapter    *USBToLightningAdapter
	data       int
}

func makeUSBToLightningAdapterSut() usbToLightningAdapterSut {
	data := 10
	usbPortSpy := NewUSBPortSpy(data)
	adapter := NewUSBToLightningAdapter(usbPortSpy)

	return usbToLightningAdapterSut{usbPortSpy, adapter, data}
}
