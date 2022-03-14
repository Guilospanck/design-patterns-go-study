package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Radio_TurnOn(t *testing.T) {
	t.Run("should turn radio on", func(t *testing.T) {
		// arrange
		sut := makeRadioSut()

		// act
		sut.usecase.TurnOn()

		// assert
		assert.Equal(t, true, sut.usecase.isRunning)
	})
}

func Test_Radio_TurnOff(t *testing.T) {
	t.Run("should turn radio off", func(t *testing.T) {
		// arrange
		sut := makeRadioSut()

		// act
		sut.usecase.TurnOff()

		// assert
		assert.Equal(t, false, sut.usecase.isRunning)
	})
}

type radioSut struct {
	usecase *Radio
}

func makeRadioSut() radioSut {
	usecase := NewRadio()

	return radioSut{usecase}
}
