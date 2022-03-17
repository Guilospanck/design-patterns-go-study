package implementations

import (
	"base/CreationalPatterns/AbstractFactory/SimpleFactory/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GunFactory_NewAk47(t *testing.T) {
	t.Run("Should create AK47 from Gun Factory", func(t *testing.T) {
		// arrange
		sut := makeGunFactorySut()
		gunType := "ak47"

		// act
		result, err := NewGunFactory(gunType)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, sut.ak47)
		assert.Equal(t, result.GetName(), sut.ak47.GetName())
		assert.Equal(t, result.GetPower(), sut.ak47.GetPower())
	})
}

func Test_GunFactory_NewMusket(t *testing.T) {
	t.Run("Should create Musket from Gun Factory", func(t *testing.T) {
		// arrange
		sut := makeGunFactorySut()
		gunType := "musket"

		// act
		result, err := NewGunFactory(gunType)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, sut.musket)
		assert.Equal(t, result.GetName(), sut.musket.GetName())
		assert.Equal(t, result.GetPower(), sut.musket.GetPower())
	})
}

type gunFactorySut struct {
	ak47   interfaces.IGun
	musket interfaces.IGun
}

func makeGunFactorySut() gunFactorySut {
	ak47 := NewAK47()
	musket := NewMusket()

	return gunFactorySut{ak47, musket}
}
