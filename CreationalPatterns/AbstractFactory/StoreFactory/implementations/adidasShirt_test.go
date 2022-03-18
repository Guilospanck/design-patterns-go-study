package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdidasShirt_SetLogo(t *testing.T) {
	t.Run("Should set right logo in the adidas shirt", func(t *testing.T) {
		// arrange
		sut := makeAdidasShirtSut()

		// act
		sut.shirt.SetLogo(sut.logo)

		// assert
		assert.Equal(t, sut.logo, sut.shirt.logo)
	})
}

func Test_AdidasShirt_GetLogo(t *testing.T) {
	t.Run("Should get right logo in the adidas shirt", func(t *testing.T) {
		// arrange
		sut := makeAdidasShirtSut()
		sut.shirt.logo = sut.logo

		// act
		result := sut.shirt.GetLogo()

		// assert
		assert.Equal(t, sut.logo, result)
	})
}

func Test_AdidasShirt_SetSize(t *testing.T) {
	t.Run("Should set right size in the adidas shirt", func(t *testing.T) {
		// arrange
		sut := makeAdidasShirtSut()

		// act
		sut.shirt.SetSize(sut.size)

		// assert
		assert.Equal(t, sut.size, sut.shirt.size)
	})
}

func Test_AdidasShirt_GetSize(t *testing.T) {
	t.Run("Should get right size in the adidas shirt", func(t *testing.T) {
		// arrange
		sut := makeAdidasShirtSut()
		sut.shirt.size = sut.size

		// act
		result := sut.shirt.GetSize()

		// assert
		assert.Equal(t, sut.size, result)
	})
}

type adidasShirtSut struct {
	shirt *adidasShirt
	logo  string
	size  int
}

func makeAdidasShirtSut() adidasShirtSut {
	shirt := &adidasShirt{}
	logo := "logo"
	size := 1

	return adidasShirtSut{shirt, logo, size}
}
