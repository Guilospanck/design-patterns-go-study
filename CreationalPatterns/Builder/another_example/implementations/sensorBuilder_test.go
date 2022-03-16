package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SensorsBuilder_Reset(t *testing.T) {
	t.Run("should reset sensor options", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()
		sut.sensorsBuilder.obj.Between = true
		sut.sensorsBuilder.obj.TSample = 10

		// act
		sut.sensorsBuilder.Reset()

		// assert
		assert.Equal(t, false, sut.sensorsBuilder.obj.Between)
		assert.Equal(t, uint(0), sut.sensorsBuilder.obj.TSample)
	})
}

func Test_SensorsBuilder_SetThreshold(t *testing.T) {
	t.Run("should set threshold properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetThreshold(10)

		// assert
		assert.Equal(t, int(10), sut.sensorsBuilder.obj.Threshold)
	})
}

func Test_SensorsBuilder_SetHigher(t *testing.T) {
	t.Run("should set higher properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetHigher(true)

		// assert
		assert.Equal(t, true, sut.sensorsBuilder.obj.Higher)
	})
}

func Test_SensorsBuilder_SetLower(t *testing.T) {
	t.Run("should set lower properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetLower(true)

		// assert
		assert.Equal(t, true, sut.sensorsBuilder.obj.Lower)
	})
}

func Test_SensorsBuilder_SetTSample(t *testing.T) {
	t.Run("should set tsample properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetTSample(10)

		// assert
		assert.Equal(t, uint(10), sut.sensorsBuilder.obj.TSample)
	})
}

func Test_SensorsBuilder_SetBorder(t *testing.T) {
	t.Run("should set border properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetBorder(true)

		// assert
		assert.Equal(t, true, sut.sensorsBuilder.obj.Border)
	})
}

func Test_SensorsBuilder_SetRisingEdge(t *testing.T) {
	t.Run("should set rising edge properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetRisingEdge(true)

		// assert
		assert.Equal(t, true, sut.sensorsBuilder.obj.RisingEdge)
	})
}

func Test_SensorsBuilder_SetFallingEdge(t *testing.T) {
	t.Run("should set falling edge properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetFallingEdge(true)

		// assert
		assert.Equal(t, true, sut.sensorsBuilder.obj.FallingEdge)
	})
}

func Test_SensorsBuilder_SetBetween(t *testing.T) {
	t.Run("should set between properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetBetween(true)

		// assert
		assert.Equal(t, true, sut.sensorsBuilder.obj.Between)
	})
}

func Test_SensorsBuilder_SetFirstThreshold(t *testing.T) {
	t.Run("should set first threshold", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetFirstThreshold(10)

		// assert
		assert.Equal(t, int(10), sut.sensorsBuilder.obj.FirstThreshold)
	})
}

func Test_SensorsBuilder_SetSecondThreshold(t *testing.T) {
	t.Run("should set second threshold properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetSecondThreshold(10)

		// assert
		assert.Equal(t, int(10), sut.sensorsBuilder.obj.SecondThreshold)
	})
}

func Test_SensorsBuilder_SetType(t *testing.T) {
	t.Run("should set type properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetType(10)

		// assert
		assert.Equal(t, uint(10), sut.sensorsBuilder.obj.Type)
	})
}

func Test_SensorsBuilder_SetSocialContact(t *testing.T) {
	t.Run("should set social contact properly", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()

		// act
		sut.sensorsBuilder.SetSocialContact("test@test.com")

		// assert
		assert.Equal(t, "test@test.com", sut.sensorsBuilder.obj.SocialContact)
	})
}

func Test_SensorsBuilder_GetSensor(t *testing.T) {
	t.Run("should get right sensor with sensorType defined and reset the one inside the builder", func(t *testing.T) {
		// arrange
		sut := makeSensorsBuilderSut()
		sut.sensorsBuilder.obj.Between = true
		sut.sensorsBuilder.obj.FallingEdge = true
		sut.sensorsBuilder.obj.RisingEdge = true

		// act
		result := sut.sensorsBuilder.GetSensor()

		// assert
		assert.Equal(t, true, result[sut.sensorType].Between)
		assert.Equal(t, true, result[sut.sensorType].FallingEdge)
		assert.Equal(t, true, result[sut.sensorType].RisingEdge)
		assert.Equal(t, false, sut.sensorsBuilder.obj.Between)
		assert.Equal(t, false, sut.sensorsBuilder.obj.FallingEdge)
		assert.Equal(t, false, sut.sensorsBuilder.obj.RisingEdge)
	})
}

type sensorsBuilderSut struct {
	sensorsBuilder *SensorsBuilder
	sensorType     string
}

func makeSensorsBuilderSut() sensorsBuilderSut {
	sensorType := "test"
	sensorsBuilder := NewSensorsBuilder(sensorType)

	return sensorsBuilderSut{sensorsBuilder, sensorType}
}
