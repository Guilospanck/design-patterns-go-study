package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NikeShirt_SetLogo(t *testing.T) {
	t.Run("Should set right logo in the nike shirt", func(t *testing.T) {
		// arrange
		sut := makeNikeShirtSut()

		// act
		sut.shirt.SetLogo(sut.logo)

		// assert
		assert.Equal(t, sut.logo, sut.shirt.logo)
	})
}

func Test_NikeShirt_GetLogo(t *testing.T) {
	t.Run("Should get right logo in the nike shirt", func(t *testing.T) {
		// arrange
		sut := makeNikeShirtSut()
		sut.shirt.logo = sut.logo

		// act
		result := sut.shirt.GetLogo()

		// assert
		assert.Equal(t, sut.logo, result)
	})
}

func Test_NikeShirt_SetSize(t *testing.T) {
	t.Run("Should set right size in the nike shirt", func(t *testing.T) {
		// arrange
		sut := makeNikeShirtSut()

		// act
		sut.shirt.SetSize(sut.size)

		// assert
		assert.Equal(t, sut.size, sut.shirt.size)
	})
}

func Test_NikeShirt_GetSize(t *testing.T) {
	t.Run("Should get right size in the nike shirt", func(t *testing.T) {
		// arrange
		sut := makeNikeShirtSut()
		sut.shirt.size = sut.size

		// act
		result := sut.shirt.GetSize()

		// assert
		assert.Equal(t, sut.size, result)
	})
}

type nikeShirtSut struct {
	shirt *NikeShirt
	logo  string
	size  int
}

func makeNikeShirtSut() nikeShirtSut {
	shirt := &NikeShirt{}
	logo := "logo"
	size := 1

	return nikeShirtSut{shirt, logo, size}
}
