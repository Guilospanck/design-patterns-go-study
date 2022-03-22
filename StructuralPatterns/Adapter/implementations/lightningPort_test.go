package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LightningPort_GetData(t *testing.T) {
	t.Run("Should return the right data", func(t *testing.T) {
		// arrange
		sut := makeLightningPortSut()

		// act
		result := sut.sut.GetData()

		// assert
		assert.Equal(t, sut.data, result)
	})
}

type lightningPortSut struct {
	data float64
	sut  *LightningPort
}

func makeLightningPortSut() lightningPortSut {
	data := 85.2
	sut := NewLightningPort(data)

	return lightningPortSut{data, sut}
}
