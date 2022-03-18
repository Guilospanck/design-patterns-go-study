package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NikeShoe_SetLogo(t *testing.T) {
	t.Run("Should set right logo in the nike shoe", func(t *testing.T) {
		// arrange
		sut := makeNikeShoeSut()

		// act
		sut.shoe.SetLogo(sut.logo)

		// assert
		assert.Equal(t, sut.logo, sut.shoe.logo)
	})
}

func Test_NikeShoe_GetLogo(t *testing.T) {
	t.Run("Should get right logo in the nike shoe", func(t *testing.T) {
		// arrange
		sut := makeNikeShoeSut()
		sut.shoe.logo = sut.logo

		// act
		result := sut.shoe.GetLogo()

		// assert
		assert.Equal(t, sut.logo, result)
	})
}

func Test_NikeShoe_SetSize(t *testing.T) {
	t.Run("Should set right size in the nike shoe", func(t *testing.T) {
		// arrange
		sut := makeNikeShoeSut()

		// act
		sut.shoe.SetSize(sut.size)

		// assert
		assert.Equal(t, sut.size, sut.shoe.size)
	})
}

func Test_NikeShoe_GetSize(t *testing.T) {
	t.Run("Should get right size in the nike shoe", func(t *testing.T) {
		// arrange
		sut := makeNikeShoeSut()
		sut.shoe.size = sut.size

		// act
		result := sut.shoe.GetSize()

		// assert
		assert.Equal(t, sut.size, result)
	})
}

type nikeShoeSut struct {
	shoe *NikeShoe
	logo string
	size int
}

func makeNikeShoeSut() nikeShoeSut {
	shoe := &NikeShoe{}
	logo := "logo"
	size := 1

	return nikeShoeSut{shoe, logo, size}
}
