package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Adidas_MakeShoe(t *testing.T) {
	t.Run("should create right adidas shoe", func(t *testing.T) {
		// arrange
		sut := makeAdidasSut()
		sut.adidasShoeSpy.On("SetLogo", sut.logo).Return()
		sut.adidasShoeSpy.On("SetSize", sut.size).Return()

		// act
		result := sut.adidas.MakeShoe()

		// assert
		assert.Equal(t, result, sut.adidasShoeSpy)
		sut.adidasShoeSpy.AssertCalled(t, "SetLogo", sut.logo)
		sut.adidasShoeSpy.AssertCalled(t, "SetSize", sut.size)
		sut.adidasShoeSpy.AssertNumberOfCalls(t, "SetLogo", 1)
		sut.adidasShoeSpy.AssertNumberOfCalls(t, "SetSize", 1)
		sut.adidasShoeSpy.AssertExpectations(t)
	})
}

func Test_Adidas_MakeShirt(t *testing.T) {
	t.Run("should create right adidas shirt", func(t *testing.T) {
		// arrange
		sut := makeAdidasSut()
		sut.adidasShirtSpy.On("SetLogo", sut.logo).Return()
		sut.adidasShirtSpy.On("SetSize", sut.size).Return()

		// act
		result := sut.adidas.MakeShirt()

		// assert
		assert.Equal(t, result, sut.adidasShirtSpy)
		sut.adidasShirtSpy.AssertCalled(t, "SetLogo", sut.logo)
		sut.adidasShirtSpy.AssertCalled(t, "SetSize", sut.size)
		sut.adidasShirtSpy.AssertNumberOfCalls(t, "SetLogo", 1)
		sut.adidasShirtSpy.AssertNumberOfCalls(t, "SetSize", 1)
		sut.adidasShirtSpy.AssertExpectations(t)
	})
}

type adidasSut struct {
	adidas         *adidas
	adidasShirtSpy *adidasShirtSpy
	adidasShoeSpy  *adidasShoeSpy
	logo           string
	size           int
}

func makeAdidasSut() adidasSut {
	adidasShirtSpy := &adidasShirtSpy{}
	adidasShoeSpy := &adidasShoeSpy{}

	adidas := &adidas{adidasShoeSpy, adidasShirtSpy}

	logo := "adidas"
	size := 14

	return adidasSut{adidas, adidasShirtSpy, adidasShoeSpy, logo, size}
}
