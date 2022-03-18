package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Nike_MakeShoe(t *testing.T) {
	t.Run("should create right nike shoe", func(t *testing.T) {
		// arrange
		sut := makeNikeSut()
		sut.nikeShoeSpy.On("SetLogo", sut.logo).Return()
		sut.nikeShoeSpy.On("SetSize", sut.size).Return()

		// act
		result := sut.nike.MakeShoe()

		// assert
		assert.Equal(t, result, sut.nikeShoeSpy)
		sut.nikeShoeSpy.AssertCalled(t, "SetLogo", sut.logo)
		sut.nikeShoeSpy.AssertCalled(t, "SetSize", sut.size)
		sut.nikeShoeSpy.AssertNumberOfCalls(t, "SetLogo", 1)
		sut.nikeShoeSpy.AssertNumberOfCalls(t, "SetSize", 1)
		sut.nikeShoeSpy.AssertExpectations(t)
	})
}

func Test_Nike_MakeShirt(t *testing.T) {
	t.Run("should create right nike shirt", func(t *testing.T) {
		// arrange
		sut := makeNikeSut()
		sut.nikeShirtSpy.On("SetLogo", sut.logo).Return()
		sut.nikeShirtSpy.On("SetSize", sut.size).Return()

		// act
		result := sut.nike.MakeShirt()

		// assert
		assert.Equal(t, result, sut.nikeShirtSpy)
		sut.nikeShirtSpy.AssertCalled(t, "SetLogo", sut.logo)
		sut.nikeShirtSpy.AssertCalled(t, "SetSize", sut.size)
		sut.nikeShirtSpy.AssertNumberOfCalls(t, "SetLogo", 1)
		sut.nikeShirtSpy.AssertNumberOfCalls(t, "SetSize", 1)
		sut.nikeShirtSpy.AssertExpectations(t)
	})
}

type nikeSut struct {
	nike         *nike
	nikeShirtSpy *nikeShirtSpy
	nikeShoeSpy  *nikeShoeSpy
	logo         string
	size         int
}

func makeNikeSut() nikeSut {
	nikeShirtSpy := &nikeShirtSpy{}
	nikeShoeSpy := &nikeShoeSpy{}

	nike := &nike{nikeShoeSpy, nikeShirtSpy}

	logo := "nike"
	size := 14

	return nikeSut{nike, nikeShirtSpy, nikeShoeSpy, logo, size}
}
