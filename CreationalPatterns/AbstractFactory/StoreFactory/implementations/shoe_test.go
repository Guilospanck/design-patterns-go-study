package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Shoe_SetLogo(t *testing.T) {
	t.Run("Should set right logo in the shoe", func(t *testing.T) {
		// arrange
		sut := makeShoeSut()

		// act
		sut.shoe.SetLogo(sut.logo)

		// assert
		assert.Equal(t, sut.logo, sut.shoe.logo)
	})
}

func Test_Shoe_GetLogo(t *testing.T) {
	t.Run("Should get right logo in the shoe", func(t *testing.T) {
		// arrange
		sut := makeShoeSut()
		sut.shoe.logo = sut.logo

		// act
		result := sut.shoe.GetLogo()

		// assert
		assert.Equal(t, sut.logo, result)
	})
}

func Test_Shoe_SetSize(t *testing.T) {
	t.Run("Should set right size in the shoe", func(t *testing.T) {
		// arrange
		sut := makeShoeSut()

		// act
		sut.shoe.SetSize(sut.size)

		// assert
		assert.Equal(t, sut.size, sut.shoe.size)
	})
}

func Test_Shoe_GetSize(t *testing.T) {
	t.Run("Should get right size in the shoe", func(t *testing.T) {
		// arrange
		sut := makeShoeSut()
		sut.shoe.size = sut.size

		// act
		result := sut.shoe.GetSize()

		// assert
		assert.Equal(t, sut.size, result)
	})
}

type shoeSut struct {
	shoe *Shoe
	logo string
	size int
}

func makeShoeSut() shoeSut {
	shoe := NewShoe()
	logo := "logo"
	size := 1

	return shoeSut{shoe, logo, size}
}
