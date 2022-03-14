package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TV_TurnOn(t *testing.T) {
	t.Run("it should turn tv on", func(t *testing.T) {
		// arrange
		sut := makeTVSut()

		// act
		sut.usecase.TurnOn()

		// assert
		assert.Equal(t, true, sut.usecase.isRunning)
	})
}

func Test_TV_TurnOff(t *testing.T) {
	t.Run("it should turn tv off", func(t *testing.T) {
		// arrange
		sut := makeTVSut()

		// act
		sut.usecase.TurnOff()

		// assert
		assert.Equal(t, false, sut.usecase.isRunning)
	})
}

type tvSut struct {
	usecase *TV
}

func makeTVSut() tvSut {
	usecase := NewTV()

	return tvSut{usecase}
}
