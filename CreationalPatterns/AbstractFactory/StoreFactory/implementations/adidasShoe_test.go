package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdidasShoe_SetLogo(t *testing.T) {
	t.Run("Should set right logo in the adidas shoe", func(t *testing.T) {
		// arrange
		sut := makeAdidasShoeSut()

		// act
		sut.shoe.SetLogo(sut.logo)

		// assert
		assert.Equal(t, sut.logo, sut.shoe.logo)
	})
}

func Test_AdidasShoe_GetLogo(t *testing.T) {
	t.Run("Should get right logo in the adidas shoe", func(t *testing.T) {
		// arrange
		sut := makeAdidasShoeSut()
		sut.shoe.logo = sut.logo

		// act
		result := sut.shoe.GetLogo()

		// assert
		assert.Equal(t, sut.logo, result)
	})
}

func Test_AdidasShoe_SetSize(t *testing.T) {
	t.Run("Should set right size in the adidas shoe", func(t *testing.T) {
		// arrange
		sut := makeAdidasShoeSut()

		// act
		sut.shoe.SetSize(sut.size)

		// assert
		assert.Equal(t, sut.size, sut.shoe.size)
	})
}

func Test_AdidasShoe_GetSize(t *testing.T) {
	t.Run("Should get right size in the adidas shoe", func(t *testing.T) {
		// arrange
		sut := makeAdidasShoeSut()
		sut.shoe.size = sut.size

		// act
		result := sut.shoe.GetSize()

		// assert
		assert.Equal(t, sut.size, result)
	})
}

type adidasShoeSut struct {
	shoe *adidasShoe
	logo string
	size int
}

func makeAdidasShoeSut() adidasShoeSut {
	shoe := &adidasShoe{}
	logo := "logo"
	size := 1

	return adidasShoeSut{shoe, logo, size}
}
