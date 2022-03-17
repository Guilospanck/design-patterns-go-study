package implementations

import (
	"base/CreationalPatterns/AbstractFactory/SimpleFactory/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AK47_New(t *testing.T) {
	t.Run("Should create ak47 properly", func(t *testing.T) {
		// arrange
		sut := makeAK47Sut()

		// act
		name := sut.gun.GetName()
		power := sut.gun.GetPower()

		// assert
		assert.Equal(t, "AK47", name)
		assert.Equal(t, 4, power)
	})
}

type ak47Sut struct {
	gun interfaces.IGun
}

func makeAK47Sut() ak47Sut {
	gun := NewAK47()

	return ak47Sut{gun}
}
