package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Factory_GetSportsFactoryAdidas(t *testing.T) {
	t.Run("Should create adidas factory properly", func(t *testing.T) {
		// arrange
		sut := makeFactorySut()
		brand := "adidas"

		// act
		result, err := GetSportsFactory(brand)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, sut.adidas, result)

	})
}
func Test_Factory_GetSportsFactoryNike(t *testing.T) {
	t.Run("Should create nike factory properly", func(t *testing.T) {
		// arrange
		sut := makeFactorySut()
		brand := "nike"

		// act
		result, err := GetSportsFactory(brand)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, sut.nike, result)

	})
}

func Test_Factory_GetSportsFactoryDefault(t *testing.T) {
	t.Run("Should return error if no right brand specified", func(t *testing.T) {
		// arrange
		brand := "some brand"

		// act
		result, err := GetSportsFactory(brand)

		// assert
		assert.Error(t, err)
		assert.Nil(t, result)

	})
}

type factorySut struct {
	adidas *adidas
	nike   *nike

	adidasShoe *adidasShoe
	nikeShoe   *NikeShoe

	adidasShirt *adidasShirt
	nikeShirt   *NikeShirt
}

func makeFactorySut() factorySut {
	adidas := &adidas{
		shoe:  &adidasShoe{},
		shirt: &adidasShirt{},
	}
	nike := &nike{
		shoe:  &NikeShoe{},
		shirt: &NikeShirt{},
	}

	adidasShoe := &adidasShoe{}
	nikeShoe := &NikeShoe{}

	adidasShirt := &adidasShirt{}
	nikeShirt := &NikeShirt{}

	return factorySut{
		adidas,
		nike,
		adidasShoe,
		nikeShoe,
		adidasShirt,
		nikeShirt,
	}
}
