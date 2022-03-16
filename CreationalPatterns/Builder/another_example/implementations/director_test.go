package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Director_SetBuilder(t *testing.T) {
	t.Run("should set right builder", func(t *testing.T) {
		// arrange
		sut := makeDirectorSut()
		otherBuilder := NewSensorsBuilder("test_run")

		// act and assert
		assert.Equal(t, sut.sensorBuilder, sut.director.builder)

		sut.director.SetBuilder(otherBuilder)
		assert.Equal(t, otherBuilder, sut.director.builder)

	})
}

type directorSut struct {
	director      *director
	sensorBuilder *SensorsBuilder
}

func makeDirectorSut() directorSut {
	sensorBuilder := NewSensorsBuilder("test")
	director := NewDirector(sensorBuilder)

	return directorSut{director, sensorBuilder}
}
