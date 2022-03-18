package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Shirt_SetLogo(t *testing.T) {
	t.Run("Should set right logo in the shirt", func(t *testing.T) {
		// arrange
		sut := makeShirtSut()

		// act
		sut.shirt.SetLogo(sut.logo)

		// assert
		assert.Equal(t, sut.logo, sut.shirt.logo)
	})
}

func Test_Shirt_GetLogo(t *testing.T) {
	t.Run("Should get right logo in the shirt", func(t *testing.T) {
		// arrange
		sut := makeShirtSut()
		sut.shirt.logo = sut.logo

		// act
		result := sut.shirt.GetLogo()

		// assert
		assert.Equal(t, sut.logo, result)
	})
}

func Test_Shirt_SetSize(t *testing.T) {
	t.Run("Should set right size in the shirt", func(t *testing.T) {
		// arrange
		sut := makeShirtSut()

		// act
		sut.shirt.SetSize(sut.size)

		// assert
		assert.Equal(t, sut.size, sut.shirt.size)
	})
}

func Test_Shirt_GetSize(t *testing.T) {
	t.Run("Should get right size in the shirt", func(t *testing.T) {
		// arrange
		sut := makeShirtSut()
		sut.shirt.size = sut.size

		// act
		result := sut.shirt.GetSize()

		// assert
		assert.Equal(t, sut.size, result)
	})
}

type shirtSut struct {
	shirt *Shirt
	logo  string
	size  int
}

func makeShirtSut() shirtSut {
	shirt := NewShirt()
	logo := "logo"
	size := 1

	return shirtSut{shirt, logo, size}
}
