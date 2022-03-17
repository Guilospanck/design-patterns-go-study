package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Gun_GetName(t *testing.T) {
	t.Run("Should return gun's name", func(t *testing.T) {
		// arrange
		sut := makeGunSut()
		sut.gun.name = sut.name

		// act
		result := sut.gun.GetName()

		// assert
		assert.Equal(t, sut.name, result)
	})
}

func Test_Gun_SetName(t *testing.T) {
	t.Run("Should set gun's name", func(t *testing.T) {
		// arrange
		sut := makeGunSut()

		// act
		sut.gun.SetName(sut.name)

		// assert
		assert.Equal(t, sut.name, sut.gun.name)
	})
}

func Test_Gun_GetPower(t *testing.T) {
	t.Run("Should return gun's power", func(t *testing.T) {
		// arrange
		sut := makeGunSut()
		sut.gun.power = sut.power

		// act
		result := sut.gun.GetPower()

		// assert
		assert.Equal(t, sut.power, result)
	})
}

func Test_Gun_SetPower(t *testing.T) {
	t.Run("Should set gun's power", func(t *testing.T) {
		// arrange
		sut := makeGunSut()

		// act
		sut.gun.SetPower(sut.power)

		// assert
		assert.Equal(t, sut.power, sut.gun.power)
	})
}

type gunSut struct {
	gun   *Gun
	name  string
	power int
}

func makeGunSut() gunSut {
	gun := NewGun()

	name := "Gun name"
	power := 10

	return gunSut{gun, name, power}
}
