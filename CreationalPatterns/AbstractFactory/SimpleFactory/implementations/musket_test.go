package implementations

import (
	"base/CreationalPatterns/AbstractFactory/SimpleFactory/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Musket_New(t *testing.T) {
	t.Run("Should create musket properly", func(t *testing.T) {
		// arrange
		sut := makeMusketSut()

		// act
		name := sut.gun.GetName()
		power := sut.gun.GetPower()

		// assert
		assert.Equal(t, "Musket", name)
		assert.Equal(t, 1, power)
	})
}

type musketSut struct {
	gun interfaces.IGun
}

func makeMusketSut() musketSut {
	gun := NewMusket()

	return musketSut{gun}
}
